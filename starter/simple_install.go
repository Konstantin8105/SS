package starter

import (
	"fmt"
)

// SimpleInstall - base type for program just install
type SimpleInstall struct {
	ProgramName string // Name of program
}

// Run - install program without any other settings
func (s SimpleInstall) Run() (err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cannot run %v. err = %v", s.ProgramName, err)
		}
	}()

	// install
	err = IsInstalled(s.ProgramName)
	if err != nil {
		err = Install(s.ProgramName)
		if err != nil {
			return err
		}
	}

	return nil
}
