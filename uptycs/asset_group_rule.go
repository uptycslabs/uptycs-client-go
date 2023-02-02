package uptycs

import "errors"

func (T AssetGroupRule) GetID() string {
	return T.ID
}

func (T AssetGroupRule) GetName() string {
	return T.Name
}

func (T AssetGroupRule) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateAssetGroupRule(assetGroupRule AssetGroupRule) (AssetGroupRule, error) {
	return doCreate(c, assetGroupRule, "assetGroupRules", []string{})
}

func (c *Client) UpdateAssetGroupRule(assetGroupRule AssetGroupRule) (AssetGroupRule, error) {
	return doUpdate(c, assetGroupRule, "assetGroupRules", []string{})
}

func (c *Client) GetAssetGroupRules() (AssetGroupRules, error) {
	return doGetMany(c, AssetGroupRules{}, "assetGroupRules")
}

func (c *Client) GetAssetGroupRule(assetGroupRule AssetGroupRule) (AssetGroupRule, error) {
	if len(assetGroupRule.ID) == 0 {
		all, _ := c.GetAssetGroupRules()
		for _, _item := range all.Items {
			if _item.Name == assetGroupRule.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, assetGroupRule, "assetGroupRules")
	}
	return assetGroupRule, errors.New("assetGroupRule was not found")
}

func (c *Client) DeleteAssetGroupRule(assetGroupRule AssetGroupRule) (AssetGroupRule, error) {
	return doDelete(c, assetGroupRule, "assetGroupRules")
}
