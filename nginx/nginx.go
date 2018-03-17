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

func (m Nginx) Run() (err error) {
	fmt.Println("nginx : Run")
	return nil
}
