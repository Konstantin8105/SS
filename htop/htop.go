package htop

import (
	"fmt"

	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n Htop
	starter.Register(name, n)
}

var (
	name = "htop"
)

// Htop - program `htop`
type Htop struct {
}

// Run - running preparing the program
func (m Htop) Run() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cannot run %v. err = %v", name, err)
		}
	}()

	// install
	err = starter.IsInstalled(name)
	if err != nil {
		err = starter.Install(name)
		if err != nil {
			return err
		}
	}

	return nil
}
