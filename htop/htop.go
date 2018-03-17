package htop

import (
	"fmt"

	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n Htop
	starter.Register("htop", n)
}

var (
	name = "htop"
)

type Htop struct {
}

func (m Htop) Run() (err error) {
	fmt.Println("htop: Run")

	err = starter.IsInstalled(name)
	if err != nil {
		err = starter.Install(name)
		if err != nil {
			return err
		}
	}

	return nil
}
