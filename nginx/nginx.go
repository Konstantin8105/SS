package nginx

import (
	"fmt"

	"github.com/Konstantin8105/SS/starter"
)

func init() {
	var n Nginx
	starter.Register("nginx", n)
}

type Nginx struct {
}

func (m Nginx) Test() {
	fmt.Println("nginx : Test")
}

func (m Nginx) Set() {
	fmt.Println("nginx : Set ")
}
