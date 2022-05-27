package uptycs

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) GetDestinations() (Destinations, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/destinations", c.HostURL), nil)
	if err != nil {
		return Destinations{}, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Destinations{}, err
	}

	destinations := Destinations{}
	err = json.Unmarshal(body, &destinations)
	if err != nil {
		return Destinations{}, err
	}

	return destinations, nil
}
