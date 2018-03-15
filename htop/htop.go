package htop

import (
	"fmt"

	"github.com/Konstantin8105/SS/starter"
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

func (m Htop) Test() error {
	fmt.Println("htop: Test")
	err := starter.IsInstalled(name)
	if err != nil {
		return err
	}
	return nil
}

func (m Htop) Set() {
	fmt.Println("htop : Set ")
}
