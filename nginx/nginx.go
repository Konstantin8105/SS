package nginx

import (
	"fmt"

	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n Nginx
	n.ProgramName = "nginx"
	starter.Register(n.ProgramName, n)
}

// Nginx - program `nginx`
type Nginx struct {
	starter.SimpleInstall
}

// Run - running preparing the program
func (n Nginx) Run() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cannot run %v. err = %v", n.ProgramName, err)
		}
	}()

	// install
	n.SimpleInstall.Run()
	if err != nil {
		return err
	}

	return nil
}
