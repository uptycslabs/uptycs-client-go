package uptycs

func (T WindowsDefenderPreference) GetID() string {
	return T.ID
}

func (T WindowsDefenderPreference) GetName() string {
	return T.Name
}

func (T WindowsDefenderPreference) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateWindowsDefenderPreference(windowsDefenderPreference WindowsDefenderPreference) (WindowsDefenderPreference, error) {
	panic("No methods are supported for windows defender preferences yet. The API has no reference data.")
	//nolint:govet
	return WindowsDefenderPreference{}, nil
}

func (c *Client) UpdateWindowsDefenderPreference(windowsDefenderPreference WindowsDefenderPreference) (WindowsDefenderPreference, error) {
	panic("No methods are supported for windows defender preferences yet. The API has no reference data.")
	//nolint:govet
	return WindowsDefenderPreference{}, nil
}

func (c *Client) GetWindowsDefenderPreferences() (WindowsDefenderPreferences, error) {
	panic("No methods are supported for windows defender preferences yet. The API has no reference data.")
	//nolint:govet
	return WindowsDefenderPreferences{}, nil
}

func (c *Client) GetWindowsDefenderPreference(windowsDefenderPreference WindowsDefenderPreference) (WindowsDefenderPreference, error) {
	panic("No methods are supported for windows defender preferences yet. The API has no reference data.")
	//nolint:govet
	return WindowsDefenderPreference{}, nil
}

func (c *Client) DeleteWindowsDefenderPreference(windowsDefenderPreference WindowsDefenderPreference) (WindowsDefenderPreference, error) {
	panic("No methods are supported for windows defender preferences yet. The API has no reference data.")
	//nolint:govet
	return WindowsDefenderPreference{}, nil
}
