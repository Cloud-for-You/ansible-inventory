package multipass

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type MultipassHost struct {
	Name					string			`json:"name"`
	Release				string			`json:"release"`
	State					string			`json:"state"`
	IPv4					[]string		`json:"ipv4"`

}

type MultipassHostList struct {
  List					[]MultipassHost			`json:"list"`
}

func GetHosts() ([]MultipassHost, error) {
	cmd := exec.Command("multipass", "list", "--format", "json")
	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Chyba při spouštění příkazu Multipass: %v", err)
	}

	// Zpracování výstupu a získání jmen serverů
	var hostList MultipassHostList
	err = json.Unmarshal(stdout, &hostList)
	if err != nil {
		return nil, fmt.Errorf("Chyba při dekódování JSON: %v", err)
	}

	return hostList.List, nil
}