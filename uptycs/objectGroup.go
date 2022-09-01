package uptycs

import "errors"

func (T ObjectGroup) GetID() string {
	return T.ID
}

func (T ObjectGroup) GetName() string {
	return T.Name
}

func (T ObjectGroup) KeysToDelete() []string {
	return []string{}
}

func (c *Client) GetObjectGroups() (ObjectGroups, error) {
	return doGetMany(c, ObjectGroups{}, "objectGroups")
}

func (c *Client) GetObjectGroup(objectGroup ObjectGroup) (ObjectGroup, error) {
	if len(objectGroup.ID) == 0 {
		all, _ := c.GetObjectGroups()
		for _, _item := range all.Items {
			if _item.Name == objectGroup.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, objectGroup, "objectGroups")
	}
	return objectGroup, errors.New("user was not found")
}

// TODO: Only support GET for now until we fully understand what objectGroups are.

//func (c *Client) DeleteObjectGroup(objectGroup ObjectGroup) (ObjectGroup, error) {
//	return doDelete(c, objectGroup, "objectGroups")
//}
//
//func (c *Client) CreateObjectGroup(objectGroup ObjectGroup) (ObjectGroup, error) {
//	return doCreate(c, objectGroup, "objectGroups")
//}
//
//func (c *Client) UpdateObjectGroup(objectGroup ObjectGroup) (ObjectGroup, error) {
//	return doUpdate(c, objectGroup, "objectGroups")
//}
