package uptycs

func (T BlockRule) GetID() string {
	return T.ID
}

func (T BlockRule) GetName() string {
	return T.Name
}

func (T BlockRule) KeysToDelete() []string {
	return []string{}
}

func (c *Client) CreateBlockRule(blockRule BlockRule) (BlockRule, error) {
	panic("No methods are supported for block rules yet. The API returns a 404")
	//nolint:govet
	return BlockRule{}, nil
}

func (c *Client) UpdateBlockRule(blockRule BlockRule) (BlockRule, error) {
	panic("No methods are supported for block rules yet. The API returns a 404")
	//nolint:govet
	return BlockRule{}, nil
}

func (c *Client) GetBlockRules() (BlockRules, error) {
	panic("No methods are supported for block rules yet. The API returns a 404")
	//nolint:govet
	return BlockRules{}, nil
}

func (c *Client) GetBlockRule(blockRule BlockRule) (BlockRule, error) {
	panic("No methods are supported for block rules yet. The API returns a 404")
	//nolint:govet
	return BlockRule{}, nil
}

func (c *Client) DeleteBlockRule(blockRule BlockRule) (BlockRule, error) {
	panic("No methods are supported for block rules yet. The API returns a 404")
	//nolint:govet
	return BlockRule{}, nil
}
