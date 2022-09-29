package main

import (
	"errors"

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
	Action: func(c *cli.Context) error {
		var retErr error

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

		ret := hedge.Start_vm(kernel, name, core, mem,
			blk, net, cmdline)
		if ret == 0 {
			return nil
		}
		return retErr
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
	Action: func(c *cli.Context) error {
		var retErr error

		name := c.String("name")
		if name == "" {
			return errors.New("please specify a VM name")
		}
		ret := hedge.Stop_vm(name)
		if ret == 0 {
			return nil
		}
		return retErr
	},
}

var showCommand = cli.Command{
	Name:        "show",
	Usage:       "show all running VMs",
	Description: "show all running VMs",
	Action: func(c *cli.Context) error {
		hedge.Show_vms()
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
	Action: func(c *cli.Context) error {
		vm := c.Int("vm")

		hedge.Show_cons(vm)
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
