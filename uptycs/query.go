package uptycs

import "errors"

func (T Query) GetID() string {
	return T.ID
}

func (T Query) GetName() string {
	return T.Name
}

func (T Query) KeysToDelete() []string {
	return []string{
		"interval",
		"removed",
		"runNow",
		"snapshot",
		"verified",
	}
}

func (c *Client) CreateQuery(query Query) (Query, error) {
	return doCreate(c, query, "queries", []string{})
}

func (c *Client) UpdateQuery(query Query) (Query, error) {
	return doUpdate(c, query, "queries", []string{})
}

func (c *Client) GetQueries() (Queries, error) {
	return doGetMany(c, Queries{}, "queries")
}

func (c *Client) GetQuery(query Query) (Query, error) {
	if len(query.ID) == 0 {
		all, _ := c.GetQueries()
		for _, _item := range all.Items {
			if _item.Name == query.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, query, "queries")
	}
	return query, errors.New("query was not found")
}

func (c *Client) DeleteQuery(query Query) (Query, error) {
	return doDelete(c, query, "queries")
}
