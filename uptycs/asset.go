package uptycs

import "errors"

func (T Asset) GetID() string {
	return T.ID
}

func (T Asset) GetName() string {
	return T.Name
}

func (T Asset) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateAsset(asset Asset) (Asset, error) {
	// TODO maybe we should support this but it does not currently make sense
	return Asset{}, errors.New("POST is not supported for assets")
}

func (c *Client) UpdateAsset(asset Asset) (Asset, error) {
	// TODO maybe we should support this but it does not currently make sense
	return Asset{}, errors.New("PUT is not supported for assets")
}

func (c *Client) GetAssets() (Assets, error) {
	return doGetMany(c, Assets{}, "assets")
}

func (c *Client) GetAsset(asset Asset) (Asset, error) {
	if len(asset.ID) == 0 {
		all, _ := c.GetAssets()
		for _, _item := range all.Items {
			if _item.Name == asset.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, asset, "assets")
	}
	return asset, errors.New("asset was not found")
}

func (c *Client) DeleteAsset(asset Asset) (Asset, error) {
	// TODO maybe we should support this but it does not currently make sense
	return Asset{}, errors.New("DELETE is not supported for assets")
}
