package client_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"

	client "github.com/nl2go/hrobot-go"
	"github.com/nl2go/hrobot-go/models"
	. "gopkg.in/check.v1"
)

func (s *ClientSuite) TestFirewallTemplateGetListSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/firewall_template_list.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	templates, err := robotClient.FirewallTemplateGetList()
	c.Assert(err, IsNil)
	c.Assert(len(templates), Equals, 2)
	c.Assert(templates[0].ID, Equals, testFirewallTemplateID)
	c.Assert(templates[1].ID, Equals, testFirewallTemplateID2)
}

func (s *ClientSuite) TestFirewallTemplateGetListServerError(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.FirewallTemplateGetList()
	c.Assert(err, Not(IsNil))
}

func (s *ClientSuite) TestFirewallTemplateGetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/firewall_template_get.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	template, err := robotClient.FirewallTemplateGet(testFirewallTemplateID2)
	c.Assert(err, IsNil)
	c.Assert(template.ID, Equals, testFirewallTemplateID2)
}

func (s *ClientSuite) TestFirewallTemplateGetNotFound(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/firewall_template_get_404.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	_, err := robotClient.FirewallTemplateGet(testFirewallTemplateID)
	c.Assert(err, NotNil)
}

func (s *ClientSuite) TestFirewallTemplateSetSuccess(c *C) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reqContentType := r.Header.Get("Content-Type")
		c.Assert(reqContentType, Equals, "application/x-www-form-urlencoded")

		body, bodyErr := ioutil.ReadAll(r.Body)
		c.Assert(bodyErr, IsNil)
		c.Assert(string(body), Matches, ".+name=MyTemplate.+")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		pwd, pwdErr := os.Getwd()
		c.Assert(pwdErr, IsNil)

		data, readErr := ioutil.ReadFile(fmt.Sprintf("%s/test/response/firewall_template_get.json", pwd))
		c.Assert(readErr, IsNil)

		_, err := w.Write(data)
		c.Assert(err, IsNil)
	}))
	defer ts.Close()

	robotClient := client.NewBasicAuthClient("user", "pass")
	robotClient.SetBaseURL(ts.URL)

	inputRule := models.FirewallRulesInput{
		Name: "Rule1",
	}

	input := &models.FirewallTemplate{
		Name:  "MyTemplate",
		ID:    testFirewallTemplateID2,
		Rules: models.FirewallRules{Input: []models.FirewallRulesInput{inputRule}},
	}

	template, err := robotClient.FirewallTemplateSet(testFirewallTemplateID2, input)
	c.Assert(err, IsNil)
	c.Assert(template.ID, Equals, testFirewallTemplateID2)
}
