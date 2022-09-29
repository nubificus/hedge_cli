package hedge_api

import (
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const mon_endpoint = "/proc/monitor"
const cons_endpoint = "/proc/vmcons"

func StartVM(bin_path string, vm_name string, cpu int, mem int,
	blk string, net string, cmdline string) (retError error) {
	defer func() {
		if retError != nil {
			log.WithError(retError).Error(retError.Error())
		}
	}()

	if bin_path == "" {
		return &VMParamError{Param: "kernel", Value: bin_path}
	}
	if vm_name == "" {
		return &VMParamError{Param: "name", Value: vm_name}
	}
	if cpu < 0 {
		return &VMParamError{Param: "core", Value: strconv.Itoa(cpu)}
	}
	if mem < 0 {
		return &VMParamError{Param: "mem", Value: strconv.Itoa(mem)}
	}

	cmd := fmt.Sprintf("load|%s", bin_path)
	err := os.WriteFile(mon_endpoint, []byte(cmd), 0777)
	if err != nil {
		return err
	}
	cmd = fmt.Sprintf("start|%s|%s|%d|%d|%s|%s|%s", vm_name,
		bin_path, cpu, mem, blk, net, cmdline)
	err = os.WriteFile(mon_endpoint, []byte(cmd), 0777)
	return err
}

func StopVM(vm_name string) (retError error) {
	defer func() {
		if retError != nil {
			log.WithError(retError).Error(retError.Error())
		}
	}()

	if vm_name == "" {
		return &VMParamError{Param: "name", Value: vm_name}
	}

	cmd := fmt.Sprintf("stop|%s", vm_name)
	err := os.WriteFile(mon_endpoint, []byte(cmd), 0777)
	return err
}

func ShowVMs() (vms string, retError error) {
	defer func() {
		if retError != nil {
			log.WithError(retError).Error(retError.Error())
		}
	}()

	var content []byte
	content, err := os.ReadFile(mon_endpoint)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func ShowConsole(vm_id int) (console string, retError error) {
	defer func() {
		if retError != nil {
			log.WithError(retError).Error(retError.Error())
		}
	}()
	cmd := fmt.Sprintf("%d", vm_id)
	err := os.WriteFile(cons_endpoint, []byte(cmd), 0777)
	if err != nil {
		return "", err
	}
	// This value of cmd is never used
	cmd = fmt.Sprintf("cat /proc/vmcons\n")
	content, err := os.ReadFile(cons_endpoint)
	if err != nil {
		return "", err
	}
	return string(content), nil

}
