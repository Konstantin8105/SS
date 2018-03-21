package starter

import (
	"sync"
)

// Starter - base interface
type Starter interface {
	Run() error
}

type StarterItem struct {
	Name string
	S    Starter
}

var (
	starters      []StarterItem
	m             sync.Mutex
	commandPrefix string
)

// List returns a sorted list of the names of the registered starters.
func List() []StarterItem {
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

// GetCommandPrefix return used command prefix
func GetCommandPrefix() string {
	m.Lock()
	defer m.Unlock()

	return commandPrefix
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
	for _, item := range starters {
		if item.Name == name {
			panic("Starter called twice for " + name)
		}
	}

	starters = append(starters, StarterItem{
		Name: name,
		S:    starter,
	})
}
