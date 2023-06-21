package hedge_api

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	MONITOR_ENDPOINT = "/proc/monitor"
	CONSOLE_ENDPOINT = "/proc/vmcons"
)

type VMConfig struct {
	Name    string
	Binary  string
	CPU     int
	Mem     int
	Blk     string
	Net     string
	CmdLine string
}

func (conf VMConfig) Validate() error {
	if conf.Binary == "" {
		return fmt.Errorf("path to kernel cannot be empty")
	}
	if conf.Name == "" {
		return fmt.Errorf("vm name cannot be empty")
	}
	if conf.CPU < 0 {
		return fmt.Errorf("cpu core cannot be negative")
	}
	if conf.Mem <= 0 {
		return fmt.Errorf("memory count must be positive")
	}
	if conf.CmdLine == "" {
		return fmt.Errorf("cmdline cannot be empty")
	}
	return nil
}

type VM struct {
	ID      int
	Name    string
	ModID   int
	ModName string
}

func loadBinary(binary string) error {
	cmd := fmt.Sprintf("load|%s\n", binary)
	err := os.WriteFile(MONITOR_ENDPOINT, []byte(cmd), 0777)
	if err != nil {
		return err
	}
	return nil
}

func Status() error {
	f, err := os.Stat(MONITOR_ENDPOINT)
	if err != nil {
		return err
	}
	if f.IsDir() {
		return fmt.Errorf("%s is a directory", MONITOR_ENDPOINT)
	}
	return nil
}

func StartVM(conf VMConfig) error {
	err := conf.Validate()
	if err != nil {
		return err
	}
	err = loadBinary(conf.Binary)
	if err != nil {
		return err
	}
	cmd := fmt.Sprintf("start|%s|%s|%d|%d|%s|%s|%s\n", conf.Name,
		conf.Binary, conf.CPU, conf.Mem, conf.Blk, conf.Net, conf.CmdLine)
	err = os.WriteFile(MONITOR_ENDPOINT, []byte(cmd), 0777)
	if err != nil {
		return err
	}
	return nil
}

func StopVM(name string) error {
	if name == "" {
		return fmt.Errorf("vm name cannot be empty")
	}
	cmd := fmt.Sprintf("stop|%s\n", name)
	err := os.WriteFile(MONITOR_ENDPOINT, []byte(cmd), 0777)
	if err != nil {
		return err
	}
	return nil
}

func ListVMs() ([]VM, error) {
	content, err := os.ReadFile(MONITOR_ENDPOINT)
	if err != nil {
		return []VM{}, err
	}
	response := string(content)
	response = strings.Trim(response, "\n")
	lines := strings.Split(response, "\n")
	if len(lines) == 0 {
		return []VM{}, fmt.Errorf("%s did not return any value", MONITOR_ENDPOINT)
	}
	if len(lines) == 1 {
		return []VM{}, nil
	}
	lines = lines[1:]

	var vms []VM
	for _, value := range lines {
		parts := strings.Fields(value)
		id := strings.TrimSpace(parts[0])
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return []VM{}, err
		}
		modIdInt, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			return []VM{}, err
		}
		temp := VM{
			ID:      idInt,
			Name:    strings.TrimSpace(parts[1]),
			ModID:   modIdInt,
			ModName: strings.TrimSpace(parts[3]),
		}
		vms = append(vms, temp)
	}
	return vms, nil
}

func Console(id int) (string, error) {
	consoleEndpoint := fmt.Sprintf("%s/vm%d", CONSOLE_ENDPOINT, id)
	content, err := os.ReadFile(consoleEndpoint)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
