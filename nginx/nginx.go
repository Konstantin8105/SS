package nginx

import (
	"fmt"

	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n Nginx
	starter.Register("nginx", n)
}

type Nginx struct {
}

func (m Nginx) Test() (err error) {
	fmt.Println("nginx : Test")
	return nil
}

func (m Nginx) Set() {
	fmt.Println("nginx : Set ")
}
