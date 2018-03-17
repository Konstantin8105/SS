package starter

import "testing"

func TestSimpleInstallFail(t *testing.T) {
	s := SimpleInstall{ProgramName: "SomeWrongProgramName"}
	err := s.Run()
	if err == nil {
		t.Errorf("Cannot fail for program with wrong name")
	}
}
