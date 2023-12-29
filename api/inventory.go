package inventory

type Host struct {
	Name					string
	Address				string
}

type Group struct {
	Name					string
	Hosts					[]Host
}

type Inventory struct {
	Hosts					[]Host
}