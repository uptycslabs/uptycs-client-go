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

	allCarves, _ := c.GetCarves()
	if len(allCarves.Items) > 0 {
		log.Println("Found a Carve %s with Path %s", allCarves.Items[0].Name, allCarves.Items[0].Path)
	}

	// Create a carve
	newCarve, err := c.CreateCarve(uptycs.Carve{
		Path: "marcus test",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("Created carve '%s' with id '%s'", newCarve.Path, newCarve.ID))

	// Update a carve by ID
	log.Println(fmt.Sprintf("Attempting to update carve with id '%s': '%s' to 'marcus test updated'", newCarve.ID, newCarve.Path))
	updatedCarve, err := c.UpdateCarve(uptycs.Carve{
		ID:   newCarve.ID,
		Path: "marcus test updated",
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("Updated carve '%s' with id '%s'", updatedCarve.Path, updatedCarve.ID))

	// Delete a carve by ID
	_, err = c.DeleteCarve(uptycs.Carve{
		ID: newCarve.ID,
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(fmt.Sprintf("Deleted carve '%s' with id '%s'", updatedCarve.Path, newCarve.ID))

}
