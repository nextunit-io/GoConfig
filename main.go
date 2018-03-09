package config

import (
	"errors"
	"os"

	log "github.com/sirupsen/logrus"
)

var Cfg = Config{
	lookupDone: map[interface{}]bool{},
	variables:  map[interface{}]interface{}{},
}

type Config struct {
	lookupDone map[interface{}]bool
	variables  map[interface{}]interface{}
}

func (cfg *Config) SetDefault(name interface{}, value interface{}) {
	cfg.variables[name] = value
}

func (cfg *Config) SetDefaults(defaults map[interface{}]interface{}) {
	for k, v := range defaults {
		cfg.SetDefault(k, v)
	}
}

func (cfg *Config) Get(name interface{}) (interface{}, error) {
	if !cfg.lookupDone[name] {
		log.Debugf("Configuration: New lookup for variable '%s'.", name)

		cfg.lookupDone[name] = true

		if stringName, ok := name.(string); ok {
			env := os.Getenv(stringName)

			if env != "" {
				cfg.variables[name] = env
			}
		}
	}

	if val, ok := cfg.variables[name]; ok {
		return val, nil
	}

	log.Errorf("Configuration: Variable '%s' not found.", name)
	return "", errors.New("Variable not found")
}
