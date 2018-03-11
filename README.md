# GO Config Module

## Usage

It is easy to use this module. It is kind of a variable store and you can override variables via environment variables.

```go
package x

import (
	"fmt"

	"github.com/nextunit-io/go-config"
)

func init() {
	createDefaultConfig()
}

func main() {
	val, err := config.Cfg.Get("TEST_ENTRY_ONE")
	if err == nil {
		fmt.Printf("Found entry for 'TEST_ENTRY_ONE': %s", val.(string))
	}

	val, err = config.Cfg.Get("TEST_ENTRY_TWO")
	if err == nil {
		if val.(bool) {
			fmt.Print("Found entry for 'TEST_ENTRY_TWO' and it is true")
		}
    }
    
    os.Setenv("TEST_ENTRY_THREE", "Overridden config")
    val, err = config.Cfg.Get("TEST_ENTRY_THREE")
	if err == nil {
		fmt.Printf("Found entry for 'TEST_ENTRY_THREE': %s", val.(string))
	}

	val, err = config.Cfg.Get("TEST_INVALID_ENTRY")
	if err != nil {
		fmt.Print("No entry for 'TEST_INVALID_ENTRY' found.")
	}
}

func createDefaultConfig() {
	config.Cfg.SetDefaults(map[interface{}]interface{}{
		"TEST_ENTRY_ONE": "this-is-the-value",
		"TEST_ENTRY_TWO": true,
		"TEST_ENTRY_THREE": "hidden",
	})
}
```

Output:
```
Found entry for 'TEST_ENTRY_ONE': this-is-the-value
Found entry for 'TEST_ENTRY_TWO' and it is true
Found entry for 'TEST_ENTRY_TREE': Overridden config
No entry for 'TEST_INVALID_ENTRY' found.
```