package uptycs

import "errors"

func (T YaraGroupRule) GetID() string {
	return T.ID
}

func (T YaraGroupRule) GetName() string {
	return T.Name
}

func (T YaraGroupRule) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateYaraGroupRule(yaraGroupRule YaraGroupRule) (YaraGroupRule, error) {
	return doCreate(c, yaraGroupRule, "yaraGroupRules", []string{})
}

func (c *Client) UpdateYaraGroupRule(yaraGroupRule YaraGroupRule) (YaraGroupRule, error) {
	return doUpdate(c, yaraGroupRule, "yaraGroupRules", []string{})
}

func (c *Client) GetYaraGroupRules() (YaraGroupRules, error) {
	return doGetMany(c, YaraGroupRules{}, "yaraGroupRules")
}

func (c *Client) GetYaraGroupRule(yaraGroupRule YaraGroupRule) (YaraGroupRule, error) {
	if len(yaraGroupRule.ID) == 0 {
		all, _ := c.GetYaraGroupRules()
		for _, _item := range all.Items {
			if _item.Name == yaraGroupRule.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, yaraGroupRule, "yaraGroupRules")
	}
	return yaraGroupRule, errors.New("yaraGroupRule was not found")
}

func (c *Client) DeleteYaraGroupRule(yaraGroupRule YaraGroupRule) (YaraGroupRule, error) {
	return doDelete(c, yaraGroupRule, "yaraGroupRules")
}
