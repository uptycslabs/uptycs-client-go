package uptycs

import "errors"

func (T FlagProfile) GetID() string {
	return T.ID
}

func (T FlagProfile) GetName() string {
	return T.Name
}

func (T FlagProfile) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateFlagProfile(flagProfile FlagProfile) (FlagProfile, error) {
	return doCreate(c, flagProfile, "flagProfiles", []string{})
}

func (c *Client) UpdateFlagProfile(flagProfile FlagProfile) (FlagProfile, error) {
	return doUpdate(c, flagProfile, "flagProfiles", []string{})
}

func (c *Client) GetFlagProfiles() (FlagProfiles, error) {
	return doGetMany(c, FlagProfiles{}, "flagProfiles")
}

func (c *Client) GetFlagProfile(flagProfile FlagProfile) (FlagProfile, error) {
	if len(flagProfile.ID) == 0 {
		all, _ := c.GetFlagProfiles()
		for _, _item := range all.Items {
			if _item.Name == flagProfile.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, flagProfile, "flagProfiles")
	}
	return flagProfile, errors.New("flagProfile was not found")
}

func (c *Client) DeleteFlagProfile(flagProfile FlagProfile) (FlagProfile, error) {
	return doDelete(c, flagProfile, "flagProfiles")
}
