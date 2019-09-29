package iris

import (
	"github.com/kataras/iris"
	"strings"
)

type Config struct {
	iris *iris.Application
}

// Quickly convert all other configurations
// @todo waiting testing
func (c *Config) Get(key string) interface{} {
	if strings.Index(key, `.`) != -1 {
		keys := strings.Split(key, `.`)
		var mapValue map[interface{}]interface{}
		length := len(keys)
		for i, k := range keys {
			//last
			if length-1 == i {
				return mapValue[k]
			} else if i == 0 {
				mapValue = c.iris.ConfigurationReadOnly().GetOther()[k].(map[interface{}]interface{})
			} else {
				mapValue = mapValue[k].(map[interface{}]interface{})
			}
		}
	}

	return c.iris.ConfigurationReadOnly().GetOther()[key]
}

// Quickly convert all other configurations
// If is nil return default value
func (c *Config) GetDefault(key string, defaultValue interface{}) interface{} {
	value := c.Get(key)

	// empty string conversion return defaultValue
	if v, ok := value.(string); ok && v == "" {
		return defaultValue
	}

	return value
}



func NewConfig(iris *iris.Application) *Config {
	return &Config{
		iris: iris,
	}
}
