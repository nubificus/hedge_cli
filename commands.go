package main

import (
	"errors"
	"fmt"

	hedge "github.com/nubificus/hedge_cli/hedge_api"

	"github.com/urfave/cli"
)

var startCommand = cli.Command{
	Name:        "start",
	Usage:       "start a VM",
	Description: "start a VM",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "kernel",
			Usage:       "path to kernel file",
			EnvVar:      "",
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			Destination: new(string),
		},
		cli.StringFlag{
			Name:        "name",
			Usage:       "unique name of the VM",
			EnvVar:      "",
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			Destination: new(string),
		},
		cli.IntFlag{
			Name:        "core",
			Usage:       "CPU core, where the VM will run",
			EnvVar:      "",
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			Value:       0,
			Destination: new(int),
		},
		cli.IntFlag{
			Name:        "mem",
			Usage:       "VM memory in MBs",
			EnvVar:      "",
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			Value:       0,
			Destination: new(int),
		},
		cli.StringFlag{
			Name:        "blk",
			Usage:       "path to block device",
			EnvVar:      "",
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			Destination: new(string),
		},
		cli.StringFlag{
			Name:        "net",
			Usage:       "network interface",
			EnvVar:      "",
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			Destination: new(string),
		},
		cli.StringFlag{
			Name:        "cmdline",
			Usage:       "command line argument for the guest kernel",
			EnvVar:      "",
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			Destination: new(string),
		},
	},
	Action: func(c *cli.Context) (retError error) {
		defer func() {
			if retError != nil {
				fmt.Println(retError.Error())
			}
		}()
		kernel := c.String("kernel")
		name := c.String("name")
		blk := c.String("blk")
		net := c.String("net")
		cmdline := c.String("cmdline")
		stringFlags := []string{kernel, name, blk, net, cmdline}

		for _, stringFlag := range stringFlags {
			if stringFlag == "" {
				return errors.New("please specify a VM name")
			}
		}
		mem := c.Int("mem")
		if mem == 0 {
			return errors.New("memory must be greater than 0Mbs")
		}

		core := c.Int("core")

		err := hedge.StartVM(kernel, name, core, mem,
			blk, net, cmdline)
		if err != nil {
			return err
		}
		fmt.Println("VM started")
		return nil
	},
}

var stopCommand = cli.Command{
	Name:        "stop",
	Usage:       "stop a VM",
	Description: "stop a VM",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "name",
			Usage:       "name of the VM to sto",
			EnvVar:      "",
			FilePath:    "",
			Required:    true,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			Destination: new(string),
		},
	},
	Action: func(c *cli.Context) (retError error) {
		defer func() {
			if retError != nil {
				fmt.Println(retError.Error())
			}
		}()
		name := c.String("name")
		if name == "" {
			return errors.New("please specify a VM name")
		}
		err := hedge.StopVM(name)
		if err != nil {
			return err
		}
		fmt.Println("VM stopped")
		return nil
	},
}

var showCommand = cli.Command{
	Name:        "show",
	Usage:       "show all running VMs",
	Description: "show all running VMs",
	Action: func(c *cli.Context) (retError error) {
		defer func() {
			if retError != nil {
				fmt.Println(retError.Error())
			}
		}()
		vms, err := hedge.ShowVMs()
		if err != nil {
			return err
		}
		fmt.Println(vms)
		return nil
	},
}

var consoleCommand = cli.Command{
	Name:        "console",
	Usage:       "show the console of VM with given ID",
	Description: "show the console of VM with given ID",
	Flags: []cli.Flag{
		cli.IntFlag{
			Name:     "vm",
			Usage:    "ID of the VM to show console",
			EnvVar:   "",
			FilePath: "",
			Required: true,
			Hidden:   false,
		},
	},
	Action: func(c *cli.Context) (retError error) {
		defer func() {
			if retError != nil {
				fmt.Println(retError.Error())
			}
		}()
		vm := c.Int("vm")
		console, err := hedge.ShowConsole(vm)
		if err != nil {
			return err
		}
		fmt.Println(console)
		return nil
	},
}

func Commands() []cli.Command {
	commands := []cli.Command{
		startCommand,
		stopCommand,
		showCommand,
		consoleCommand,
	}
	return commands
}
