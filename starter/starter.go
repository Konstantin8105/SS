package starter

import (
	"sort"
)

type iStarter interface {
	Test()
	Set()
}

var (
	starters map[string]iStarter
)

func init() {
	starters = map[string]iStarter{}
}

// Drivers returns a sorted list of the names of the registered drivers.
func Drivers() []string {
	// driversMu.RLock()
	// defer driversMu.RUnlock()
	var list []string
	for name := range starters {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

// Register makes a database driver available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, starter iStarter) {
	// driversMu.Lock()
	// defer driversMu.Unlock()

	if starter == nil {
		panic("sql: Register driver is nil")
	}

	if _, dup := starters[name]; dup {
		panic("sql: Register called twice for driver " + name)
	}

	starters[name] = starter
}
