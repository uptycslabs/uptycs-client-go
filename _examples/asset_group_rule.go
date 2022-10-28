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

	assetGroupRuleByID, _ := c.GetAssetGroupRule(uptycs.AssetGroupRule{
		ID: "d774ac13-ad82-4fb2-8bc9-893a1b957264",
	})
	log.Println("Got AssetGroupRule with Name %s", assetGroupRuleByID.Name)

	// Create a assetGroupRule

	newAssetGroupRule, err := c.CreateAssetGroupRule(uptycs.AssetGroupRule{
		Name:     "marcus_test",
		Query:    "select 'servers' as value from ec2_instance_metadata\nwhere instance_id is not null\nunion\nselect 'servers' as value from gce_instance_metadata\nwhere id is not null",
		Interval: 3600,
		Platform: "all",
		Enabled:  false,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created assetGroupRule '%s' with id '%s'", newAssetGroupRule.Name, newAssetGroupRule.ID))

	// Update a assetGroupRule by by ID

	log.Println(fmt.Sprintf("Attempting to update assetGroupRule with id '%s': '%s' to 'marcus test updated'", newAssetGroupRule.ID, newAssetGroupRule.Name))
	updatedAssetGroupRule, err := c.UpdateAssetGroupRule(uptycs.AssetGroupRule{
		ID:   newAssetGroupRule.ID,
		Name: "marcus_test_updated",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated assetGroupRule '%s' with id '%s'", updatedAssetGroupRule.Name, updatedAssetGroupRule.ID))

	// Delete a assetGroupRule by ID

	_, err = c.DeleteAssetGroupRule(uptycs.AssetGroupRule{
		ID: newAssetGroupRule.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted assetGroupRule '%s' with id '%s'", updatedAssetGroupRule.Name, newAssetGroupRule.ID))

}
