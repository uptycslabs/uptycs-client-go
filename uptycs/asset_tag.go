package uptycs

import "errors"

func (T AssetTag) GetID() string {
	return T.ID
}

func (T AssetTag) GetName() string {
	return T.Name
}

func (T AssetTag) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateAssetTag(assetTag AssetTag) (AssetTag, error) {
	return doCreate(c, assetTag, "assetTags/tags", []string{})
}

func (c *Client) UpdateAssetTag(assetTag AssetTag) (AssetTag, error) {
	return doUpdate(c, assetTag, "assetTags/tags", []string{})
}

func (c *Client) GetAssetTags() (AssetTags, error) {
	return AssetTags{}, errors.New("GET is not supported for assetTags")
}

func (c *Client) GetAssetTag(assetTag AssetTag) (AssetTag, error) {
	return AssetTag{}, errors.New("GET is not supported for assetTags")
}

func (c *Client) DeleteAssetTag(assetTag AssetTag) (AssetTag, error) {
	return doDelete(c, assetTag, "assetTags/tags")
}
