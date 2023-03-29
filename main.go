package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "hedge_cli"
	app.Usage = "A CLI tool to manage hedge VMs"
	app.Version = "0.0.3"
	app.Commands = append(app.Commands, commands...)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
