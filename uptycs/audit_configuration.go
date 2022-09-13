package uptycs

import "errors"

func (T AuditConfiguration) GetID() string {
	return T.ID
}

func (T AuditConfiguration) GetName() string {
	return T.Name
}

func (T AuditConfiguration) KeysToDelete() []string {
	return []string{}
}

func (c *Client) GetAuditConfigurations() (AuditConfigurations, error) {
	return doGetMany(c, AuditConfigurations{}, "auditConfigurations")
}

func (c *Client) GetAuditConfiguration(auditConfiguration AuditConfiguration) (AuditConfiguration, error) {
	if len(auditConfiguration.ID) == 0 {
		all, _ := c.GetAuditConfigurations()
		for _, _item := range all.Items {
			if _item.Name == auditConfiguration.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, auditConfiguration, "auditConfigurations")
	}
	return auditConfiguration, errors.New("auditConfiguration was not found")
}
