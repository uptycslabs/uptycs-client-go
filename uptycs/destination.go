package uptycs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) CreateDestination(destination Destination) (Destination, error) {
	rb, err := json.Marshal(destination)
	if err != nil {
		return destination, err
	}

	var destinationInterface interface{}
	if err := json.Unmarshal([]byte(rb), &destinationInterface); err != nil {
		panic(err)
	}
	if m, ok := destinationInterface.(map[string]interface{}); ok {
		for _, item := range []string{
			"id",
			"customerId",
			"createdAt",
			"createdBy",
			"updatedAt",
			"updatedBy",
		} {
			delete(m, item)
		}
	}

	slimmedDestination, err := json.Marshal(destinationInterface)
	if err != nil {
		return destination, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/destinations", c.HostURL),
		strings.NewReader(string(slimmedDestination)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return destination, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return destination, err
	}

	newDestination := Destination{}
	err = json.Unmarshal(body, &newDestination)
	if err != nil {
		return Destination{}, err
	}

	return newDestination, nil
}

func (c *Client) UpdateDestination(destination Destination) (Destination, error) {
	if len(destination.ID) == 0 {
		return destination, fmt.Errorf("ID of the destination is required")
	}

	rb, err := json.Marshal(destination)
	if err != nil {
		return destination, err
	}

	var destinationInterface interface{}
	if err := json.Unmarshal([]byte(rb), &destinationInterface); err != nil {
		panic(err)
	}
	if m, ok := destinationInterface.(map[string]interface{}); ok {
		for _, item := range []string{
			"id",
			"customerId",
			"createdAt",
			"createdBy",
			"updatedAt",
			"updatedBy",
			"links",
		} {
			delete(m, item)
		}
	}

	slimmedDestination, err := json.Marshal(destinationInterface)
	if err != nil {
		return destination, err
	}
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/destinations/%s", c.HostURL, destination.ID),
		strings.NewReader(string(slimmedDestination)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return destination, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return destination, err
	}
	if err != nil {
		return destination, err
	}

	return destination, nil
}

func (c *Client) GetDestination(destination Destination) (Destination, error) {
	urlStr := fmt.Sprintf("%s/destinations/%s", c.HostURL, destination.ID)
	if len(destination.ID) == 0 {
		urlStr = fmt.Sprintf("%s/destinations", c.HostURL)
	}

	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		return destination, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return Destination{}, err
	}

	foundDestination := Destination{}

	if len(destination.ID) == 0 {
		urlStr = fmt.Sprintf("%s/destinations", c.HostURL)
		destinations := Destinations{}
		err = json.Unmarshal(body, &destinations)
		if err != nil {
			return Destination{}, err
		}
		for _, dest := range destinations.Items {
			if dest.Name == destination.Name {
				foundDestination = dest
				break
			}
		}
	} else {
		err = json.Unmarshal(body, &foundDestination)
		if err != nil {
			return Destination{}, err
		}
	}

	return foundDestination, nil
}

func (c *Client) DeleteDestination(destination Destination) (Destination, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/destinations/%s", c.HostURL, destination.ID), nil)
	if err != nil {
		return destination, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return destination, err
	}

	return destination, nil
}
