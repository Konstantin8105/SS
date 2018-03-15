package starter

import (
	"bytes"
	"fmt"
	"os/exec"
)

func IsInstalled(name string) error {
	out, err := exec.Command(
		fmt.Sprintf("dpkg -s %s | grep Status", name)).CombinedOutput()
	if err != nil {
		t.Fatalf("cannot build program. err = %v", err)
	}

	if !bytes.Contains(out, []byte("install ok installed")) {
		return fmt.Errorf("Program %s is not installed : %v",
			name, string(out))
	}

	return nil
}
