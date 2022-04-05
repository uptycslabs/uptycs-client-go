package uptycs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateAlertRule(alertRule AlertRule) (AlertRule, error) {
	rb, err := json.Marshal(alertRule)
	if err != nil {
		return alertRule, err
	}

	var alertRuleInterface interface{}
	if err := json.Unmarshal([]byte(rb), &alertRuleInterface); err != nil {
		panic(err)
	}
	if m, ok := alertRuleInterface.(map[string]interface{}); ok {
		for _, item := range []string{
			"id",
			"customerId",
			"seedId",
			"alertRuleQueries",
			"enabled",
			"custom",
			"throttled",
			"createdAt",
			"isInternal",
			"createdBy",
			"updatedAt",
			"timeSuppresionStart",
			"timeSuppresionDuration",
			"updatedBy",
			"lock",
			"alertTags",
			"alertNotifyInterval",
			"alertNotifyCount",
			"destinations",
			"script_config",
			"alertRuleQueries",
			"links",
		} {
			delete(m, item)
		}
	}

	slimmedAlertRule, err := json.Marshal(alertRuleInterface)
	if err != nil {
		return alertRule, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/alertRules", c.HostURL),
		strings.NewReader(string(slimmedAlertRule)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return alertRule, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return alertRule, err
	}

	newAlertRule := AlertRule{}
	err = json.Unmarshal(body, &newAlertRule)
	if err != nil {
		return AlertRule{}, err
	}

	return newAlertRule, nil
}

func (c *Client) UpdateAlertRule(alertRule AlertRule) (AlertRule, error) {
	if len(alertRule.ID) == 0 {
		return alertRule, fmt.Errorf("ID of the Alert Rule is required")
	}

	rb, err := json.Marshal(alertRule)
	if err != nil {
		return alertRule, err
	}

	var alertRuleInterface interface{}
	if err := json.Unmarshal([]byte(rb), &alertRuleInterface); err != nil {
		panic(err)
	}
	if m, ok := alertRuleInterface.(map[string]interface{}); ok {
		for _, item := range []string{
			"id",
			"customerId",
			"seedId",
			"type",
			"alertRuleExceptions",
			"alertRuleQueries",
			"enabled",
			"custom",
			"throttled",
			"createdAt",
			"isInternal",
			"createdBy",
			"updatedAt",
			"timeSuppresionStart",
			"timeSuppresionDuration",
			"updatedBy",
			"lock",
			"alertTags",
			"alertNotifyInterval",
			"alertNotifyCount",
			"destinations",
			"script_config",
			"alertRuleQueries",
			"links",
		} {
			delete(m, item)
		}
	}

	slimmedAlertRule, err := json.Marshal(alertRuleInterface)
	if err != nil {
		return alertRule, err
	}
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/alertRules/%s", c.HostURL, alertRule.ID),
		strings.NewReader(string(slimmedAlertRule)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return alertRule, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return alertRule, err
	}
	if err != nil {
		return alertRule, err
	}

	return alertRule, nil
}

func (c *Client) GetAlertRule(alertRule AlertRule) (AlertRule, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/alertRules/%s", c.HostURL, alertRule.ID), nil)
	if err != nil {
		return alertRule, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return AlertRule{}, err
	}

	foundAlertRule := AlertRule{}
	err = json.Unmarshal(body, &foundAlertRule)
	if err != nil {
		return AlertRule{}, err
	}

	return foundAlertRule, nil
}

func (c *Client) DeleteAlertRule(alertRule AlertRule) (AlertRule, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/alertRules/%s", c.HostURL, alertRule.ID), nil)
	if err != nil {
		return alertRule, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return alertRule, err
	}

	return alertRule, nil
}
