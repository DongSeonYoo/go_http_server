package main

import (
	"flag"

	"github.com/dongseonyoo/go_http_server/init/cmd"
)

/*
	entry point in application
*/

var configPathFlag = flag.String("config", "./init/config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configPathFlag)
}
