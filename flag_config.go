package flag

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
	path string
}

// 新的配置文件
func newConfig(path string) Config {
	if _, err := os.Stat(path); err != nil {
		os.Create(path)
	}
	var c Config
	c.path = path
	c.Viper = viper.New()
	c.Viper.AddConfigPath(path)
	c.Viper.SetConfigFile(path)
	config := strings.Split(path, ".")
	if len(config) >= 2 {
		c.Viper.SetConfigType(config[len(config)-1])
	} else {
		c.Viper.SetConfigType("toml")
	}
	c.Viper.ReadInConfig()
	return c
}

// 写入配置文件
func (c *Config) WriteToConfig(title, key string, value any) {
	var readKey string
	if title != "" && key != "" {
		readKey = title + "." + key
	} else {
		if title == "" {
			readKey = key
		}
		if key == "" {
			readKey = title
		}
	}
	c.Viper.Set(readKey, value)
	c.Viper.WriteConfig()
}

// 读取配置信息为任意类型
func (c *Config) ReadConfigToAny(title, key string) any {
	var readKey string
	if title != "" && key != "" {
		readKey = title + "." + key
	} else {
		if title == "" {
			readKey = key
		}
		if key == "" {
			readKey = title
		}
	}
	return c.Viper.Get(readKey)
}

// 读取配置信息为string类型
func (c *Config) ReadConfigToString(title, key string) string {
	var readKey string
	if title != "" && key != "" {
		readKey = title + "." + key
	} else {
		if title == "" {
			readKey = key
		}
		if key == "" {
			readKey = title
		}
	}
	return c.GetString(readKey)
}

// 写入配置文件
func (c *Config) WriteStringToConfig(title, key, value string) {
	c.WriteToConfig(title, key, value)
}

// 读取配置信息为string类型map
func (c *Config) ReadConfigToStringMap(title, key string) map[string]string {
	var readKey string
	if title != "" && key != "" {
		readKey = title + "." + key
	} else {
		if title == "" {
			readKey = key
		}
		if key == "" {
			readKey = title
		}
	}
	return c.Viper.GetStringMapString(readKey)
}

func (c *Config) ReadConfigToKVMap(title, key string) map[string]any {
	var readKey string
	if title != "" && key != "" {
		readKey = title + "." + key
	} else {
		if title == "" {
			readKey = key
		}
		if key == "" {
			readKey = title
		}
	}
	return c.Viper.GetStringMap(readKey)
}

// 写入配置文件
func (c *Config) WriteKVMapToConfig(title, key string, value map[string]any) {
	c.WriteToConfig(title, key, value)
}

func (c *Config) WriteStringMapToConfig(title, key string, value map[string]string) {
	c.WriteToConfig(title, key, value)
}

// 读取配置信息为string类型map
func (c *Config) ReadConfigToAnySlice(title, key string) []any {
	var readKey string
	if title != "" && key != "" {
		readKey = title + "." + key
	} else {
		if title == "" {
			readKey = key
		}
		if key == "" {
			readKey = title
		}
	}
	data, ok := c.Viper.Get(readKey).([]any)
	if ok {
		return data
	} else {
		return []any{c.Viper.Get(readKey)}
	}
}

// 写入配置文件
func (c *Config) WriteStringMapsToConfig(title, key string, value []map[string]any) {
	c.WriteToConfig(title, key, value)
}

// 删除配置文件的值
func (c *Config) DeleteConfigValue(title, key string) {
	c.WriteToConfig(title, key, "")
}

// 获取配置文件名称
func (c *Config) GetConfigPath() string {
	return c.Viper.ConfigFileUsed()
}

// 获取配置文件所有Key
func (c *Config) GetAllConfigKeys() []string {
	return c.Viper.AllKeys()
}

// 获取配置文件所有Key和Value
func (c *Config) GetAllConfigValues() map[string]any {
	return c.Viper.AllSettings()
}
