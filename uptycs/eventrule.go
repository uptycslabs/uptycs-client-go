package uptycs

import (
	"encoding/json"
	"errors"
)

func (T EventRule) GetID() string {
	return T.ID
}

func (T EventRule) GetName() string {
	return T.Name
}

func (T EventRule) KeysToDelete() []string {
	return []string{
		"throttled",
		"isInternal",
	}
}

func (c *Client) GetEventRules() (EventRules, error) {
	return doGetMany(c, EventRules{}, "eventRules")
}

func (c *Client) GetEventRule(eventRule EventRule) (EventRule, error) {
	if len(eventRule.ID) == 0 {
		all, _ := c.GetEventRules()
		for _, _item := range all.Items {
			if _item.Name == eventRule.Name {
				return _item, nil
			}
		}
	} else {
		return doGet(c, eventRule, "eventRules")
	}
	return eventRule, errors.New("event rule was not found")
}

func (c *Client) DeleteEventRule(eventRule EventRule) (EventRule, error) {
	return doDelete(c, eventRule, "eventRules")
}

func (c *Client) CreateEventRule(eventRule EventRule) (EventRule, error) {
	if len(eventRule.BuilderConfigJson) > 0 {
		builderConfig := BuilderConfig{}
		if err := json.Unmarshal([]byte(eventRule.BuilderConfigJson), &builderConfig); err != nil {
			panic(err)
		}
		eventRule.BuilderConfig = builderConfig
		eventRule.BuilderConfigJson = ""
	}

	if len(eventRule.BuilderConfig.FiltersJson) > 0 {
		filters := BuilderConfigFilter{}
		if err := json.Unmarshal([]byte(eventRule.BuilderConfig.FiltersJson), &filters); err != nil {
			panic(err)
		}
		eventRule.BuilderConfig.Filters = filters
		eventRule.BuilderConfig.FiltersJson = ""
	}

	return doCreate(c, eventRule, "eventRules")
}

func (c *Client) UpdateEventRule(eventRule EventRule) (EventRule, error) {
	if len(eventRule.BuilderConfigJson) > 0 {
		builderConfig := BuilderConfig{}
		if err := json.Unmarshal([]byte(eventRule.BuilderConfigJson), &builderConfig); err != nil {
			panic(err)
		}
		eventRule.BuilderConfig = builderConfig
		eventRule.BuilderConfigJson = ""
	}

	if len(eventRule.BuilderConfig.FiltersJson) > 0 {
		filters := BuilderConfigFilter{}
		if err := json.Unmarshal([]byte(eventRule.BuilderConfig.FiltersJson), &filters); err != nil {
			panic(err)
		}
		eventRule.BuilderConfig.Filters = filters
		eventRule.BuilderConfig.FiltersJson = ""
	}

	return doUpdate(c, eventRule, "eventRules")
}
