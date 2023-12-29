package hosts

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Host struct {
	Name					string			`json:"name"`
	IPv4					[]string		`json:"ipv4"`
}

type HostList struct {
  List					[]Host			`json:"list"`
}

func GetHosts(filePath string) ([]Host, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Chyba pri cteni souboru: %v", err)
	}
	defer file.Close()

	// Zpracování výstupu a získání jmen serverů
	var hostList HostList
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore lines starting with #
		if strings.HasPrefix(line, "#") {
			continue
		}
		fields := strings.Fields(line)

		if len(fields) < 2 {
			// Skip lines without enough fields (e.g., comments or malformed lines)
			continue
		}

		ip := fields[0]
		hostName := fields[1]

		host := Host{
			Name: hostName,
			IPv4: []string{ip},
		}

		hostList.List = append(hostList.List, host)
	}

	if err := scanner.Err(); err != nil {
		return hostList.List, err
	}

	return hostList.List, nil

}