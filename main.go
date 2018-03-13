package main

import (
	"flag"
	"fmt"

	_ "github.com/Konstantin8105/SS/nginx"
	"github.com/Konstantin8105/SS/starter"
)

var (
	versionFlag = flag.Bool("v", false, "print the version and exit")
	testFlag    = flag.Bool("t", false, "testing settings")
	setFlag     = flag.Bool("s", false, "set settings")
)

func main() {
	flag.Parse()
	flag.Usage()
	fmt.Println("starters : ", starter.List())
	// nginx
	// ssh
	// backup
	// systemd
	// git server
	// git web
	// localhost
	// router settings
}
