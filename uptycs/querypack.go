package uptycs

import "errors"

func (T Querypack) GetID() string {
	return T.ID
}

func (T Querypack) GetName() string {
	return T.Name
}

func (T Querypack) KeysToDelete() []string {
	return []string{
		"isInternal",
		"queries",
		"resourceType",
	}
}

func (c *Client) CreateQuerypack(querypack Querypack) (Querypack, error) {
	return doCreate(c, querypack, "querypacks", []string{})
}

func (c *Client) UpdateQuerypack(querypack Querypack) (Querypack, error) {
	return doUpdate(c, querypack, "querypacks", []string{})
}

func (c *Client) GetQuerypacks() (Querypacks, error) {
	return doGetMany(c, Querypacks{}, "querypacks")
}

func (c *Client) GetQuerypack(querypack Querypack) (Querypack, error) {
	if len(querypack.ID) == 0 {
		all, _ := c.GetQuerypacks()
		for _, _item := range all.Items {
			if _item.Name == querypack.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, querypack, "querypacks")
	}
	return querypack, errors.New("querypack was not found")
}

func (c *Client) DeleteQuerypack(querypack Querypack) (Querypack, error) {
	return doDelete(c, querypack, "querypacks")
}
