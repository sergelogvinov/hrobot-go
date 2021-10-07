package client

import (
	"encoding/json"
	"fmt"

	"github.com/nl2go/hrobot-go/models"
)

// func (c *Client) FirewallGetList() ([]models.VSwitch, error) {
// 	url := c.baseURL + "/firewall"
// 	bytes, err := c.doGetRequest(url)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var vswitchs []models.VSwitch
// 	err = json.Unmarshal(bytes, &vswitchs)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return vswitchs, nil
// }

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
