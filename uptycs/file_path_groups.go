package uptycs

import "errors"

func (T FilePathGroup) GetID() string {
	return T.ID
}

func (T FilePathGroup) GetName() string {
	return T.Name
}

func (T FilePathGroup) KeysToDelete() []string {
	return []string{
		"checkSignature",
	}
}

func (c *Client) CreateFilePathGroup(filePathGroup FilePathGroup) (FilePathGroup, error) {
	// Allow creating with yaragrouprules by name. turn them into id's.
	for ind, ypg := range filePathGroup.YaraGroupRules {
		if len(ypg.Name) > 0 && len(ypg.ID) == 0 {
			_ygr, _ := c.GetYaraGroupRule(YaraGroupRule{
				Name: ypg.Name,
			})
			filePathGroup.YaraGroupRules[ind] = YaraGroupRule{
				ID: _ygr.ID,
			}
		}
	}
	return doCreate(c, filePathGroup, "filePathGroups")
}

func (c *Client) UpdateFilePathGroup(filePathGroup FilePathGroup) (FilePathGroup, error) {
	// Allow creating with yaragrouprules by name. turn them into id's.
	for ind, ypg := range filePathGroup.YaraGroupRules {
		if len(ypg.Name) > 0 && len(ypg.ID) == 0 {
			_ygr, _ := c.GetYaraGroupRule(YaraGroupRule{
				Name: ypg.Name,
			})
			filePathGroup.YaraGroupRules[ind] = YaraGroupRule{
				ID: _ygr.ID,
			}
		}
	}
	return doUpdate(c, filePathGroup, "filePathGroups")
}

func (c *Client) GetFilePathGroups() (FilePathGroups, error) {
	return doGetMany(c, FilePathGroups{}, "filePathGroups")
}

func (c *Client) GetFilePathGroup(filePathGroup FilePathGroup) (FilePathGroup, error) {
	if len(filePathGroup.ID) == 0 {
		all, _ := c.GetFilePathGroups()
		for _, _item := range all.Items {
			if _item.Name == filePathGroup.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, filePathGroup, "filePathGroups")
	}
	return filePathGroup, errors.New("filePathGroup was not found")
}

func (c *Client) DeleteFilePathGroup(filePathGroup FilePathGroup) (FilePathGroup, error) {
	return doDelete(c, filePathGroup, "filePathGroups")
}
