package uptycs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetEventExcludeProfiles() (EventExcludeProfiles, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/eventExcludeProfiles", c.HostURL), nil)
	if err != nil {
		return EventExcludeProfiles{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return EventExcludeProfiles{}, err
	}

	eventExcludeProfiles := EventExcludeProfiles{}
	err = json.Unmarshal(body, &eventExcludeProfiles)
	if err != nil {
		return EventExcludeProfiles{}, err
	}

	return eventExcludeProfiles, nil
}
