package ssh

import (
	"fmt"

	"github.com/Konstantin8105/ss/starter"
)

func init() {
	var n SSH
	n.ProgramName = "ssh"
	starter.Register(n.ProgramName, n)
}

// SSH - program `ssh`
type SSH struct {
	starter.SimpleInstall
}

// Run - running preparing the program
func (n SSH) Run() (err error) {
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

	// change settings
	location := "/etc/ssh/sshd_config"
	err = starter.WriteFile(location, []byte(sshConfig))
	if err != nil {
		return
	}

	// restart ssh
	if err = starter.ServiceAction(starter.ServiceRestart, "ssh"); err != nil {
		return
	}

	// change permissions for configs

	// TODO:

	// status ssh
	if err = starter.ServiceAction(starter.ServiceStatus, "ssh"); err != nil {
		return
	}

	return nil
}
