package nginx

import (
	"fmt"
)

func init() {
	var n Nginx
	_ = n
	// starter.Register("nginx", n)
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
