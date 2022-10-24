package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "hedgecli"
	app.Usage = "A CLI tool to manage hedge VMs"
	app.Version = "0.0.3"
	// 	app.Description = `
	// hedgecli is a tool that allows users to start, stop and manage hedge VMs`
	app.Commands = Commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
