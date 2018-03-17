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

	// vim
	_ "github.com/Konstantin8105/ss/vim"

	// nano
	_ "github.com/Konstantin8105/ss/nano"

	// ssh
	_ "github.com/Konstantin8105/ss/ssh"

	// backup
	// systemd
	// git server
	// git web
	// localhost
	// router settings
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
	installFlag = flag.Bool("i", false, "install settings")
	prefixFlag  = flag.String("prefix", "", "prefix before each command. Typically used :\"sudo\" or \"ssh tom@localhost sudo\" or ...")
)

/*
Notes:
* https://blog.golang.org/docker
* https://stackoverflow.com/questions/26411594/executing-docker-command-using-golang-exec-fails
* https://github.com/betweenbrain/ubuntu-web-server-build-script
* https://medium.com/statuscode/golang-docker-for-development-and-production-ce3ad4e69673
*/

func main() {
	flag.Parse()
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

var output io.Writer = os.Stdout

func run() (err error) {

	if len(*prefixFlag) != 0 {
		starter.SetCommandPrefix(*prefixFlag)
	}

	switch {
	case *versionFlag:
		// version flag
		fmt.Fprintf(output, "Version : %s\n", version)

	case *listFlag:
		// list of modules
		fmt.Fprintf(output, "List of starters :\n")
		list := starter.List()
		var inx int
		for name := range list {
			inx++
			fmt.Fprintf(output, "%2d%20s\n", inx, name)
		}
		fmt.Fprintf(output, "Amount of starters : %2d\n", len(list))

	case *installFlag:
		// set settings
		list := starter.List()
		var inx int
		for name, s := range list {
			inx++
			fmt.Fprintf(output, "%2d%20s\n", inx, name)
			err = s.Run()
			if err != nil {
				return err
			}
		}

	default:
		// help flag
		flag.Usage()
	}
	return nil
}
