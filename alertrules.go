package uptycs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetAlertRules - Returns list of AlertRules (no auth required)
//func (c *Client) GetAlertRules() ([]AlertRule, error) {
func (c *Client) GetAlertRules() (AlertRules, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/alertRules", c.HostURL), nil)
	if err != nil {
		return AlertRules{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return AlertRules{}, err
	}

	alertRules := AlertRules{}
	err = json.Unmarshal(body, &alertRules)
	if err != nil {
		return AlertRules{}, err
	}

	return alertRules, nil
}
