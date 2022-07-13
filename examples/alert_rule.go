package main

import (
	"fmt"
	"github.com/uptycslabs/uptycs-client-go/uptycs"
	"log"
	"os"
)

func main() {
	c, _ := uptycs.NewClient(uptycs.UptycsConfig{
		Host:       os.Getenv("UPTYCS_HOST"),
		ApiKey:     os.Getenv("UPTYCS_API_KEY"),
		ApiSecret:  os.Getenv("UPTYCS_API_SECRET"),
		CustomerID: os.Getenv("UPTYCS_CUSTOMER_ID"),
	})
	// Get all alert rules
	rules, _ := c.GetAlertRules()
	for _, alert := range rules.Items {
		if alert.Name == "marcus test" {
			log.Println(fmt.Sprintf("%s has id %s", alert.Name, alert.ID))
		}
	}

	// Get an alert rule by ID

	rule, _ := c.GetAlertRule(uptycs.AlertRule{
		ID: "1d4720ce-19a9-4a03-bb3a-905717b8a60f",
	})
	log.Println(fmt.Sprintf("Found rule by ID with name %s as %s", rule.Name, rule.ID))

	// Get an alert rule by Name

	anotherRule, _ := c.GetAlertRule(uptycs.AlertRule{
		Name: "7zip.exe execution detected - T1560.001 - Archive via Utility - Windows",
	})
	log.Println(fmt.Sprintf("Found rule by Name with name %s as %s", anotherRule.Name, anotherRule.ID))

	// Create an alert rule

	rule, err := c.CreateAlertRule(uptycs.AlertRule{
		Name:        "marcus test",
		Description: "marcus test",
		Grouping:    "MITRE",
		GroupingL2:  "Impact",
		GroupingL3:  "T1560",
		SQLConfig: &uptycs.SQLConfig{
			IntervalSeconds: 3600,
		},
		Code: "test_marc",
		Type: "sql",
		Rule: "select * from processes limit 1 :to;",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("Created Rule '%s' with id %s", rule.Name, rule.ID))

	// Update an alert rule by by ID

	updatedRule, err := c.UpdateAlertRule(uptycs.AlertRule{
		ID:          rule.ID,
		Name:        "marcus test updated",
		Description: "marcus test updated",
		Grouping:    "MITRE",
		GroupingL2:  "Impact",
		GroupingL3:  "T1580",
		SQLConfig: &uptycs.SQLConfig{
			IntervalSeconds: 1800,
		},
		Code: "test_marc2",
		Type: "sql",
		Rule: "select * from processes limit 2 :to;",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("Updated Rule '%s' with id %s", updatedRule.Name, updatedRule.ID))

	// Delete an alert rule by ID

	_, err = c.DeleteAlertRule(uptycs.AlertRule{
		ID: rule.ID,
	})
	log.Println(fmt.Sprintf("Deleted Rule '%s' with id %s", rule.Name, rule.ID))

}
