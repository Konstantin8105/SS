package starter

import (
	"bytes"
	"fmt"
	"os/exec"
)

func IsInstalled(name string) error {
	out, err := exec.Command("dpkg", "-s", name).CombinedOutput()
	if err != nil {
		return fmt.Errorf("cannot build program. err = %v", err)
	}

	if !bytes.Contains(out, []byte("install ok installed")) {
		return fmt.Errorf("Program %s is not installed : %v",
			name, string(out))
	}

	fmt.Printf("Program `%s` is installed\n", name)

	return nil
}
