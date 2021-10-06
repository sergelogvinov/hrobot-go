package client

import (
	"encoding/json"
	"fmt"

	"github.com/nl2go/hrobot-go/models"
)

func (c *Client) VSwitchGetList() ([]models.VSwitch, error) {
	url := c.baseURL + "/vswitch"
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var vswitchs []models.VSwitch
	err = json.Unmarshal(bytes, &vswitchs)
	if err != nil {
		return nil, err
	}

	return vswitchs, nil
}

func (c *Client) VSwitchGet(id int) (*models.VSwitch, error) {
	url := fmt.Sprintf(c.baseURL+"/vswitch/%d", id)
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var vswitch models.VSwitch
	err = json.Unmarshal(bytes, &vswitch)
	if err != nil {
		return nil, err
	}

	return &vswitch, nil
}
