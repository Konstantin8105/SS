package mc

import (
	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n starter.SimpleInstall
	n.ProgramName = "mc"
	starter.Register(n.ProgramName, n)
}
