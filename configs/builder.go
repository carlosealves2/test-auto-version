package configs

import (
	"os"
	"strconv"
)

type ConfigBuilder struct {
	cfg *ApplicationConfig
}

func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{}
}

func (cb *ConfigBuilder) WithEnv() *ConfigBuilder {
	cb.cfg = &ApplicationConfig{
		Env:     getEnv("APP_ENV", "development"),
		Port:    getEnv("APP_PORT", "8080"),
		Version: getEnv("APP_VERSION", "0.1.0"),
	}

	return cb
}

func (cb *ConfigBuilder) Validate() *ConfigBuilder {
	if cb.cfg == nil {
		panic("ConfigBuilder: cfg is nil, please call WithEnv first")
	}

	return cb
}

func (cb *ConfigBuilder) Build() *ApplicationConfig {
	return cb.cfg
}

func getEnv[T any](key string, defaultValue T) T {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	var result any

	switch any(defaultValue).(type) {
	case string:
		result = value
	case int:
		intValue, err := strconv.Atoi(value)
		if err != nil {
			result = defaultValue
		} else {
			result = intValue
		}
	case bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			result = defaultValue
		} else {
			result = boolValue
		}
	default:
		result = defaultValue
	}

	return result.(T)
}
