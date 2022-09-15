package uptycs

import "errors"

func (T Role) GetID() string {
	return T.ID
}

func (T Role) GetName() string {
	return T.Name
}

func (T Role) KeysToDelete() []string {
	return []string{
		"UserRole",
	}
}

func (c *Client) GetRoles() (Roles, error) {
	return doGetMany(c, Roles{}, "roles")
}

func (c *Client) GetRole(role Role) (Role, error) {
	if len(role.ID) == 0 {
		all, _ := c.GetRoles()
		for _, _item := range all.Items {
			if _item.Name == role.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, role, "roles")
	}
	return role, errors.New("role was not found")
}

func (c *Client) DeleteRole(role Role) (Role, error) {
	return doDelete(c, role, "roles")
}

func (c *Client) CreateRole(role Role) (Role, error) {
	return doCreate(c, role, "roles")
}

func (c *Client) UpdateRole(role Role) (Role, error) {
	return doUpdate(c, role, "roles")
}
