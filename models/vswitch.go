package models

type VSwitch struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Vlan      int    `json:"Vlan"`
	Cancelled bool   `json:"cancelled"`
	Server    []struct {
		ServerIP     string `json:"server_ip"`
		ServerIP6Net string `json:"server_ipv6_net"`
		ServerNumber int    `json:"server_number"`
		Status       string `json:"status"`
	} `json:"server,omitempty"`
	Subnet []struct {
		Ip      string `json:"ip"`
		Mask    int    `json:"mask"`
		Gateway string `json:"gateway"`
	} `json:"subnet,omitempty"`
	CloudNetwork []struct {
		ID      int    `json:"id"`
		Ip      string `json:"ip"`
		Mask    int    `json:"mask"`
		Gateway string `json:"gateway"`
	} `json:"cloud_network,omitempty"`
}
