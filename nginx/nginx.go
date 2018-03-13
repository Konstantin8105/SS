package nginx

import "github.com/Konstantin8105/SS/module"

func init() {
	var n Nginx
	module.Register("nginx", n)
}

type Nginx struct {
}

func (m Nginx) Test() {
}

func (m Nginx) Set() {
}
