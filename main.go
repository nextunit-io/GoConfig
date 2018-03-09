package config

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

var Cfg = Config{}

type Config struct {
	lookupDone map[string]bool
	variables  map[string]interface{}
}

func (cfg *Config) SetDefault(name string, value interface{}) {
	cfg.variables[name] = value
}

func (cfg *Config) SetDefaults(defaults map[string]interface{}) {
	for k, v := range defaults {
		cfg.SetDefault(k, v)
	}
}

func (cfg *Config) Get(name string) (interface{}, error) {
	if !cfg.lookupDone[name] {
		log.Debugf("Configuration: New lookup for variable '%s'.", name)

		cfg.lookupDone[name] = true
		env := os.Getenv(name)

		if env != "" {
			cfg.variables[name] = env
		}
	}

	if val, ok := cfg.variables[name]; ok {
		return val, nil
	}

	log.Errorf("Configuration: Variable '%s' not found.", name)
	return "", errors.New("Variable not found")
}
