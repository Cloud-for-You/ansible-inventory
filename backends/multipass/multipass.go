package multipass

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Host struct {
	Name					string			`json:"name"`
	Release				string			`json:"release"`
	State					string			`json:"state"`
	IPv4					[]string		`json:"ipv4"`
}

type HostList struct {
  List					[]Host			`json:"list"`
}

func GetHosts() ([]Host, error) {
	cmd := exec.Command("multipass", "list", "--format", "json")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Chyba při spouštění příkazu Multipass: %v", err)
	}

	// Zpracování výstupu a získání jmen serverů
	var hostList HostList
	err = json.Unmarshal(stdout, &hostList)
	if err != nil {
		return nil, fmt.Errorf("Chyba při dekódování JSON: %v", err)
	}

	return hostList.List, nil
}