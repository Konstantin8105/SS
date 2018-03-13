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
	version string = "0.1"
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

	// help flag
	if *helpFlag {
		flag.Usage()
		return
	}

	// version flag
	if *versionFlag {
		fmt.Println("Version : ", version)
		return
	}

	// list of modules
	if *listFlag {
		fmt.Println("List of starters :")
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Printf("%2d%20s\n", inx, s)
		}
		fmt.Printf("Amount of starters : %2d\n", len(list))
		return
	}

	// testing settings
	if *testFlag {
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Printf("%2d%20s\n", inx, s)
			starter.Test(s)
		}
		return
	}

	// set settings
	if *setFlag {
		list := starter.List()
		for inx, s := range list {
			inx := inx + 1
			fmt.Printf("%2d%20s\n", inx, s)
			starter.Set(s)
		}
		return
	}
}
