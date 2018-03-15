package starter

import (
	"fmt"
	"sort"
	"sync"
)

// Starter - base interface
type Starter interface {
	Test() error
	Set()
}

var (
	starters map[string]Starter
	m        sync.Mutex
)

func init() {
	starters = map[string]Starter{}
}

// List returns a sorted list of the names of the registered starters.
func List() []string {
	m.Lock()
	defer m.Unlock()

	var list []string
	for name := range starters {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
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

func isStarter(starterName string) {
	for s := range starters {
		if s == starterName {
			return
		}
	}
	panic(fmt.Errorf("Cannot found starter with name : %s", starterName))
}
