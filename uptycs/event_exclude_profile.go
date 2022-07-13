package uptycs

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) GetEventExcludeProfile(eventExcludeProfile EventExcludeProfile) (EventExcludeProfile, error) {
	urlStr := fmt.Sprintf("%s/eventExcludeProfiles/%s", c.HostURL, eventExcludeProfile.ID)
	if len(eventExcludeProfile.ID) == 0 {
		urlStr = fmt.Sprintf("%s/eventExcludeProfiles", c.HostURL)
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return eventExcludeProfile, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return EventExcludeProfile{}, err
	}

	foundEventExcludeProfile := EventExcludeProfile{}

	if len(eventExcludeProfile.ID) == 0 {
		urlStr = fmt.Sprintf("%s/eventExcludeProfiles", c.HostURL)
		eventExcludeProfiles := EventExcludeProfiles{}
		err = json.Unmarshal(body, &eventExcludeProfiles)
		if err != nil {
			return EventExcludeProfile{}, err
		}
		for _, dest := range eventExcludeProfiles.Items {
			if dest.Name == eventExcludeProfile.Name {
				foundEventExcludeProfile = dest
				break
			}
		}
	} else {
		err = json.Unmarshal(body, &foundEventExcludeProfile)
		if err != nil {
			return EventExcludeProfile{}, err
		}
	}

	return foundEventExcludeProfile, nil
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

	slimmedEventExcludeProfile, err := SlimStructAsJsonString(eventExcludeProfile, []string{
		"resourceType",
	})
	if err != nil {
		return eventExcludeProfile, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/eventExcludeProfiles", c.HostURL),
		strings.NewReader(string(slimmedEventExcludeProfile)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return eventExcludeProfile, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return eventExcludeProfile, err
	}

	newEventExcludeProfile := EventExcludeProfile{}
	err = json.Unmarshal(body, &newEventExcludeProfile)
	if err != nil {
		return EventExcludeProfile{}, err
	}

	return newEventExcludeProfile, nil
}

func (c *Client) DeleteEventExcludeProfile(eventExcludeProfile EventExcludeProfile) (EventExcludeProfile, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/eventExcludeProfiles/%s", c.HostURL, eventExcludeProfile.ID), nil)
	if err != nil {
		return eventExcludeProfile, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return eventExcludeProfile, err
	}

	return eventExcludeProfile, nil
}

func (c *Client) UpdateEventExcludeProfile(eventExcludeProfile EventExcludeProfile) (EventExcludeProfile, error) {
	if len(eventExcludeProfile.ID) == 0 {
		return eventExcludeProfile, fmt.Errorf("ID of the eventExcludeProfile is required")
	}

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

	slimmedEventExcludeProfile, err := SlimStructAsJsonString(eventExcludeProfile, []string{
		"resourceType",
	})
	if err != nil {
		return eventExcludeProfile, err
	}

	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/eventExcludeProfiles/%s", c.HostURL, eventExcludeProfile.ID),
		strings.NewReader(string(slimmedEventExcludeProfile)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return eventExcludeProfile, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return eventExcludeProfile, err
	}

	return eventExcludeProfile, nil
}
