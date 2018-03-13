package main

import (
	"flag"
	"fmt"

	// nginx
	_ "github.com/Konstantin8105/SS/nginx"

	// ssh
	// backup
	// systemd
	// git server
	// git web
	// localhost
	// router settings

	// base `starter` package
	"github.com/Konstantin8105/SS/starter"
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

func main() {
	flag.Parse()

	switch {
	case *versionFlag:
		// version flag
		fmt.Printf("Version : %s\n", version)

	case *listFlag:
		// list of modules
		fmt.Printf("List of starters :\n")
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Printf("%2d%20s\n", inx, s)
		}
		fmt.Printf("Amount of starters : %2d\n", len(list))

	case *testFlag:
		// testing settings
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Printf("%2d%20s\n", inx, s)
			starter.Test(s)
		}

	case *setFlag:
		// set settings
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Printf("%2d%20s\n", inx, s)
			starter.Set(s)
		}

	default:
		// help flag
		flag.Usage()
	}
}
