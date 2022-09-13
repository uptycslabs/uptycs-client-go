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

	tagRuleByName, _ := c.GetTagRule(uptycs.TagRule{
		Name: "AWS Account ID Tag",
	})
	log.Printf("Got TagRule with ID %s\n", tagRuleByName.ID)

	tagRuleByID, _ := c.GetTagRule(uptycs.TagRule{
		ID: "3b231f3a-f852-4d86-8611-2d3a1a368851",
	})
	log.Printf("Got TagRule with Name %s\n", tagRuleByID.Name)

	// Create a tagRule

	newTagRule, err := c.CreateTagRule(uptycs.TagRule{
		Description: "a test",
		Source:      "realtime",
		Platform:    "all",
		Interval:    30,
		Name:        "marcus test",
		Query:       "SELECT upt-mac-edr AS tag FROM system_info WHERE name = 'Mac OS X'",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created tagRule '%s' with id '%s'", newTagRule.Name, newTagRule.ID))

	// Delete a tagRule by ID

	_, err = c.DeleteTagRule(uptycs.TagRule{
		ID: newTagRule.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted tagRule '%s' with id '%s'", newTagRule.Name, newTagRule.ID))

}
