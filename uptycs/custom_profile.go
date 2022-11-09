package uptycs

import "errors"

func (T CustomProfile) GetID() string {
	return T.ID
}

func (T CustomProfile) GetName() string {
	return T.Name
}

func (T CustomProfile) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateCustomProfile(customProfile CustomProfile) (CustomProfile, error) {
	return doCreate(c, customProfile, "customProfiles")
}

func (c *Client) UpdateCustomProfile(customProfile CustomProfile) (CustomProfile, error) {
	return doUpdate(c, customProfile, "customProfiles")
}

func (c *Client) GetCustomProfiles() (CustomProfiles, error) {
	return doGetMany(c, CustomProfiles{}, "customProfiles")
}

func (c *Client) GetCustomProfile(customProfile CustomProfile) (CustomProfile, error) {
	if len(customProfile.ID) == 0 {
		all, _ := c.GetCustomProfiles()
		for _, _item := range all.Items {
			if _item.Name == customProfile.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, customProfile, "customProfiles")
	}
	return customProfile, errors.New("customProfile was not found")
}

func (c *Client) DeleteCustomProfile(customProfile CustomProfile) (CustomProfile, error) {
	return doDelete(c, customProfile, "customProfiles")
}
