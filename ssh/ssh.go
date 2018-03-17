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

	// TODO:
	/*
		Create `SSH` connection:
		```
		$ sudo vim /etc/ssh/sshd_config
		```
		Generate RSA

		Change config:
		```
		Port 222
		MaxAuthTries 3 # add after port
		ChallengeResponseAuthentication no
		PasswordAuthentication no
		UsePAM no
		```
	*/

	return nil
}
