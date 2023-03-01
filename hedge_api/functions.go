package hedge_api

import (
	"fmt"
	"os"
)

const mon_endpoint = "/proc/monitor"
const cons_endpoint = "/proc/vmcons"

func Start_vm(bin_path string, vm_name string, cpu int, mem int,
	blk string, net string, cmdline string) int {

	var ret int

	ret = 0
	if bin_path == "" {
		fmt.Println("Path to kernel was not specified. Please use --kernel argument")
		ret = -1
	}
	if vm_name == "" {
		fmt.Println("Name of VM was not specified. Please use the --name argument")
		ret = -1
	}
	if cpu < 0 {
		fmt.Println("CPU core (--core) can not be negative")
		ret = -1
	}
	if mem < 0 {
		fmt.Println("VM memory (--mem) can not be a negative value")
		ret = -1
	}
	if ret == 0 {
		cmd := fmt.Sprintf("load|%s\n", bin_path)
		err := os.WriteFile(mon_endpoint, []byte(cmd), 0777)
		if err != nil {
			fmt.Println(err)
			ret = 1
		}
		cmd = fmt.Sprintf("start|%s|%s|%d|%d|%s|%s|%s\n", vm_name,
			bin_path, cpu, mem, blk, net, cmdline)
		err = os.WriteFile(mon_endpoint, []byte(cmd), 0777)
		if err != nil {
			fmt.Println(err)
			ret = 1
		}
	}
	return ret
}

func Stop_vm(vm_name string) int {
	var ret int

	ret = 0
	if vm_name == "" {
		return 1
	}

	cmd := fmt.Sprintf("stop|%s\n", vm_name)
	err := os.WriteFile(mon_endpoint, []byte(cmd), 0777)
	if err != nil {
		fmt.Println(err)
		ret = 1
	}
	return ret
}

func Show_vms() {

	var content []byte

	content, err := os.ReadFile(mon_endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}
}

func Show_cons(vm_id int) {
	cmd := fmt.Sprintf("%d", vm_id)
	err := os.WriteFile(cons_endpoint, []byte(cmd), 0777)
	if err != nil {
		fmt.Println(err)
	}
	// This value of cmd is never used
	cmd = fmt.Sprintf("cat /proc/vmcons\n")
	content, err := os.ReadFile(cons_endpoint)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(content))
	}
}
