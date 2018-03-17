package nano

import (
	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n starter.SimpleInstall
	n.ProgramName = "nano"
	starter.Register(n.ProgramName, n)
}
