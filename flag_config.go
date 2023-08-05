package flag

import (
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	config *viper.Viper
	path   string
}

// 新的配置文件
func newConfig(path string) Config {
	if _, err := os.Stat(path); err != nil {
		os.Create(path)
	}
	var c Config
	c.path = path
	c.config = viper.New()
	c.config.AddConfigPath(path)
	c.config.SetConfigFile(path)
	config := strings.Split(path, ".")
	if len(config) >= 2 {
		c.config.SetConfigType(config[len(config)-1])
	} else {
		c.config.SetConfigType("ini")
	}
	c.config.ReadInConfig()
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
	c.config.Set(readKey, value)
	c.config.WriteConfig()
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
	return c.config.Get(readKey)
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

func (c *Config) GetString(key string) string {
	return c.config.GetString(key)
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
	return c.config.GetStringMapString(readKey)
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
	return c.config.GetStringMap(readKey)
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
	data, ok := c.config.Get(readKey).([]any)
	if ok {
		return data
	} else {
		return []any{c.config.Get(readKey)}
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
	return c.config.ConfigFileUsed()
}

// 获取配置文件所有Key
func (c *Config) GetAllConfigKeys() []string {
	return c.config.AllKeys()
}

// 获取配置文件所有Key和Value
func (c *Config) GetAllConfigValues() map[string]any {
	return c.config.AllSettings()
}
