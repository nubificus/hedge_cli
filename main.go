package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/nubificus/hedge_cli"
)

func usage() {

	fmt.Println("Usage of Hedge's command line tool:")
	fmt.Printf("%s <command> [<args>]\n\n", os.Args[0])
	fmt.Println("The supported commands are:")
	fmt.Println("\tstart Args\t Start a new VM. Args:")
	fmt.Println("\t\t\t --kernel <path_to_kernel_file> --name <unique_name_of_VM>")
	fmt.Println("\t\t\t [--core=<CPU_core] [--mem=VM_memory_in_MBs]")
	fmt.Println("\t\t\t [--blk=path_to_block_device] [--net=network_interface]")
	fmt.Println("\t\t\t [--cmdline=\"command line argument for the guest kernel\"]")
	fmt.Println("\t--stop <vm_name> Stop the VM, with the given name")
	fmt.Println("\t--show\t\t Show all running VMs")
	fmt.Println("\t--cons <vm_id>\t Show the console of the VM with the given id")
}

func main() {
	var mem int
	var core_id int
	var cons_id int
	var show bool

	/*
	 * Generic flags to determine the Hedge command
	 */
	start_cmd := flag.NewFlagSet("start", flag.ExitOnError)

	/*
	 * Subcommands for start command
	 */
	binary := start_cmd.String("kernel", "", "Path to the kernel")
	start_vm_name := start_cmd.String("name", "", "Name of the new VM")
	start_cmd.IntVar(&core_id, "core", 0, "Id of the core, where the VM will run")
	start_cmd.IntVar(&mem, "mem", 128, "Memory for the guest in MBs. Default is 128 MBs")
	blk_dev := start_cmd.String("blk", "", "Path to block device, if any")
	net_dev := start_cmd.String("net", "", "Name of the tap  device to use, if any")
	cmdline := start_cmd.String("cmdline", "", "Command line argument for the kernel")

	/*
	 * Stop command
	 */
	stop_vm_name := flag.String("stop", "", "Name of the VM to stop")

	/*
	 * Command for console
	 */
	flag.IntVar(&cons_id, "cons", -1, "Id of the VM to show console")

	/*
	 * Show all running VMs command
	 */
	flag.BoolVar(&show, "show", false, "Show the running VMs")

	flag.Usage = usage

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}

	if os.Args[1] == "start" {
		start_cmd.Parse(os.Args[2:])
		ret := hedge_cli.Start_vm(*binary, *start_vm_name, core_id, mem, *blk_dev,
				*net_dev, *cmdline)
		os.Exit(ret)
	}

	flag.Parse()

	if *stop_vm_name != "" {
		ret := hedge_cli.Stop_vm(*stop_vm_name)
		os.Exit(ret)
	}
	if cons_id >= 0 {
		hedge_cli.Show_cons(cons_id)
		os.Exit(0)
	}
	if show {
		hedge_cli.Show_vms()
		os.Exit(0)
	}
	fmt.Printf("Unsupported command %s\n\n", os.Args[1])
	os.Exit(1)
}

