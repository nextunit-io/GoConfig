package config

import (
	"os"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestConfiguration(t *testing.T) {
	Cfg.SetDefaults(map[interface{}]interface{}{
		"test":  "test1",
		"test2": "test3",
		"test4": "test5",
		"test6": true,
	})

	testGetValue(t, "test", "test1")
	testGetValue(t, "test2", "test3")
	testGetValue(t, "test4", "test5")
	testGetValue(t, "test6", true)

	Cfg.SetDefault("test5", "test6")
	Cfg.SetDefault("test4", "test7")

	testGetValue(t, "test5", "test6")
	testGetValue(t, "test4", "test7")

	Cfg.SetDefault("test8", "invalid")
	os.Setenv("test8", "bla")
	testGetValue(t, "test8", "bla")

	_, err := Cfg.Get("invalid_not_used")
	assert.Equal(t, err, VariableNotFoundError)
}

func testGetValue(t *testing.T, name string, expected interface{}) {
	val, err := Cfg.Get(name)

	assert.Equal(t, err, nil)
	assert.Equal(t, val, expected)
}
