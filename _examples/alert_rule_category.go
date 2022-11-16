package main

import (
	"fmt"
	"github.com/uptycslabs/uptycs-client-go/uptycs"
	"log"
	"os"
)

func main() {
	c, _ := uptycs.NewClient(uptycs.Config{
		Host:       os.Getenv("UPTYCS_HOST"),
		APIKey:     os.Getenv("UPTYCS_API_KEY"),
		APISecret:  os.Getenv("UPTYCS_API_SECRET"),
		CustomerID: os.Getenv("UPTYCS_CUSTOMER_ID"),
	})

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
		AlertRuleExceptions: []uptycs.RuleException{},
		Destinations:        []uptycs.AlertRuleDestination{},
		Code:                "test_marc",
		Type:                "sql",
		Rule:                "select * from processes limit 1 :to;",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("Created Rule '%s' with id '%s'", rule.Name, rule.ID))

	newAlertRuleCategory, err := c.CreateAlertRuleCategory(uptycs.AlertRuleCategory{
		Name:   "marcus test",
		RuleID: rule.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created alertRuleCategory '%s' with id '%s'", newAlertRuleCategory.Name, newAlertRuleCategory.ID))

	// Update a alertRuleCategory by ID

	log.Println(fmt.Sprintf("Attempting to update alertRuleCategory with id '%s': '%s' to 'marcus test updated'", newAlertRuleCategory.ID, newAlertRuleCategory.Name))
	updatedAlertRuleCategory, err := c.UpdateAlertRuleCategory(uptycs.AlertRuleCategory{
		ID:   newAlertRuleCategory.ID,
		Name: "marcus test updated",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated alertRuleCategory '%s' with id '%s'", updatedAlertRuleCategory.Name, updatedAlertRuleCategory.ID))

	// Delete a alertRuleCategory by ID

	_, err = c.DeleteAlertRuleCategory(uptycs.AlertRuleCategory{
		ID: newAlertRuleCategory.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted alertRuleCategory '%s' with id '%s'", updatedAlertRuleCategory.Name, newAlertRuleCategory.ID))
	_, err = c.DeleteAlertRule(uptycs.AlertRule{
		ID: rule.ID,
	})
	log.Println(fmt.Sprintf("Deleted Rule '%s' with id '%s'", rule.Name, rule.ID))

}
