package uptycs

import (
	"errors"
)

func (T AlertRule) GetID() string {
	return T.ID
}

func (T AlertRule) GetName() string {
	return T.Name
}

func (T AlertRule) KeysToDelete() []string {
	keysToDelete := []string{
		"enabled",
		"isInternal",
		"timeSuppresionStart",
		"timeSuppresionDuration",
		"seedId",
		"throttled",
		"links",
	}

	if T.Type != "sql" {
		keysToDelete = append(keysToDelete, "sqlConfig")
	}

	return keysToDelete
}

func (c *Client) GetAlertRules() (AlertRules, error) {
	return doGetMany(c, AlertRules{}, "alertRules")
}

func (c *Client) GetAlertRule(alertRule AlertRule) (AlertRule, error) {
	if len(alertRule.ID) == 0 {
		all, _ := c.GetAlertRules()
		for _, _item := range all.Items {
			if _item.Name == alertRule.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, alertRule, "alertRules")
	}
	return alertRule, errors.New("alertRule was not found")
}

func (c *Client) DeleteAlertRule(alertRule AlertRule) (AlertRule, error) {
	return doDelete(c, alertRule, "alertRules")
}

func (c *Client) CreateAlertRule(alertRule AlertRule) (AlertRule, error) {
	if alertRule.AlertTags == nil {
		// For some reason this sometimes defaults to null
		alertRule.AlertTags = []string{}
	}

	if alertRule.BuilderConfig == nil {
		return doCreate(c, alertRule, "alertRules", []string{"builderConfig"})
	}
	return doCreate(c, alertRule, "alertRules", []string{})
}

func (c *Client) UpdateAlertRule(alertRule AlertRule) (AlertRule, error) {
	if alertRule.AlertTags == nil {
		// For some reason this sometimes defaults to null
		alertRule.AlertTags = []string{}
	}
	if alertRule.BuilderConfig == nil {
		return doUpdate(c, alertRule, "alertRules", []string{"builderConfig"})
	}
	return doUpdate(c, alertRule, "alertRules", []string{})
}
