package uptycs

import "errors"

func (T RegistryPath) GetID() string {
	return T.ID
}

func (T RegistryPath) GetName() string {
	return T.Name
}

func (T RegistryPath) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateRegistryPath(registryPath RegistryPath) (RegistryPath, error) {
	return doCreate(c, registryPath, "registryPaths", []string{})
}

func (c *Client) UpdateRegistryPath(registryPath RegistryPath) (RegistryPath, error) {
	return doUpdate(c, registryPath, "registryPaths", []string{})
}

func (c *Client) GetRegistryPaths() (RegistryPaths, error) {
	return doGetMany(c, RegistryPaths{}, "registryPaths")
}

func (c *Client) GetRegistryPath(registryPath RegistryPath) (RegistryPath, error) {
	if len(registryPath.ID) == 0 {
		all, _ := c.GetRegistryPaths()
		for _, _item := range all.Items {
			if _item.Name == registryPath.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, registryPath, "registryPaths")
	}
	return registryPath, errors.New("registryPath was not found")
}

func (c *Client) DeleteRegistryPath(registryPath RegistryPath) (RegistryPath, error) {
	return doDelete(c, registryPath, "registryPaths")
}
