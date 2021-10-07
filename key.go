package client

import (
	"encoding/json"
	"fmt"
	neturl "net/url"

	"github.com/nl2go/hrobot-go/models"
)

func (c *Client) KeyGetList() ([]models.Key, error) {
	url := c.baseURL + "/key"
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var keys []models.KeyResponse
	err = json.Unmarshal(bytes, &keys)
	if err != nil {
		return nil, err
	}

	var data []models.Key
	for _, key := range keys {
		data = append(data, key.Key)
	}

	return data, nil
}

func (c *Client) KeyGet(id string) (*models.Key, error) {
	url := fmt.Sprintf(c.baseURL+"/key/%s", id)
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var key models.KeyResponse
	err = json.Unmarshal(bytes, &key)
	if err != nil {
		return nil, err
	}

	return &key.Key, nil
}

func (c *Client) KeySet(name, data string) (*models.Key, error) {
	// parts := strings.Fields(input.Data)
	// if len(parts) < 2 {
	// 	return nil, fmt.Errorf("bad sshkey")
	// }

	// key, err := base64.StdEncoding.DecodeString(parts[1])
	// if err != nil {
	// 	return nil, fmt.Errorf("bad sshkey")
	// }

	// fingerprint := ""
	// fingerprintMD5 := md5.Sum([]byte(key))

	// for i, b := range fingerprintMD5 {
	// 	fingerprint = fingerprint + fmt.Sprintf("%02x", b)
	// 	if i < len(fingerprintMD5)-1 {
	// 		fingerprint = fingerprint + ":"
	// 	}
	// }

	url := c.baseURL + "/key"
	formData := neturl.Values{
		"name": {name},
		"data": {data},
	}

	bytes, err := c.doPostFormRequest(url, formData)
	if err != nil {
		return nil, err
	}

	var keyResp models.KeyResponse
	err = json.Unmarshal(bytes, &keyResp)
	if err != nil {
		return nil, err
	}

	return &keyResp.Key, nil
}
