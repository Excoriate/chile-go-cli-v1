package main

import (
	config "github.com/Excoriate/chile-go-cli-v1/pkg/config/commands"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

var (
	version = "0.0.0" // populated at compile time
	commit  = "HEAD"  // populated at compile time
	secret  = "dev"   // populated at compile time
	updates chan string
)

func initLogging() {
	log.SetOutput(os.Stderr)
}

func getCommands() []*cli.Command {
	var commands []*cli.Command
	
	commands = append(commands, config.GetCommands()...)
	return commands
	
	//return append(commands, logs.GetCommands()...)
}

func main() {
	initLogging()

	flags := []cli.Flag{
		&cli.BoolFlag{
			Name:  "debug",
			Usage: "enables debug level output",
		},
		&cli.BoolFlag{
			Name:  "skip-version-check",
			Usage: "disables version checking for this run",
		},
		&cli.BoolFlag{
			Name:  "clear-cache",
			Usage: "clears the local cache",
		},
	}

	app := cli.NewApp()
	app.Version = version
	app.Commands = getCommands()
	app.Flags = flags
	//app.Before = beforeAction
	app.Name = "ChileCLI"
	app.Usage = "Simple lab to have fun doing weird things with data :)"
	app.EnableBashCompletion = true
	//app.Metadata = map[string]interface{}{
	//	dazncli.ContextKey: context.Background(),
	//}

	if err := app.Run(os.Args); err != nil {
		panic(err)
	}

	if updates == nil {
		return
	}

	select {
	case updateMsg := <-updates:
		if updateMsg != "" {
			c := color.New(color.FgRed)
			c.Fprintf(os.Stderr, "\n%s\n", updateMsg)
		}
	case <-time.After(3 * time.Second):
		logrus.Debug("Timeout checking for updates")
	}
}
