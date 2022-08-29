package uptycs

import "errors"

func (T Tag) GetID() string {
	return T.ID
}

func (T Tag) GetName() string {
	return T.Name
}

func (T Tag) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateTag(tag Tag) (Tag, error) {
	return doCreate(c, tag, "tags")
}

func (c *Client) UpdateTag(tag Tag) (Tag, error) {
	return doUpdate(c, tag, "tags")
}

func (c *Client) GetTags() (Tags, error) {
	return doGetMany(c, Tags{}, "tags")
}

func (c *Client) GetTag(tag Tag) (Tag, error) {
	if len(tag.ID) == 0 {
		all, _ := c.GetTags()
		for _, _item := range all.Items {
			if _item.Key == tag.Key && _item.Value == tag.Value {
				return _item, nil
			}
		}
	} else {
		return doGet(c, tag, "tags")
	}
	return tag, errors.New("tag was not found")
}

func (c *Client) DeleteTag(tag Tag) (Tag, error) {
	return doDelete(c, tag, "tags")
}
