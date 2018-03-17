package starter

import (
	"sync"
)

// Starter - base interface
type Starter interface {
	Run() error
}

var (
	starters      map[string]Starter
	m             sync.Mutex
	commandPrefix string
)

func init() {
	starters = map[string]Starter{}
}

// List returns a sorted list of the names of the registered starters.
func List() map[string]Starter {
	m.Lock()
	defer m.Unlock()

	return starters
}

// SetCommandPrefix - add prefix before each command.
// Typically used for ssh, docker, ...
// Example for ssh prefix is `ssh some@localhost`
func SetCommandPrefix(prefix string) {
	m.Lock()
	defer m.Unlock()

	commandPrefix = prefix
}

// Register makes a starter available by the provided name.
// If Register is called twice with the same name or if starter is nil,
// it panics.
func Register(name string, starter Starter) {
	m.Lock()
	defer m.Unlock()

	if starter == nil {
		panic("Starter is nil")
	}
	if _, dup := starters[name]; dup {
		panic("Starter called twice for " + name)
	}

	starters[name] = starter
}
