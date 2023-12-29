package multipass

import (
	"fmt"
	"os/exec"
	"strings"
)

func GetHosts() ([]string, error) {
	cmd := exec.Command("multipass", "list")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Chyba při spouštění příkazu Multipass: %v", err)
	}

	// Zpracování výstupu a získání jmen serverů
	var hosts []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines[2:] {
		fields := strings.Fields(line)
		if len(fields) > 0 {
			hosts = append(hosts, fields[0])
		}
	}

	return hosts, nil
}