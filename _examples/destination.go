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

	destinationByName, _ := c.GetDestination(uptycs.Destination{
		Name: "#reddiconnect-alerts",
	})
	log.Println("Got Destination with ID %s", destinationByName.ID)

	destinationByID, _ := c.GetDestination(uptycs.Destination{
		ID: "b7c9c973-e2a3-4913-a755-919026267679",
	})
	log.Println("Got Destination with Name %s", destinationByID.Name)

	// Create a destination

	newDestination, err := c.CreateDestination(uptycs.Destination{
		Name:    "marcus test",
		Type:    "email",
		Address: "test@email.com",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created destination '%s' with id %s", newDestination.Name, newDestination.ID))

	// Update a destination by by ID

	log.Println(fmt.Sprintf("Attempting to update destination with id %s: '%s' to 'marcus test updated'", newDestination.ID, newDestination.Name))
	updatedDestination, err := c.UpdateDestination(uptycs.Destination{
		ID:   newDestination.ID,
		Name: "marcus test updated",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated destination '%s' with id %s", updatedDestination.Name, updatedDestination.ID))

	// Delete a destination by ID

	_, err = c.DeleteDestination(uptycs.Destination{
		ID: newDestination.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted destination '%s' with id %s", updatedDestination.Name, newDestination.ID))

}
