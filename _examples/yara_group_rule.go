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

	yaraGroupRuleByName, _ := c.GetYaraGroupRule(uptycs.YaraGroupRule{
		Name: "AmazonAccessKeyId",
	})
	log.Println("Got YaraGroupRule with ID %s", yaraGroupRuleByName.ID)

	yaraGroupRuleByID, _ := c.GetYaraGroupRule(uptycs.YaraGroupRule{
		ID: "c6655aac-abfd-42d4-b2bc-b0a59e98057a",
	})
	log.Println("Got YaraGroupRule with Name %s", yaraGroupRuleByID.Name)

	// Create a yaraGroupRule

	newYaraGroupRule, err := c.CreateYaraGroupRule(uptycs.YaraGroupRule{
		Name:        "marcus test",
		Description: "my test",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created yaraGroupRule '%s' with id '%s'", newYaraGroupRule.Name, newYaraGroupRule.ID))

	// Update a yaraGroupRule by by ID

	log.Println(fmt.Sprintf("Attempting to update yaraGroupRule with id '%s': '%s' to 'marcus test updated'", newYaraGroupRule.ID, newYaraGroupRule.Name))
	updatedYaraGroupRule, err := c.UpdateYaraGroupRule(uptycs.YaraGroupRule{
		ID:          newYaraGroupRule.ID,
		Name:        "marcus test updated",
		Description: "my test",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated yaraGroupRule '%s' with id '%s'", updatedYaraGroupRule.Name, updatedYaraGroupRule.ID))

	// Delete a yaraGroupRule by ID

	_, err = c.DeleteYaraGroupRule(uptycs.YaraGroupRule{
		ID: newYaraGroupRule.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted yaraGroupRule '%s' with id '%s'", updatedYaraGroupRule.Name, newYaraGroupRule.ID))

}
