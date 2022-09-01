package main

import (
	"fmt"
	"log"
	"os"

	"github.com/MakeNowJust/heredoc"
	"github.com/uptycslabs/uptycs-client-go/uptycs"
)

func main() {
	c, _ := uptycs.NewClient(uptycs.Config{
		Host:       os.Getenv("UPTYCS_HOST"),
		APIKey:     os.Getenv("UPTYCS_API_KEY"),
		APISecret:  os.Getenv("UPTYCS_API_SECRET"),
		CustomerID: os.Getenv("UPTYCS_CUSTOMER_ID"),
	})
	//Get all event rules

	rules, _ := c.GetEventRules()
	log.Println(len(rules.Items))
	for _, event := range rules.Items {
		if event.Name == "AWS AMI not encrypted for data that is at rest" {
			log.Println(fmt.Sprintf("%s has id %s", event.Name, event.ID))
		}
	}

	// Get an event rule by ID
	eventRule, err := c.GetEventRule(uptycs.EventRule{
		ID: "159dcff5-cee1-46a6-8b1e-93d611e69818",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("Found rule by ID with name %s as %s", eventRule.Name, eventRule.ID))

	rule, err := c.CreateEventRule(uptycs.EventRule{
		Name:        "marc_is_awesomer",
		Description: "this is a test",
		Code:        "1651259159841CODE",
		Type:        "builder",
		Rule:        "builder",
		BuilderConfig: uptycs.BuilderConfig{
			TableName:     "process_open_sockets",
			Added:         true,
			MatchesFilter: true,
			Filters: uptycs.BuilderConfigFilterString(heredoc.Doc(`{
			    "and": [
			      {
			        "not": true,
			        "name": "remote_address",
			        "value": "^172.(1[6-9]|2[0-9]|3[01])|^10.|^192.168.",
			        "operator": "MATCHES_REGEX"
			      }
			    ]
			  }
			`)),
			Severity:        "low",
			Key:             "Test",
			ValueField:      "pid",
			AutoAlertConfig: uptycs.AutoAlertConfig{},
		},
		EventTags: []string{
			"Tactic=Persistence",
			"Version=1.1",
			"Permissions Required=User",
		},
		Grouping:   "builderRules",
		GroupingL2: "Impact",
		GroupingL3: "T1560",
	})

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("Created Rule '%s' with id %s", rule.Name, rule.ID))

	// Delete an event rule by ID
	_, err = c.DeleteEventRule(uptycs.EventRule{
		ID: rule.ID,
	})
	log.Println(fmt.Sprintf("Deleted Rule '%s' with id %s", rule.Name, rule.ID))
}
