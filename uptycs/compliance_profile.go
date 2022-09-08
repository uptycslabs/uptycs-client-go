package uptycs

import (
	"errors" 
	
)

func (T ComplianceProfile) GetID() string {
	return T.ID
}

func (T ComplianceProfile) GetName() string {
	return T.Name
}

func (T ComplianceProfile) KeysToDelete() []string {
	return []string{
		
	}
}

func (c *Client) GetComplianceProfiles() (ComplianceProfiles, error) {
	return doGetMany(c, ComplianceProfiles{}, "complianceProfiles")
}

func (c *Client) GetComplianceProfile(complianceProfile ComplianceProfile) (ComplianceProfile, error) {
	if len(complianceProfile.GetID()) == 0 {
		all, _ := c.GetComplianceProfiles()
		for _, _item := range all.Items {
			if _item.GetName() == complianceProfile.GetName() {
				return _item, nil
			}
		}
	} else {
		return doGet(c, complianceProfile, "complianceProfiles")
	}
	return complianceProfile, errors.New("Compliance profile was not found")
}