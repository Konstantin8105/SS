package starter

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func run(programArgs ...string) (_ []byte, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Cannot run program. err = %v", err)
		}
	}()
	command := strings.Split(commandPrefix, " ")
	command = append(command, programArgs...)
again:
	for inx := range command {
		c := strings.TrimSpace(command[inx])
		if len(c) == 0 {
			if inx == len(command) {
				command = command[:inx]
				continue
			}
			command = append(command[:inx], command[inx+1:]...)
			goto again
		}
	}

	// with stderr analyze
	cmd := exec.Command(command[0], command[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("err = %v.\n stderr = %v", err, stderr.String())
		return
	}
	return out.Bytes(), nil
}

func IsInstalled(name string) error {
	out, err := run("dpkg", "-s", name)
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

func Install(name string) (err error) {
	_, err = run("apt", "install", "-y", name)
	if err != nil {
		return fmt.Errorf("cannot build program. err = %v", err)
	}

	fmt.Printf("Install program `%s`\n", name)

	return nil
}
