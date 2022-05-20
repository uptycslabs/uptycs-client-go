package uptycs

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func (c *Client) UpdateEventRule(eventRule EventRule) (EventRule, error) {
	if len(eventRule.ID) == 0 {
		return eventRule, fmt.Errorf("ID of the Event Rule is required")
	}

	rb, err := json.Marshal(eventRule)
	if err != nil {
		return eventRule, err
	}

	var eventRuleInterface interface{}
	if err := json.Unmarshal([]byte(rb), &eventRuleInterface); err != nil {
		panic(err)
	}
	if m, ok := eventRuleInterface.(map[string]interface{}); ok {
		for _, item := range []string{
			"id",
			"customerId",
			"seedId",
			"throttled",
			"createdAt",
			"isInternal",
			"createdBy",
			"updatedAt",
			"updatedBy",
			"sqlConfig",
			"scriptConfig",
			"links",
		} {
			delete(m, item)
		}
	}

	slimmedEventRule, err := json.Marshal(eventRuleInterface)
	if err != nil {
		return eventRule, err
	}
	req, err := http.NewRequest(
		"PUT",
		fmt.Sprintf("%s/eventRules/%s", c.HostURL, eventRule.ID),
		strings.NewReader(string(slimmedEventRule)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return eventRule, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return eventRule, err
	}
	if err != nil {
		return eventRule, err
	}

	return eventRule, nil
}

func (c *Client) DeleteEventRule(eventRule EventRule) (EventRule, error) {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/eventRules/%s", c.HostURL, eventRule.ID), nil)
	if err != nil {
		return eventRule, err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return eventRule, err
	}

	return eventRule, nil
}

func (c *Client) CreateEventRule(eventRule EventRule) (EventRule, error) {

	if len(eventRule.BuilderConfig.FiltersJson) > 0 {
		filters := BuilderConfigFilter{}
		if err := json.Unmarshal([]byte(eventRule.BuilderConfig.FiltersJson), &filters); err != nil {
			panic(err)
		}
		eventRule.BuilderConfig.Filters = filters
		eventRule.BuilderConfig.FiltersJson = ""
	}

	rb, err := json.Marshal(eventRule)
	if err != nil {
		return eventRule, err
	}

	var eventRuleInterface interface{}
	if err := json.Unmarshal([]byte(rb), &eventRuleInterface); err != nil {
		panic(err)
	}
	if m, ok := eventRuleInterface.(map[string]interface{}); ok {
		for _, item := range []string{
			"id",
			"customerId",
			"seedId",
			"throttled",
			"createdAt",
			"isInternal",
			"createdBy",
			"updatedAt",
			"updatedBy",
			"sqlConfig",
			"scriptConfig",
			"links",
		} {
			delete(m, item)
		}
	}

	slimmedEventRule, err := json.Marshal(eventRuleInterface)
	if err != nil {
		return eventRule, err
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf("%s/eventRules", c.HostURL),
		strings.NewReader(string(slimmedEventRule)),
	)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return eventRule, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return eventRule, err
	}

	newEventRule := EventRule{}
	err = json.Unmarshal(body, &newEventRule)
	if err != nil {
		return EventRule{}, err
	}

	return newEventRule, nil
}

func (c *Client) GetEventRule(eventRule EventRule) (EventRule, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/eventRules/%s", c.HostURL, eventRule.ID), nil)
	if err != nil {
		return eventRule, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return EventRule{}, err
	}

	foundEventRule := EventRule{}
	err = json.Unmarshal(body, &foundEventRule)
	if err != nil {
		return EventRule{}, err
	}

	return foundEventRule, nil
}
