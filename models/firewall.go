package models

type FirewallTemplateResponse struct {
	FirewallTemplate FirewallTemplate `json:"firewall_template"`
}

type FirewallTemplate struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	WhitelistHos bool          `json:"whitelist_hos"`
	IsDefault    bool          `json:"is_default"`
	Rules        FirewallRules `json:"rules,omitempty"`
}

type Firewall struct {
	ServerIP     string        `json:"server_ip"`
	ServerNumber int           `json:"server_number"`
	Status       string        `json:"status"`
	WhitelistHos bool          `json:"whitelist_hos"`
	Port         string        `json:"port"`
	Rules        FirewallRules `json:"rules"`
}

type FirewallRules struct {
	Input []FirewallRulesInput `json:"input"`
}

type FirewallRulesInput struct {
	IPVersion string `json:"ip_version"`
	Name      string `json:"name"`
	DstIP     string `json:"dst_ip,omitempty"`
	DstPort   string `json:"dst_port,omitempty"`
	SrcIP     string `json:"src_ip,omitempty"`
	SrcPort   string `json:"src_port,omitempty"`
	Protocol  string `json:"protocol,omitempty"`
	TCPFlags  string `json:"tcp_flags,omitempty"`
	Action    string `json:"action"`
}
