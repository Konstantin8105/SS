package main

import (
	"flag"
	"fmt"

	"github.com/Konstantin8105/SS/module"
	_ "github.com/Konstantin8105/SS/nginx"
)

var (
	versionFlag = flag.Bool("v", false, "print the version and exit")
	testFlag    = flag.Bool("t", false, "testing settings")
	setFlag     = flag.Bool("s", false, "set settings")
)

func main() {
	flag.Parse()
	flag.Usage()
	fmt.Println("drivers : ", module.Drivers())
	// nginx
	// ssh
	// backup
	// systemd
	// git server
	// git web
}
