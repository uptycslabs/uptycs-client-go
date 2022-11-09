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

	flagProfileByName, _ := c.GetFlagProfile(uptycs.FlagProfile{
		Name: "#examplecorpconnect-alerts",
	})
	log.Println("Got FlagProfile with ID ", flagProfileByName.ID)

	flagProfileByID, _ := c.GetFlagProfile(uptycs.FlagProfile{
		ID: "b7c9c973-e2a3-4913-a755-919026267679",
	})
	log.Println("Got FlagProfile with Name ", flagProfileByID.Name)

	// Create a flagProfile
	newFlagProfile, err := c.CreateFlagProfile(uptycs.FlagProfile{
		Name:        "marcus test",
		Description: "marcus test",
		Priority:    50,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created flagProfile '%s' with id '%s'", newFlagProfile.Name, newFlagProfile.ID))

	// Update a flagProfile by ID
	log.Println(fmt.Sprintf("Attempting to update flagProfile with id '%s': '%s' to 'marcus test updated'", newFlagProfile.ID, newFlagProfile.Name))
	updatedFlagProfile, err := c.UpdateFlagProfile(uptycs.FlagProfile{
		ID:   newFlagProfile.ID,
		Name: "marcus test updated",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated flagProfile '%s' with id '%s'", updatedFlagProfile.Name, updatedFlagProfile.ID))

	// Delete a flagProfile by ID
	_, err = c.DeleteFlagProfile(uptycs.FlagProfile{
		ID: newFlagProfile.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted flagProfile '%s' with id '%s'", updatedFlagProfile.Name, newFlagProfile.ID))

}
