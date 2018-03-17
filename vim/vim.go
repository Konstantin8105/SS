package vim

import (
	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n starter.SimpleInstall
	n.ProgramName = "vim"
	starter.Register(n.ProgramName, n)
}
