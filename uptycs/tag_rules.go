package uptycs

import "errors"

func (T TagRule) GetID() string {
	return T.ID
}

func (T TagRule) GetName() string {
	return T.Name
}

func (T TagRule) KeysToDelete() []string {
	return []string{
		"resourceType",
	}
}

func (c *Client) CreateTagRule(tagRule TagRule) (TagRule, error) {
	return doCreate(c, tagRule, "tagRules")
}

func (c *Client) UpdateTagRule(tagRule TagRule) (TagRule, error) {
	return doUpdate(c, tagRule, "tagRules")
}

func (c *Client) GetTagRules() (TagRules, error) {
	return doGetMany(c, TagRules{}, "tagRules")
}

func (c *Client) GetTagRule(tagRule TagRule) (TagRule, error) {
	if len(tagRule.ID) == 0 {
		all, _ := c.GetTagRules()
		for _, _item := range all.Items {
			if _item.Name == tagRule.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, tagRule, "tagRules")
	}
	return tagRule, errors.New("tagRule was not found")
}

func (c *Client) DeleteTagRule(tagRule TagRule) (TagRule, error) {
	return doDelete(c, tagRule, "tagRules")
}
