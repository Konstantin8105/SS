package htop

import (
	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n starter.SimpleInstall
	n.ProgramName = "htop"
	starter.Register(n.ProgramName, n)
}
