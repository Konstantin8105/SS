package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	// nginx
	_ "github.com/Konstantin8105/ss/nginx"

	// htop
	_ "github.com/Konstantin8105/ss/htop"

	// ssh
	// backup
	// systemd
	// git server
	// git web
	// localhost
	// router settings
	// vim
	// system update

	// base `starter` package
	"github.com/Konstantin8105/ss/starter"
)

const (
	version string = "0.2"
)

var (
	helpFlag    = flag.Bool("h", false, "give this help list")
	versionFlag = flag.Bool("v", false, "print the version and exit")
	listFlag    = flag.Bool("l", false, "show list of modules")
	testFlag    = flag.Bool("t", false, "testing settings")
	setFlag     = flag.Bool("s", false, "set settings")
)

/*
Notes:
* https://blog.golang.org/docker
* https://stackoverflow.com/questions/26411594/executing-docker-command-using-golang-exec-fails
* https://github.com/betweenbrain/ubuntu-web-server-build-script
* https://medium.com/statuscode/golang-docker-for-development-and-production-ce3ad4e69673
*/

func main() {
	code := run()
	os.Exit(code)
}

var output io.Writer = os.Stdout

func run() int {
	flag.Parse()

	switch {
	case *versionFlag:
		// version flag
		fmt.Fprintf(output, "Version : %s\n", version)

	case *listFlag:
		// list of modules
		fmt.Fprintf(output, "List of starters :\n")
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Fprintf(output, "%2d%20s\n", inx, s)
		}
		fmt.Fprintf(output, "Amount of starters : %2d\n", len(list))

	case *testFlag:
		// testing settings
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Fprintf(output, "%2d%20s\n", inx, s)
			err := starter.Test(s)
			if err != nil {
				return 1
			}
		}

	case *setFlag:
		// set settings
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Fprintf(output, "%2d%20s\n", inx, s)
			starter.Set(s)
		}

	default:
		// help flag
		flag.Usage()
	}
	return 0
}
