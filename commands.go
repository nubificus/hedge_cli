package main

import (
	"errors"
	"strconv"

	hedge "github.com/nubificus/hedge_cli/hedge_api"

	"github.com/urfave/cli"
)

var startCommand = cli.Command{
	Name:        "start",
	Usage:       "Start a VM",
	Description: "Start a VM",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:        "kernel",
			Usage:       "Path to kernel file",
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
			Usage:       "Unique name of the VM",
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
			Usage:       "CPU core, where the VM will run (default: 0)",
			EnvVar:      "",
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			Value:       0,
			Destination: new(int),
		},
		cli.IntFlag{
			Name:        "mem",
			Usage:       "VM memory in MBs (default: 512)",
			EnvVar:      "",
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			Value:       512,
			Destination: new(int),
		},
		cli.StringFlag{
			Name:        "blk",
			Usage:       "Path to block device",
			EnvVar:      "",
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			Destination: new(string),
		},
		cli.StringFlag{
			Name:        "net",
			Usage:       "Network interface",
			EnvVar:      "",
			FilePath:    "",
			Required:    false,
			Hidden:      false,
			TakesFile:   false,
			Value:       "",
			Destination: new(string),
		},
		cli.StringFlag{
			Name:        "cmdline",
			Usage:       "Command line argument for the guest kernel",
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

		stringFlags := []string{kernel, name, cmdline}

		for _, stringFlag := range stringFlags {
			if stringFlag == "" {
				return errors.New("Please specify the kernel path, the name of the VM and the cmdline")
			}
		}
		mem := c.Int("mem")
		if mem <= 0 {
			return errors.New("Memory must be greater than 0 MBs")
		}

		core := c.Int("core")
		if core < 0 {
			return errors.New("Core cannot be a negative value")
		}

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
	Usage:       "Stop a VM, with the given name",
	Description: "Stop a VM, with the given name",
	Action: func(c *cli.Context) error {
		var retErr error

		name := c.Args().First()
		if name == "" {
			return errors.New("Please specify a VM name")
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
	Usage:       "Show all running VMs",
	Description: "Show all running VMs",
	Action: func(c *cli.Context) error {
		hedge.Show_vms()
		return nil
	},
}

var consoleCommand = cli.Command{
	Name:        "console",
	Usage:       "Show the console of VM with given ID",
	Description: "Show the console of VM with given ID",
	Action: func(c *cli.Context) error {
		vm := c.Int("vm")

		vm, err := strconv.Atoi(c.Args().First())
		if err != nil || vm < 0 {
			return errors.New("Please specify a valid VM id")
		}
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
