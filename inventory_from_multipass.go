package main

import (
	ansible "Cloud-for-You/ansible-inventory/api"
	multipass "Cloud-for-You/ansible-inventory/backends/multipass"
	"encoding/json"
	"fmt"
)

func main() {
	// Ziskame vsechny hosty z backendu
	hosts, err := multipass.GetHosts()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Inicializace inventory
  jsonInventory := ansible.JSONInventory{}
	// Inicializace mapy pro hostvars
	jsonInventory.Meta.HostVars = make(map[string]ansible.HostVars)
	jsonInventory.All.Children = append(jsonInventory.All.Children, "ungrouped")
	
	// Vsechny ziskane hosty z backendu ulozime do inventory
	for _, host := range hosts {
		// Vlozeni do skupiny ungrouped
		jsonInventory.Ungrouped.Hosts = append(jsonInventory.Ungrouped.Hosts, host.Name)
		// Definujeme potrebne hostvars pro pripojenik hostu
		jsonInventory.Meta.HostVars[host.Name] = ansible.HostVars{
			AnsibleHost: host.IPv4[0],
		}
	}

  // Vypiseme inventory na stdout
	inventoryData, err := json.MarshalIndent(jsonInventory, "", " ")
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(inventoryData))
}