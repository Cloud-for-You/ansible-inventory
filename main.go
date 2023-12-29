package main

import (
	multipass "Cloud-for-You/ansible-inventory/backends"
	"fmt"
)

func main() {
	hosts, err := multipass.GetHosts()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Println(hosts)
}