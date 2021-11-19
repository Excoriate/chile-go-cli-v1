package config

import (
	"fmt"
	"github.com/urfave/cli/v2"
)


const (
	ContextKey = "ctx"
)

func PingHandler(c *cli.Context) error {
	fmt.Println("Pong")
	return nil
}
