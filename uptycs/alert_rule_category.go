package uptycs

import "errors"

func (T AlertRuleCategory) GetID() string {
	return T.ID
}

func (T AlertRuleCategory) GetName() string {
	return T.Name
}

func (T AlertRuleCategory) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateAlertRuleCategory(alertRuleCategory AlertRuleCategory) (AlertRuleCategory, error) {
	return doCreate(c, alertRuleCategory, "alertRuleCategories", []string{})
}

func (c *Client) UpdateAlertRuleCategory(alertRuleCategory AlertRuleCategory) (AlertRuleCategory, error) {
	return doUpdate(c, alertRuleCategory, "alertRuleCategories", []string{})
}

func (c *Client) GetAlertRuleCategories() (AlertRuleCategories, error) {
	return doGetMany(c, AlertRuleCategories{}, "alertRuleCategories")
}

func (c *Client) GetAlertRuleCategory(alertRuleCategory AlertRuleCategory) (AlertRuleCategory, error) {
	if len(alertRuleCategory.ID) == 0 {
		all, _ := c.GetAlertRuleCategories()
		for _, _item := range all.Items {
			if _item.Name == alertRuleCategory.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, alertRuleCategory, "alertRuleCategories")
	}
	return alertRuleCategory, errors.New("alertRuleCategory was not found")
}

func (c *Client) DeleteAlertRuleCategory(alertRuleCategory AlertRuleCategory) (AlertRuleCategory, error) {
	return doDelete(c, alertRuleCategory, "alertRuleCategories")
}
