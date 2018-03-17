package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/Konstantin8105/ss/starter"
)

func TestListSize(t *testing.T) {
	if len(starter.List()) == 0 {
		t.Fatalf("starter list is empty")
	}
	if len(starter.List()) != 2 {
		t.Fatalf("starter list have uncorrect size : %v",
			len(starter.List()))
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
  -l	show list of modules
  -r	install settings
  -v	print the version and exit
`)

	if len(out) != len(expect) {
		t.Errorf("Length is %v\nExpected length is %v\n",
			len(out), len(expect))
	}

	if bytes.Compare(out, expect) != 0 {
		t.Errorf("result of help argument is not expected.\n%v\nExpected:\n%v\n",
			string(out), string(expect))
	}
}

// Inside travis no need to in docker container
func TestTravis(t *testing.T) {
	if os.Getenv("TRAVIS") != "true" {
		return
	}
	starter.SetCommandPrefix("")

	f := true
	runFlag = &f
	defer func() {
		f = false
		runFlag = &f
	}()

	err := run()
	if err != nil {
		t.Errorf("Travis test error : %v", err)
	}
	// return value back
}

func TestLocally(t *testing.T) {
	if os.Getenv("TRAVIS") == "true" {
		return
	}
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

	var err error
	{
		var fl bool
		fl = true
		runFlag = &fl
	}

	err = run()
	if err != nil {
		t.Errorf("Locally test error : %v", err)
	}

	{
		var fl bool
		fl = false
		runFlag = &fl
	}
}
