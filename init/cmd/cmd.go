package cmd

import (
	"fmt"

	"github.com/dongseonyoo/go_http_server/config"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filepath),
	}

	fmt.Println(c.config.Server.Port)

	return c
}
