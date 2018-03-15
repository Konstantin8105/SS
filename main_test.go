package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/Konstantin8105/ss/starter"
	"github.com/bradleyjkemp/cupaloy"
)

func setupTest(args []string) (*bytes.Buffer, func()) {
	buf := &bytes.Buffer{}
	oldOutput := output
	oldArgs := os.Args

	output = buf
	os.Args = args

	return buf, func() {
		output = oldOutput
		os.Args = oldArgs
	}
}

var cliTests = map[string][]string{
	"List": {"./ss", "-l"},
}

func TestCLI(t *testing.T) {
	for testName, args := range cliTests {
		t.Run(testName, func(t *testing.T) {
			output, teardown := setupTest(args)
			defer teardown()

			run()

			err := cupaloy.SnapshotMulti(testName, output)
			if err != nil {
				t.Fatalf("error: %s", err)
			}
		})
	}
	// return value back
	f := false
	listFlag = &f
}

func TestEmptyList(t *testing.T) {
	if len(starter.List()) == 0 {
		t.Fatalf("starter list is empty")
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
  -s	set settings
  -t	testing settings
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

func TestSettingTest(t *testing.T) {
	os.Args = []string{"program", "-t"}
	code := run()
	if code != 0 {
		t.Errorf("Exit code for flag '-t' is %v", code)
	}
	// return value back
	f := false
	testFlag = &f
}

/*
func TestIntegration(t *testing.T) {
	if os.Getenv("TRAVIS") != "true" {
		// This is not inside travis, so
		// run inside docker
	}
}
*/
