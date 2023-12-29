package ansible

// HostVars struct represents the structure of host-specific variables.
type HostVars struct{
	AnsibleHost string `json:"ansible_host,omitempty"`
  AnsiblePort string `json:"ansible_port,omitempty"`	
}

// Group struct represents the structure of a group of hosts.
type Group struct {
	Hosts []string `json:"hosts,omitempty"`
}

// Meta struct represents the structure of the "_meta" section in the JSON.
type Meta struct {
	HostVars map[string]HostVars `json:"hostvars,omitempty"`
}

// All struct represents the structure of the "all" section in the JSON.
type All struct {
	Children []string `json:"children,omitempty"`
}

type Ungrouped struct {
	Hosts []string `json:"hosts,omitempty"`
}

// JSONStructure struct represents the overall structure of the JSON.
type JSONInventory struct {
	Meta      Meta            `json:"_meta,omitempty"`
	All       All             `json:"all,omitempty"`
	Ungrouped Ungrouped       `json:"ungrouped,omitempty"`
}