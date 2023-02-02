package uptycs

import (
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
		"exceptions", //"exceptions\" is not allowed" api response if included
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
	return eventRule, errors.New("eventRule was not found")
}

func (c *Client) DeleteEventRule(eventRule EventRule) (EventRule, error) {
	return doDelete(c, eventRule, "eventRules")
}

func (c *Client) CreateEventRule(eventRule EventRule) (EventRule, error) {
	if eventRule.BuilderConfig == nil {
		return doCreate(c, eventRule, "eventRules", []string{"builderConfig"})
	}
	return doCreate(c, eventRule, "eventRules", []string{})
}

func (c *Client) UpdateEventRule(eventRule EventRule) (EventRule, error) {
	if eventRule.BuilderConfig == nil {
		return doUpdate(c, eventRule, "eventRules", []string{"builderConfig"})
	}
	return doUpdate(c, eventRule, "eventRules", []string{})
}
