package main

import (
	"fmt"
	"strconv"

	hedge "github.com/nubificus/hedge_cli/hedge_api"
	"github.com/urfave/cli/v2"
)

var commands []*cli.Command = []*cli.Command{{
	Name:        "start",
	Usage:       "Start a VM",
	Description: "Start a VM",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "kernel",
			Aliases:  []string{"k"},
			Usage:    "Path to kernel file",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "name",
			Usage:    "Unique name of the VM",
			Aliases:  []string{"n"},
			Required: true,
		},
		&cli.IntFlag{
			Name:     "core",
			Aliases:  []string{"c"},
			Usage:    "CPU core, where the VM will run (default: 0)",
			Required: false,
			Value:    0,
		},
		&cli.IntFlag{
			Name:     "memory",
			Aliases:  []string{"mem", "m"},
			Usage:    "VM memory in MBs (default: 512)",
			Required: false,
			Value:    512,
		},
		&cli.StringFlag{
			Name:     "block",
			Aliases:  []string{"blk", "b"},
			Usage:    "Path to block device",
			Required: false,
			Value:    "",
		},
		&cli.StringFlag{
			Name:     "net",
			Aliases:  []string{"t"},
			Usage:    "Network interface",
			Required: false,
			Value:    "",
		},
		&cli.StringFlag{
			Name:     "cmdline",
			Aliases:  []string{"cmd"},
			Usage:    "Command line argument for the guest kernel",
			Required: true,
		},
	},
	Action: func(ctx *cli.Context) error {
		newVm := hedge.VMConfig{
			Name:    ctx.String("name"),
			Binary:  ctx.String("kernel"),
			CPU:     ctx.Int("core"),
			Mem:     ctx.Int("memory"),
			Blk:     ctx.String("block"),
			Net:     ctx.String("net"),
			CmdLine: ctx.String("cmdline"),
		}
		return hedge.StartVM(newVm)
	}},
	{
		Name:        "stop",
		Usage:       "Stop a VM with the given name",
		Description: "Stop a VM with the given name",
		Action: func(ctx *cli.Context) error {
			name := ctx.Args().First()
			if name == "" {
				return fmt.Errorf("vm name cannot be empty")
			}
			return hedge.StopVM(name)
		},
	},
	{
		Name:        "status",
		Usage:       "Check if hedge is loaded",
		Description: "Check if hedge is loaded",
		Action: func(ctx *cli.Context) error {
			err := hedge.Status()
			if err == nil {
				fmt.Println("Status: OK")
				return nil
			}
			fmt.Println("Status: ERROR")
			return err
		},
	},
	{
		Name:        "show",
		Aliases:     []string{"list", "ls"},
		Usage:       "Show all running VMs",
		Description: "Show all running VMs",
		Action: func(ctx *cli.Context) error {
			vms, err := hedge.ListVMs()
			if err != nil {
				return err
			}
			prettyPrint(vms)
			return nil
		},
	},
	{
		Name:        "console",
		Usage:       "Show the console of VM with given ID",
		Description: "Show the console of VM with given ID",
		Action: func(ctx *cli.Context) error {
			vm, err := strconv.Atoi(ctx.Args().First())
			if err != nil || vm < 0 {
				return fmt.Errorf("please specify a valid VM id")
			}
			output, err := hedge.Console(vm)
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		},
	},
}
