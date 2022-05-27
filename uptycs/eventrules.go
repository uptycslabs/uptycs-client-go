package uptycs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetEventRules() (EventRules, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/eventRules", c.HostURL), nil)
	if err != nil {
		return EventRules{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return EventRules{}, err
	}

	eventRules := EventRules{}

	err = json.Unmarshal(body, &eventRules)
	if err != nil {
		return EventRules{}, err
	}

	return eventRules, nil
}
