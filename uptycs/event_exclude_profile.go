package uptycs

import (
	"encoding/json"
	"errors"
)

func (T EventExcludeProfile) GetID() string {
	return T.ID
}

func (T EventExcludeProfile) GetName() string {
	return T.Name
}

func (T EventExcludeProfile) KeysToDelete() []string {
	return []string{
		"resourceType",
	}
}

func (c *Client) GetEventExcludeProfiles() (EventExcludeProfiles, error) {
	return doGetMany(c, EventExcludeProfiles{}, "eventExcludeProfiles")
}

func (c *Client) GetEventExcludeProfile(eventExcludeProfile EventExcludeProfile) (EventExcludeProfile, error) {
	if len(eventExcludeProfile.ID) == 0 {
		all, _ := c.GetEventExcludeProfiles()
		for _, _item := range all.Items {
			if _item.Name == eventExcludeProfile.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, eventExcludeProfile, "eventExcludeProfiles")
	}
	return eventExcludeProfile, errors.New("event exclude profile was not found")
}

func (c *Client) DeleteEventExcludeProfile(eventExcludeProfile EventExcludeProfile) (EventExcludeProfile, error) {
	return doDelete(c, eventExcludeProfile, "eventExcludeProfiles")
}

func (c *Client) CreateEventExcludeProfile(eventExcludeProfile EventExcludeProfile) (EventExcludeProfile, error) {
	if len(eventExcludeProfile.MetadataJson) > 0 {
		metadata := EventExcludeProfileMetadata{}
		if err := json.Unmarshal([]byte(eventExcludeProfile.MetadataJson), &metadata); err != nil {
			panic(err)
		}
		eventExcludeProfile.Metadata = metadata
		eventExcludeProfile.MetadataJson = ""
	}

	if eventExcludeProfile.Priority > 999999999 {
		return eventExcludeProfile, errors.New("Priority is too large. Should be less than 999999999")
	}
	return doCreate(c, eventExcludeProfile, "eventExcludeProfiles")
}

func (c *Client) UpdateEventExcludeProfile(eventExcludeProfile EventExcludeProfile) (EventExcludeProfile, error) {
	if len(eventExcludeProfile.MetadataJson) > 0 {
		metadata := EventExcludeProfileMetadata{}
		if err := json.Unmarshal([]byte(eventExcludeProfile.MetadataJson), &metadata); err != nil {
			panic(err)
		}
		eventExcludeProfile.Metadata = metadata
		eventExcludeProfile.MetadataJson = ""
	}

	if eventExcludeProfile.Priority > 999999999 {
		return eventExcludeProfile, errors.New("Priority is too large. Should be less than 999999999")
	}

	return doUpdate(c, eventExcludeProfile, "eventExcludeProfiles")
}
