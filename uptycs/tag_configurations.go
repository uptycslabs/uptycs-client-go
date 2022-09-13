package uptycs

import "errors"

func (T TagConfiguration) GetID() string {
	return T.ID
}

func (T TagConfiguration) GetName() string {
	return T.Name
}

func (T TagConfiguration) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateTagConfiguration(tagConfiguration TagConfiguration) (TagConfiguration, error) {
	return doCreate(c, tagConfiguration, "tagConfigurations")
}

func (c *Client) UpdateTagConfiguration(tagConfiguration TagConfiguration) (TagConfiguration, error) {
	return doUpdate(c, tagConfiguration, "tagConfigurations")
}

func (c *Client) GetTagConfigurations() (TagConfigurations, error) {
	return doGetMany(c, TagConfigurations{}, "tagConfigurations")
}

func (c *Client) GetTagConfiguration(tagConfiguration TagConfiguration) (TagConfiguration, error) {
	if len(tagConfiguration.ID) == 0 {
		all, _ := c.GetTagConfigurations()
		for _, _item := range all.Items {
			if _item.Name == tagConfiguration.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, tagConfiguration, "tagConfigurations")
	}
	return tagConfiguration, errors.New("tagConfiguration was not found")
}

func (c *Client) DeleteTagConfiguration(tagConfiguration TagConfiguration) (TagConfiguration, error) {
	return doDelete(c, tagConfiguration, "tagConfigurations")
}
