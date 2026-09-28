package flag

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

// ConfigFileBuilder configures the optional file and its command-line path flag.
// Name and Alias should be called before Parse.
type ConfigFileBuilder struct {
	set         *FlagSet
	path        string
	aliases     []string
	invalidName bool
}

// ConfigFile opts the package's default FlagSet into loading path when no
// command-line config path is supplied.
func ConfigFile(path string) *ConfigFileBuilder {
	return GetDefaultFlagSet().ConfigFile(path)
}

// ConfigFile opts this FlagSet into loading path when no command-line config
// path is supplied. Repeated calls update the default path on this FlagSet.
func (f *FlagSet) ConfigFile(path string) *ConfigFileBuilder {
	if f.configFile == nil {
		f.configFile = &ConfigFileBuilder{set: f}
	}
	f.configFile.path = path
	f.config = nil
	return f.configFile
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
