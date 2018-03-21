package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/Konstantin8105/ss/starter"
)

func TestListSize(t *testing.T) {
	if len(starter.List()) == 0 {
		t.Fatalf("starter list is empty")
	}
	if len(starter.List()) != 6 {
		t.Fatalf("starter list have uncorrect size : %v",
			len(starter.List()))
	}
}

func TestListCLI(t *testing.T) {
	f := true
	listFlag = &f
	defer func() {
		// return value back
		f = false
		listFlag = &f
	}()

	oldOutput := output
	defer func() {
		output = oldOutput
	}()
	var buf bytes.Buffer
	output = &buf

	err := run()
	if err != nil {
		t.Errorf("List cli error : %v", err)
	}

	if buf.Len() == 0 {
		t.Fatalf("Buffer is empty")
	}

	list := starter.List()
	for name := range list {
		if !strings.Contains(buf.String(), name) {
			t.Errorf("Cannot found starter name = `%s`", name)
		}
	}
}

func TestHelp(t *testing.T) {
	_, err := exec.Command("go", "build").CombinedOutput()
	if err != nil {
		t.Fatalf("cannot build program. err = %v", err)
	}

	out, err := exec.Command("./ss").CombinedOutput()
	if err != nil {
		t.Fatalf("cannot run `./ss`. err = %v", err)
	}

	expect := []byte(`Usage of ./ss:
  -h	give this help list
  -i	install settings
  -l	show list of modules
  -prefix string
    	prefix before each command. Typically used :"sudo" or "ssh tom@localhost sudo" or ...
`)

	if len(out) != len(expect) {
		t.Errorf("Length is %v\nExpected length is %v\n",
			len(out), len(expect))
	}

	if !bytes.Equal(out, expect) {
		t.Errorf("result of help argument is not expected.\n%v\nExpected:\n%v\n",
			string(out), string(expect))
	}
}

// Inside travis no need to in docker container
func TestTravis(t *testing.T) {
	if os.Getenv("TRAVIS") != "true" {
		return
	}
	oldPrefix := starter.GetCommandPrefix()
	p := "sudo"
	prefixFlag = &p
	defer func() {
		starter.SetCommandPrefix(oldPrefix)
		prefixFlag = &oldPrefix
	}()

	f := true
	installFlag = &f
	defer func() {
		// return value back
		f = false
		installFlag = &f
	}()

	err := run()
	if err != nil {
		t.Errorf("Travis test error : %v", err)
	}
}

func TestLocally(t *testing.T) {
	if os.Getenv("TRAVIS") == "true" {
		return
	}
	oldPrefix := starter.GetCommandPrefix()
	p := "sudo"
	prefixFlag = &p
	defer func() {
		starter.SetCommandPrefix(oldPrefix)
		prefixFlag = &oldPrefix
	}()
	// # Example of creating container in according
	// # to ubuntu image
	// → docker container create ubuntu:16.04
	// 7f41440ca37ff95715c8af66f16fa432e47c341e6593b73d0de02173220ce706
	// out, err := exec.Command(
	// 	"docker", "container", "create", "ubuntu:16.04").CombinedOutput()
	// if err != nil {
	// 	t.Fatalf("cannot create container. err = %v", err)
	// }

	// starter.SetCommandPrefix(" echo ")
	// TODO : starter.SetCommandPrefix(" docker run ubuntu:16.04 ")

	f := true
	installFlag = &f
	defer func() {
		// return value back
		f = false
		installFlag = &f
	}()

	err := run()
	if err != nil {
		t.Errorf("Locally test error : %v", err)
	}
}
