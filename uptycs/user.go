package uptycs

import "errors"

func (T User) GetID() string {
	return T.ID
}

func (T User) GetName() string {
	return T.Name
}

func (T User) KeysToDelete() []string {
	return []string{
		"priorLogin",
		"rangerId",
		"password",
		"lastUpdatedByUptycs",
		"lastSyncedWithRanger",
		"superAdmin",
	}
}

func (c *Client) GetUsers() (Users, error) {
	return doGetMany(c, Users{}, "users")
}

func (c *Client) GetUser(user User) (User, error) {
	if len(user.ID) == 0 {
		all, _ := c.GetUsers()
		for _, _item := range all.Items {
			if _item.Name == user.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, user, "users")
	}
	return user, errors.New("user was not found")
}

func (c *Client) DeleteUser(user User) (User, error) {
	return doDelete(c, user, "users")
}

func (c *Client) CreateUser(user User) (User, error) {
	return doCreate(c, user, "users")
}

func (c *Client) UpdateUser(user User) (User, error) {
	return doUpdate(c, user, "users")
}
