package module

import (
	"sort"

	"github.com/Konstantin8105/SS/driver"
)

var (
	drivers map[string]driver.Driver
)

func init() {
	drivers = map[string]driver.Driver{}
}

// Drivers returns a sorted list of the names of the registered drivers.
func Drivers() []string {
	// driversMu.RLock()
	// defer driversMu.RUnlock()
	var list []string
	for name := range drivers {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, driver driver.Driver) {
	// driversMu.Lock()
	// defer driversMu.Unlock()

	if driver == nil {
		panic("sql: Register driver is nil")
	}

	if _, dup := drivers[name]; dup {
		panic("sql: Register called twice for driver " + name)
	}

	drivers[name] = driver

}
