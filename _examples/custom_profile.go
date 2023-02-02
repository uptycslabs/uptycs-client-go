package main

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
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

	customProfileByID, _ := c.GetCustomProfile(uptycs.CustomProfile{
		ID: "c6815103-33eb-41e0-bc2f-6a23cc2e1589",
	})
	log.Println("Got CustomProfile with Name ", customProfileByID.Name)

	// Create a customProfile
	newCustomProfile, err := c.CreateCustomProfile(uptycs.CustomProfile{
		Name:        "marcus test",
		Description: "Test",
		QuerySchedules: uptycs.CustomJSONString(heredoc.Doc(`{
		  "processes": 100
		}`)),
		Priority:     133715,
		ResourceType: "asset",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created customProfile '%s' with id '%s'", newCustomProfile.Name, newCustomProfile.ID))

	// Update a customProfile by ID
	log.Println(fmt.Sprintf("Attempting to update customProfile with id '%s': '%s' to 'marcus test updated'", newCustomProfile.ID, newCustomProfile.Name))
	updatedCustomProfile, err := c.UpdateCustomProfile(uptycs.CustomProfile{
		ID:   newCustomProfile.ID,
		Name: "marcus test updated",
		QuerySchedules: uptycs.CustomJSONString(heredoc.Doc(`{
          "processes": 101
        }`)),
		Priority:     2,
		ResourceType: "asset",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated customProfile '%s' with id '%s'", updatedCustomProfile.Name, updatedCustomProfile.ID))

	// Delete a customProfile by ID
	_, err = c.DeleteCustomProfile(uptycs.CustomProfile{
		ID: newCustomProfile.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted customProfile '%s' with id '%s'", updatedCustomProfile.Name, newCustomProfile.ID))

}
