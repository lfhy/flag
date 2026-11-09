package flag

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// ConfigFileBuilder configures the optional file and its command-line path flag.
// Name and Alias should be called before Parse.
type ConfigFileBuilder struct {
	set         *FlagSet
	path        string
	paths       []string
	layered     bool
	aliases     []string
	invalidName bool
}

// ConfigFile opts the package's default FlagSet into loading path when no
// command-line config path is supplied.
func ConfigFile(path string) *ConfigFileBuilder {
	return GetDefaultFlagSet().ConfigFile(path)
}

// ConfigFiles loads optional default files in order, from lowest to highest
// priority. An explicit -c file adds the highest file layer; -c= disables all
// files. Command-line and environment values still override every file.
func ConfigFiles(paths ...string) *ConfigFileBuilder {
	return GetDefaultFlagSet().ConfigFiles(paths...)
}

// ConfigFile opts this FlagSet into loading path when no command-line config
// path is supplied. Repeated calls update the default path on this FlagSet.
func (f *FlagSet) ConfigFile(path string) *ConfigFileBuilder {
	if f.configFile == nil {
		f.configFile = &ConfigFileBuilder{set: f}
	}
	f.configFile.path = path
	f.configFile.paths = nil
	f.configFile.layered = false
	f.config = nil
	return f.configFile
}

// ConfigFiles configures multiple optional default files. Later files override
// earlier values, nested tables merge, and arrays replace rather than append.
// Unlike ConfigFile, an explicit -c is an additional highest-priority layer.
func (f *FlagSet) ConfigFiles(paths ...string) *ConfigFileBuilder {
	if f.configFile == nil {
		f.configFile = &ConfigFileBuilder{set: f}
	}
	f.configFile.paths = append([]string(nil), paths...)
	f.configFile.layered = true
	f.configFile.path = ""
	for _, path := range paths {
		if path != "" {
			f.configFile.path = path
		}
	}
	f.config = nil
	return f.configFile
}

type configFileSource struct {
	path     string
	required bool
}

func (b *ConfigFileBuilder) readMerged() (*Config, error) {
	sources := make([]configFileSource, 0, len(b.paths)+1)
	for _, path := range b.paths {
		if path != "" {
			sources = append(sources, configFileSource{path: path})
		}
	}
	canonical := b.set.ConfigFlagName()
	if flag := b.set.formal[canonical]; flag != nil && b.set.actual[canonical] != nil {
		if path := flag.Value.String(); path != "" {
			sources = append(sources, configFileSource{path: path, required: true})
		} else {
			sources = nil
		}
	}
	merged := newReadOnlyConfig("")
	settings := make(map[string]interface{})
	for _, source := range sources {
		if _, err := os.Stat(source.path); err != nil {
			if os.IsNotExist(err) && !source.required {
				continue
			}
			return nil, fmt.Errorf("无法读取配置文件 %q: %w", source.path, err)
		}
		layer := newReadOnlyConfig(source.path)
		if err := layer.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("无法读取配置文件 %q: %w", source.path, err)
		}
		mergeConfigLayers(settings, layer.AllSettings())
		merged.path = source.path
		merged.SetConfigFile(source.path)
	}
	if err := merged.MergeConfigMap(settings); err != nil {
		return nil, fmt.Errorf("无法合并配置文件: %w", err)
	}
	return &merged, nil
}

// Merge only two tables recursively. A higher-priority scalar or array must
// replace an earlier table too; Viper's incremental MergeConfigMap otherwise
// logs that type change and silently retains the lower-priority table.
func mergeConfigLayers(target, overlay map[string]interface{}) {
	for key, value := range overlay {
		next, nextIsMap := value.(map[string]interface{})
		previous, previousIsMap := target[key].(map[string]interface{})
		if nextIsMap && previousIsMap {
			mergeConfigLayers(previous, next)
		} else {
			target[key] = value
		}
	}
}

// Native config arrays replace defaults as well as lower file layers. Keep
// comma-containing string elements intact instead of round-tripping via CSV.
func setLayeredConfigValue(value Value, raw interface{}, text string) error {
	var target reflect.Value
	switch value := value.(type) {
	case *StringSlice:
		target = reflect.ValueOf(value).Elem()
	case *IntSlice:
		target = reflect.ValueOf(value).Elem()
	case *int64sValue:
		target = reflect.ValueOf(value.value).Elem()
	case *uintsValue:
		target = reflect.ValueOf(value.value).Elem()
	case *uint64sValue:
		target = reflect.ValueOf(value.value).Elem()
	case *reflectSliceValue:
		target = value.elem
	default:
		return value.Set(text)
	}
	parse, err := sliceElemParser(target.Type().Elem())
	if err != nil {
		return err
	}
	slice := reflect.MakeSlice(target.Type(), 0, 0)
	source := reflect.ValueOf(raw)
	var parts []string
	if source.IsValid() && (source.Kind() == reflect.Slice || source.Kind() == reflect.Array) {
		for i := 0; i < source.Len(); i++ {
			parts = append(parts, fmt.Sprint(source.Index(i).Interface()))
		}
	} else {
		parts = splitStrings(text)
	}
	for _, part := range parts {
		element, err := parse(part)
		if err != nil {
			return err
		}
		if !element.Type().AssignableTo(target.Type().Elem()) {
			element = element.Convert(target.Type().Elem())
		}
		slice = reflect.Append(slice, element)
	}
	target.Set(slice)
	return nil
}

// Name changes the canonical config-path flag name ("c" by default).
func (b *ConfigFileBuilder) Name(name string) *ConfigFileBuilder {
	b.invalidName = name == ""
	if !b.invalidName {
		b.set.SetConfigFlagName(name)
	}
	return b
}

// Alias adds another name for the same config-path flag.
func (b *ConfigFileBuilder) Alias(name string) *ConfigFileBuilder {
	b.aliases = append(b.aliases, name)
	return b
}

// register validates before the auto-hidden flag is added, so invalid builder
// input is a Parse error under ContinueOnError rather than a registration panic.
func (b *ConfigFileBuilder) register(canonical string) error {
	if b.invalidName || canonical == "" {
		return fmt.Errorf("配置文件参数名不能为空")
	}
	if b.set.formal[canonical] == nil && b.set.Lookup(canonical) != nil {
		return fmt.Errorf("配置文件参数名 -%s 已被参数别名占用", canonical)
	}
	seen := make(map[string]bool, len(b.aliases))
	for _, alias := range b.aliases {
		if alias == "" {
			return fmt.Errorf("配置文件参数别名不能为空")
		}
		if alias == canonical {
			continue
		}
		if seen[alias] {
			return fmt.Errorf("配置文件参数别名重复定义: %s", alias)
		}
		if flag := b.set.Lookup(alias); flag != nil && flag.Name != canonical {
			return fmt.Errorf("配置文件参数别名重复定义: %s", alias)
		}
		seen[alias] = true
	}
	return nil
}

func (b *ConfigFileBuilder) installAliases(canonical string) {
	for _, alias := range b.aliases {
		if alias != canonical && b.set.Lookup(alias) == nil {
			b.set.Alias(canonical, alias)
		}
	}
}

// newReadOnlyConfig configures Viper without creating a missing file. Callers
// decide whether a read error is optional (GetConfig) or fatal (ParseFile).
func newReadOnlyConfig(path string) Config {
	c := Config{Viper: viper.New(), path: path}
	if path == "" {
		return c
	}
	c.Viper.SetConfigFile(path)
	if extension := strings.TrimPrefix(filepath.Ext(path), "."); extension != "" {
		c.Viper.SetConfigType(extension)
	} else {
		c.Viper.SetConfigType("toml")
	}
	return c
}
