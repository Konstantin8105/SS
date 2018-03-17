package starter

import (
	"bytes"
	"testing"
)

func TestRunWithEmptyArgument(t *testing.T) {
	out, err := run("", "echo", "", "\"Hello\"", "")
	if err != nil {
		t.Errorf("Error for test with empty arguments")
	}
	if bytes.Compare(out, []byte("\"Hello\"\n")) != 0 {
		t.Errorf("Cannot get correct result. output = `%s`",
			string(out))
	}
}

func TestInstallFail(t *testing.T) {
	name := "SomeStrangeProgram"
	if IsInstalled(name) == nil {
		t.Errorf("Haven`t error for check install program with fail name")
	}
	if Install(name) == nil {
		t.Errorf("Haven`t error for install program with fail name")
	}
}
