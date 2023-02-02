package uptycs

import "errors"

func (T AtcQuery) GetID() string {
	return T.ID
}

func (T AtcQuery) GetName() string {
	return T.Name
}

func (T AtcQuery) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateAtcQuery(atcQuery AtcQuery) (AtcQuery, error) {
	return doCreate(c, atcQuery, "atcQueries", []string{})
}

func (c *Client) UpdateAtcQuery(atcQuery AtcQuery) (AtcQuery, error) {
	return doUpdate(c, atcQuery, "atcQueries", []string{})
}

func (c *Client) GetAtcQueries() (AtcQueries, error) {
	return doGetMany(c, AtcQueries{}, "atcQueries")
}

func (c *Client) GetAtcQuery(atcQuery AtcQuery) (AtcQuery, error) {
	if len(atcQuery.ID) == 0 {
		all, _ := c.GetAtcQueries()
		for _, _item := range all.Items {
			if _item.Name == atcQuery.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, atcQuery, "atcQueries")
	}
	return atcQuery, errors.New("atcQuery was not found")
}

func (c *Client) DeleteAtcQuery(atcQuery AtcQuery) (AtcQuery, error) {
	return doDelete(c, atcQuery, "atcQueries")
}
