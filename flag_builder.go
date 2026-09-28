package flag

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// FlagBuilder collects a flag's settings before registering its value.
type FlagBuilder struct {
	set             *FlagSet
	option          FlagVar
	invalidConfig   bool
	configErr       error
	attempted       bool
	registered      bool
	registrationErr error
	value           Value
}

// FlagValue reads the current value of a builder-defined flag.
// Numeric and boolean getters panic if the current text cannot be converted.
type FlagValue struct {
	builder *FlagBuilder
}

// rawBuilderValue keeps pending flags as text, even if their defaults are typed.
type rawBuilderValue struct {
	text string
}

func (v *rawBuilderValue) String() string    { return v.text }
func (v *rawBuilderValue) UsageType() string { return "string" }

func (v *rawBuilderValue) Set(text string) error {
	v.text = text
	return nil
}

type boolBuilderValue struct{ *rawBuilderValue }

func (v *boolBuilderValue) Set(text string) error {
	if _, err := strconv.ParseBool(text); err != nil {
		return err
	}
	return v.rawBuilderValue.Set(text)
}

func (v *boolBuilderValue) IsBoolFlag() bool { return true }

// New starts defining a flag on the package's default FlagSet.
func New(name string) *FlagBuilder {
	return GetDefaultFlagSet().New(name)
}

// New starts defining a flag on f. A terminal method registers it immediately;
// otherwise Parse registers it with string-backed storage.
func (f *FlagSet) New(name string) *FlagBuilder {
	b := &FlagBuilder{set: f, option: FlagVar{Name: name}}
	f.pending = append(f.pending, b)
	return b
}

// Alias adds another command-line name for the flag.
func (b *FlagBuilder) Alias(alias string) *FlagBuilder {
	b.option.Aliases = append(b.option.Aliases, alias)
	return b
}

// Default sets the value used when no other source supplies one.
func (b *FlagBuilder) Default(value any) *FlagBuilder {
	b.option.DefaultValue = value
	return b
}

// Config maps the flag to a root key or a section.key path.
// Only the first dot separates the section from its key.
func (b *FlagBuilder) Config(path string) *FlagBuilder {
	section, key, hasSection := strings.Cut(path, ".")
	if !hasSection {
		key = section
		section = ""
	}
	if key == "" {
		b.invalidConfig = true
		b.configErr = fmt.Errorf("配置键不能为空: %q", path)
		b.set.handleError(b.configErr)
		return b
	}
	b.invalidConfig = false
	b.configErr = nil
	b.option.ConfigSection = section
	b.option.ConfigKey = key
	return b
}

// Env sets the environment variable read for the flag.
func (b *FlagBuilder) Env(name string) *FlagBuilder {
	b.option.Env = name
	return b
}

// Usage sets the flag's help text.
func (b *FlagBuilder) Usage(text string) *FlagBuilder {
	b.option.Usage = text
	return b
}

// Hidden omits the flag from help output.
func (b *FlagBuilder) Hidden() *FlagBuilder {
	b.option.Hidden = true
	return b
}

// Var binds a pointer to the staged flag and registers it.
func (b *FlagBuilder) Var(pointer any) {
	if b.invalidConfig {
		return
	}
	if b.attempted {
		panic(fmt.Sprintf("flag %q was already registered", b.option.Name))
	}
	b.attempted = true
	option := b.option
	option.Value = pointer
	name := option.Name
	if name == "" && option.ConfigSection != "" && option.ConfigKey != "" {
		name = option.ConfigSection + "-" + option.ConfigKey
	}
	previous := b.set.Lookup(name)
	b.set.Var(&option)
	b.option.Name = option.Name
	flag := b.set.Lookup(option.Name)
	if flag != nil && flag != previous {
		b.registered = true
		b.value = flag.Value
	} else {
		b.registrationErr = fmt.Errorf("flag %q failed to register", option.Name)
	}
}

// registerPending binds a non-terminal builder to a string-backed flag.
func (b *FlagBuilder) registerPending() error {
	if b.invalidConfig {
		return b.configErr
	}
	if b.attempted {
		return b.registrationErr
	}
	name := b.option.Name
	if name == "" && b.option.ConfigSection != "" && b.option.ConfigKey != "" {
		name = b.option.ConfigSection + "-" + b.option.ConfigKey
	}
	if b.set.Lookup(name) != nil {
		return b.failPending(fmt.Errorf("flag %q is already defined", name))
	}
	seen := make(map[string]bool, len(b.option.Aliases))
	for _, alias := range b.option.Aliases {
		if alias == "" || name == "" {
			return b.failPending(fmt.Errorf("flag %q has an empty name or alias", name))
		}
		if alias == name {
			continue
		}
		if seen[alias] || b.set.Lookup(alias) != nil {
			return b.failPending(fmt.Errorf("flag %q alias %q is already defined", name, alias))
		}
		seen[alias] = true
	}
	defaultText := defaultAsText(b.option.DefaultValue)
	_, isBool := b.option.DefaultValue.(bool)
	raw := &rawBuilderValue{text: defaultText}
	if isBool {
		b.Var(&boolBuilderValue{rawBuilderValue: raw})
		return b.registrationErr
	}
	b.Var(raw)
	return b.registrationErr
}

func (b *FlagBuilder) failPending(err error) error {
	b.attempted = true
	b.registrationErr = err
	return err
}

func defaultAsText(value any) string {
	if value == nil {
		return ""
	}
	v := reflect.ValueOf(value)
	if v.Kind() == reflect.Slice || v.Kind() == reflect.Array {
		parts := make([]string, v.Len())
		for i := range parts {
			parts[i] = fmt.Sprint(v.Index(i).Interface())
		}
		return strings.Join(parts, ",")
	}
	return fmt.Sprint(value)
}

// Value returns a read-only accessor for the registered flag's current value.
// If needed, it registers the pending builder before reading the value.
func (b *FlagBuilder) Value() *FlagValue {
	return &FlagValue{builder: b}
}

func (v *FlagValue) text() string {
	b := v.builder
	if err := b.registerPending(); err != nil {
		panic(err)
	}
	if !b.registered {
		panic(fmt.Sprintf("flag %q is not registered", b.option.Name))
	}
	return b.value.String()
}

// String returns the current value without conversion.
func (v *FlagValue) String() string { return v.text() }

// Bool parses the current value as a boolean and panics on invalid text.
func (v *FlagValue) Bool() bool {
	value, err := strconv.ParseBool(v.text())
	if err != nil {
		panic(fmt.Errorf("flag value is not a bool: %w", err))
	}
	return value
}

// Int parses the current value as an int and panics on invalid text.
func (v *FlagValue) Int() int {
	value, err := strconv.ParseInt(v.text(), 0, strconv.IntSize)
	if err != nil {
		panic(fmt.Errorf("flag value is not an int: %w", err))
	}
	return int(value)
}

// Int64 parses the current value as an int64 and panics on invalid text.
func (v *FlagValue) Int64() int64 {
	value, err := strconv.ParseInt(v.text(), 0, 64)
	if err != nil {
		panic(fmt.Errorf("flag value is not an int64: %w", err))
	}
	return value
}

// Uint parses the current value as a uint and panics on invalid text.
func (v *FlagValue) Uint() uint {
	value, err := strconv.ParseUint(v.text(), 0, strconv.IntSize)
	if err != nil {
		panic(fmt.Errorf("flag value is not a uint: %w", err))
	}
	return uint(value)
}

// Uint64 parses the current value as a uint64 and panics on invalid text.
func (v *FlagValue) Uint64() uint64 {
	value, err := strconv.ParseUint(v.text(), 0, 64)
	if err != nil {
		panic(fmt.Errorf("flag value is not a uint64: %w", err))
	}
	return value
}

// Float64 parses the current value as a float64 and panics on invalid text.
func (v *FlagValue) Float64() float64 {
	value, err := strconv.ParseFloat(v.text(), 64)
	if err != nil {
		panic(fmt.Errorf("flag value is not a float64: %w", err))
	}
	return value
}

// Duration parses the current value as a duration and panics on invalid text.
func (v *FlagValue) Duration() time.Duration {
	value, err := time.ParseDuration(v.text())
	if err != nil {
		panic(fmt.Errorf("flag value is not a duration: %w", err))
	}
	return value
}

// Bool registers a bool flag and returns its value pointer.
func (b *FlagBuilder) Bool() *bool {
	p := new(bool)
	b.Var(p)
	return p
}

// String registers a string flag and returns its value pointer.
func (b *FlagBuilder) String() *string {
	p := new(string)
	b.Var(p)
	return p
}

// Int registers an int flag and returns its value pointer.
func (b *FlagBuilder) Int() *int {
	p := new(int)
	b.Var(p)
	return p
}

// Int64 registers an int64 flag and returns its value pointer.
func (b *FlagBuilder) Int64() *int64 {
	p := new(int64)
	b.Var(p)
	return p
}

// Uint registers a uint flag and returns its value pointer.
func (b *FlagBuilder) Uint() *uint {
	p := new(uint)
	b.Var(p)
	return p
}

// Uint64 registers a uint64 flag and returns its value pointer.
func (b *FlagBuilder) Uint64() *uint64 {
	p := new(uint64)
	b.Var(p)
	return p
}

// Float64 registers a float64 flag and returns its value pointer.
func (b *FlagBuilder) Float64() *float64 {
	p := new(float64)
	b.Var(p)
	return p
}

// Duration registers a duration flag and returns its value pointer.
func (b *FlagBuilder) Duration() *time.Duration {
	p := new(time.Duration)
	b.Var(p)
	return p
}

// Strings registers a string-slice flag and returns its value pointer.
func (b *FlagBuilder) Strings() *[]string {
	p := new([]string)
	b.Var(p)
	return p
}

// Ints registers an int-slice flag and returns its value pointer.
func (b *FlagBuilder) Ints() *[]int {
	p := new([]int)
	b.Var(p)
	return p
}

// Int64s registers an int64-slice flag and returns its value pointer.
func (b *FlagBuilder) Int64s() *[]int64 {
	p := new([]int64)
	b.Var(p)
	return p
}

// Uints registers a uint-slice flag and returns its value pointer.
func (b *FlagBuilder) Uints() *[]uint {
	p := new([]uint)
	b.Var(p)
	return p
}

// Uint64s registers a uint64-slice flag and returns its value pointer.
func (b *FlagBuilder) Uint64s() *[]uint64 {
	p := new([]uint64)
	b.Var(p)
	return p
}
