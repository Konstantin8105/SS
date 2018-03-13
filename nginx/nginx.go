package nginx

import "github.com/Konstantin8105/SS/starter"

func init() {
	var n Nginx
	starter.Register("nginx", n)
}

type Nginx struct {
}

func (m Nginx) Test() {
}

func (m Nginx) Set() {
}
