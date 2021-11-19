package config

import (
	config "github.com/Excoriate/chile-go-cli-v1/pkg/config/handlers"
	"github.com/urfave/cli/v2"
)

// GetCommands returns this modules' cli commands
func GetCommands() []*cli.Command {
	var commands []*cli.Command
	
	commands = append(commands, &cli.Command{
		Name:        "config",
		Usage:       "Configuration for this CLI",
		Subcommands: GetSubCommands(),
	})
	
	return commands
}

// Get SubCommands
func GetSubCommands() []*cli.Command {
	var subcommands []*cli.Command
	
	subcommands = append(subcommands, &cli.Command{
		Name:   "ping",
		Usage:  "Send a Ping and get a Pong",
		Action: config.PingHandler,
	})
	
	return subcommands
}


