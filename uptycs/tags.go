package uptycs

import (
	"errors"
)

func (T Tag) GetID() string {
	return T.ID
}

func (T Tag) GetName() string {
	return T.Name
}

func (T Tag) KeysToDelete() []string {
	return []string{
		"customProfile",
		"system",
	}
}

func (c *Client) CreateTag(tag Tag) (Tag, error) {
	//Allow using a custom profile name instead of an ID if given
	if len(tag.CustomProfile) > 0 && len(tag.CustomProfileID) == 0 {
		_cp, _ := c.GetCustomProfile(CustomProfile{
			Name: tag.CustomProfile,
		})
		tag.CustomProfileID = _cp.ID
		tag.CustomProfile = ""
	}
	//Allow using a flag profile name instead of an ID if given
	if len(tag.FlagProfile) > 0 && len(tag.FlagProfileID) == 0 {
		_cp, _ := c.GetFlagProfile(FlagProfile{
			Name: tag.FlagProfile,
		})
		tag.FlagProfileID = _cp.ID
		tag.FlagProfile = ""
	}
	//Allow using a compliance profile name instead of an ID if given
	if len(tag.ComplianceProfile) > 0 && len(tag.ComplianceProfileID) == 0 {
		_cp, _ := c.GetComplianceProfile(ComplianceProfile{
			Name: tag.ComplianceProfile,
		})
		tag.ComplianceProfileID = _cp.ID
		tag.ComplianceProfile = ""
	}
	//Allow using a process block rule name instead of an ID if given
	if len(tag.ProcessBlockRule) > 0 && len(tag.ProcessBlockRuleID) == 0 {
		_cp, _ := c.GetBlockRule(BlockRule{
			Name: tag.ProcessBlockRule,
		})
		tag.ProcessBlockRuleID = _cp.ID
		tag.ProcessBlockRule = ""
	}
	//Allow using a dns block rule name instead of an ID if given
	if len(tag.DNSBlockRule) > 0 && len(tag.DNSBlockRuleID) == 0 {
		_cp, _ := c.GetBlockRule(BlockRule{
			Name: tag.DNSBlockRule,
		})
		tag.DNSBlockRuleID = _cp.ID
		tag.DNSBlockRule = ""
	}
	//Allow using a windows defender preference name instead of an ID if given
	if len(tag.WindowsDefenderPreference) > 0 && len(tag.WindowsDefenderPreferenceID) == 0 {
		_cp, _ := c.GetWindowsDefenderPreference(WindowsDefenderPreference{
			Name: tag.WindowsDefenderPreference,
		})
		tag.WindowsDefenderPreferenceID = _cp.ID
		tag.WindowsDefenderPreference = ""
	}
	//Allow using a tag rule name instead of an ID if given
	if len(tag.TagRule) > 0 && len(tag.TagRuleID) == 0 {
		_cp, _ := c.GetTagRule(TagRule{
			Name: tag.TagRule,
		})
		tag.TagRuleID = _cp.ID
		tag.TagRule = ""
	}

	for ind, fpg := range tag.FilePathGroups {
		if len(fpg.Name) > 0 && len(fpg.ID) == 0 {
			_fpg, _ := c.GetFilePathGroup(FilePathGroup{
				Name: fpg.Name,
			})
			tag.FilePathGroups[ind] = TagConfigurationObject{
				ID: _fpg.ID,
			}
		}
	}
	for ind, eep := range tag.EventExcludeProfiles {
		if len(eep.Name) > 0 && len(eep.ID) == 0 {
			_eep, _ := c.GetEventExcludeProfile(EventExcludeProfile{
				Name: eep.Name,
			})
			tag.EventExcludeProfiles[ind] = TagConfigurationObject{
				ID: _eep.ID,
			}
		}
	}
	for ind, rp := range tag.RegistryPaths {
		if len(rp.Name) > 0 && len(rp.ID) == 0 {
			_rp, _ := c.GetRegistryPath(RegistryPath{
				Name: rp.Name,
			})
			tag.RegistryPaths[ind] = TagConfigurationObject{
				ID: _rp.ID,
			}
		}
	}
	for ind, qp := range tag.Querypacks {
		if len(qp.Name) > 0 && len(qp.ID) == 0 {
			_qp, _ := c.GetQuerypack(Querypack{
				Name: qp.Name,
			})
			tag.Querypacks[ind] = TagConfigurationObject{
				ID: _qp.ID,
			}
		}
	}
	for ind, ygr := range tag.YaraGroupRules {
		if len(ygr.Name) > 0 && len(ygr.ID) == 0 {
			_ygr, _ := c.GetYaraGroupRule(YaraGroupRule{
				Name: ygr.Name,
			})
			tag.YaraGroupRules[ind] = TagConfigurationObject{
				ID: _ygr.ID,
			}
		}
	}
	for ind, ac := range tag.AuditConfigurations {
		if len(ac.Name) > 0 && len(ac.ID) == 0 {
			_ac, _ := c.GetAuditConfiguration(AuditConfiguration{
				Name: ac.Name,
			})
			tag.AuditConfigurations[ind] = TagConfigurationObject{
				ID: _ac.ID,
			}
		}
	}
	return doCreate(c, tag, "tags")
}

func (c *Client) UpdateTag(tag Tag) (Tag, error) {
	//Allow using a custom profile name instead of an ID if given
	if len(tag.CustomProfile) > 0 && len(tag.CustomProfileID) == 0 {
		_cp, _ := c.GetCustomProfile(CustomProfile{
			Name: tag.CustomProfile,
		})
		tag.CustomProfileID = _cp.ID
		tag.CustomProfile = ""
	}
	//Allow using a flag profile name instead of an ID if given
	if len(tag.FlagProfile) > 0 && len(tag.FlagProfileID) == 0 {
		_cp, _ := c.GetFlagProfile(FlagProfile{
			Name: tag.FlagProfile,
		})
		tag.FlagProfileID = _cp.ID
		tag.FlagProfile = ""
	}
	//Allow using a compliance profile name instead of an ID if given
	if len(tag.ComplianceProfile) > 0 && len(tag.ComplianceProfileID) == 0 {
		_cp, _ := c.GetComplianceProfile(ComplianceProfile{
			Name: tag.ComplianceProfile,
		})
		tag.ComplianceProfileID = _cp.ID
		tag.ComplianceProfile = ""
	}
	//Allow using a process block rule name instead of an ID if given
	if len(tag.ProcessBlockRule) > 0 && len(tag.ProcessBlockRuleID) == 0 {
		_cp, _ := c.GetBlockRule(BlockRule{
			Name: tag.ProcessBlockRule,
		})
		tag.ProcessBlockRuleID = _cp.ID
		tag.ProcessBlockRule = ""
	}
	//Allow using a dns block rule name instead of an ID if given
	if len(tag.DNSBlockRule) > 0 && len(tag.DNSBlockRuleID) == 0 {
		_cp, _ := c.GetBlockRule(BlockRule{
			Name: tag.DNSBlockRule,
		})
		tag.DNSBlockRuleID = _cp.ID
		tag.DNSBlockRule = ""
	}
	//Allow using a windows defender preference name instead of an ID if given
	if len(tag.WindowsDefenderPreference) > 0 && len(tag.WindowsDefenderPreferenceID) == 0 {
		_cp, _ := c.GetWindowsDefenderPreference(WindowsDefenderPreference{
			Name: tag.WindowsDefenderPreference,
		})
		tag.WindowsDefenderPreferenceID = _cp.ID
		tag.WindowsDefenderPreference = ""
	}
	//Allow using a tag rule name instead of an ID if given
	if len(tag.TagRule) > 0 && len(tag.TagRuleID) == 0 {
		_cp, _ := c.GetTagRule(TagRule{
			Name: tag.TagRule,
		})
		tag.TagRuleID = _cp.ID
		tag.TagRule = ""
	}

	for ind, fpg := range tag.FilePathGroups {
		if len(fpg.Name) > 0 && len(fpg.ID) == 0 {
			_fpg, _ := c.GetFilePathGroup(FilePathGroup{
				Name: fpg.Name,
			})
			tag.FilePathGroups[ind] = TagConfigurationObject{
				ID: _fpg.ID,
			}
		}
	}
	for ind, eep := range tag.EventExcludeProfiles {
		if len(eep.Name) > 0 && len(eep.ID) == 0 {
			_eep, _ := c.GetEventExcludeProfile(EventExcludeProfile{
				Name: eep.Name,
			})
			tag.EventExcludeProfiles[ind] = TagConfigurationObject{
				ID: _eep.ID,
			}
		}
	}
	for ind, rp := range tag.RegistryPaths {
		if len(rp.Name) > 0 && len(rp.ID) == 0 {
			_rp, _ := c.GetRegistryPath(RegistryPath{
				Name: rp.Name,
			})
			tag.RegistryPaths[ind] = TagConfigurationObject{
				ID: _rp.ID,
			}
		}
	}
	for ind, qp := range tag.Querypacks {
		if len(qp.Name) > 0 && len(qp.ID) == 0 {
			_qp, _ := c.GetQuerypack(Querypack{
				Name: qp.Name,
			})
			tag.Querypacks[ind] = TagConfigurationObject{
				ID: _qp.ID,
			}
		}
	}
	for ind, ygr := range tag.YaraGroupRules {
		if len(ygr.Name) > 0 && len(ygr.ID) == 0 {
			_ygr, _ := c.GetYaraGroupRule(YaraGroupRule{
				Name: ygr.Name,
			})
			tag.YaraGroupRules[ind] = TagConfigurationObject{
				ID: _ygr.ID,
			}
		}
	}
	for ind, ac := range tag.AuditConfigurations {
		if len(ac.Name) > 0 && len(ac.ID) == 0 {
			_ac, _ := c.GetAuditConfiguration(AuditConfiguration{
				Name: ac.Name,
			})
			tag.AuditConfigurations[ind] = TagConfigurationObject{
				ID: _ac.ID,
			}
		}
	}
	return doUpdate(c, tag, "tags")
}

func (c *Client) GetTags() (Tags, error) {
	return doGetMany(c, Tags{}, "tags")
}

func (c *Client) GetTag(tag Tag) (Tag, error) {
	if len(tag.ID) == 0 {
		all, _ := c.GetTags()
		for _, _item := range all.Items {
			if _item.Key == tag.Key && _item.Value == tag.Value {
				return _item, nil
			}
		}
	} else {
		return doGet(c, tag, "tags")
	}
	return tag, errors.New("tag was not found")
}

func (c *Client) DeleteTag(tag Tag) (Tag, error) {
	return doDelete(c, tag, "tags")
}
