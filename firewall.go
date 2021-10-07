package client

import (
	"encoding/json"
	"fmt"
	neturl "net/url"
	"strconv"

	"github.com/nl2go/hrobot-go/models"
)

func (c *Client) FirewallGet(id int) (*models.Firewall, error) {
	url := fmt.Sprintf(c.baseURL+"/firewall/%d", id)
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var firewall models.Firewall
	err = json.Unmarshal(bytes, &firewall)
	if err != nil {
		return nil, err
	}

	return &firewall, nil
}

func (c *Client) FirewallTemplateGetList() ([]models.FirewallTemplate, error) {
	url := fmt.Sprintf(c.baseURL + "/firewall/template")
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var templates []models.FirewallTemplateResponse
	err = json.Unmarshal(bytes, &templates)
	if err != nil {
		return nil, err
	}

	var data []models.FirewallTemplate
	for _, firewall := range templates {
		data = append(data, firewall.FirewallTemplate)
	}

	return data, nil
}

func (c *Client) FirewallTemplateGet(id int) (*models.FirewallTemplate, error) {
	url := fmt.Sprintf(c.baseURL+"/firewall/template/%d", id)
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var templateResp models.FirewallTemplateResponse
	err = json.Unmarshal(bytes, &templateResp)
	if err != nil {
		return nil, err
	}

	return &templateResp.FirewallTemplate, nil
}

func (c *Client) FirewallTemplateSet(id int, input *models.FirewallTemplate) (*models.FirewallTemplate, error) {
	url := fmt.Sprintf(c.baseURL+"/firewall/template/%d", id)

	formData := neturl.Values{
		"name":          {input.Name},
		"whitelist_hos": {strconv.FormatBool(input.WhitelistHos)},
		"is_default":    {strconv.FormatBool(input.IsDefault)},
	}
	for idx, rule := range input.Rules.Input {
		formData.Set(fmt.Sprintf("rules[input][%d][%s]", idx, "ip_version"), "ipv4")
		formData.Set(fmt.Sprintf("rules[input][%d][%s]", idx, "name"), rule.Name)
		formData.Set(fmt.Sprintf("rules[input][%d][%s]", idx, "dst_ip"), rule.DstIP)
		formData.Set(fmt.Sprintf("rules[input][%d][%s]", idx, "dst_port"), rule.DstPort)
		formData.Set(fmt.Sprintf("rules[input][%d][%s]", idx, "src_ip"), rule.SrcIP)
		formData.Set(fmt.Sprintf("rules[input][%d][%s]", idx, "src_port"), rule.SrcPort)
		if rule.Protocol != "" {
			formData.Set(fmt.Sprintf("rules[input][%d][%s]", idx, "protocol"), rule.Protocol)
		}
		formData.Set(fmt.Sprintf("rules[input][%d][%s]", idx, "action"), rule.Action)
	}

	bytes, err := c.doPostFormRequest(url, formData)
	if err != nil {
		return nil, err
	}

	var firewallResp models.FirewallTemplateResponse
	err = json.Unmarshal(bytes, &firewallResp)
	if err != nil {
		return nil, err
	}

	return &firewallResp.FirewallTemplate, nil
}
