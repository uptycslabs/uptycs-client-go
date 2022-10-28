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

	carveByID, _ := c.GetCarve(uptycs.Carve{
		ID: "b718f613-055d-43f3-ba58-b665f7633e48",
	})
	log.Println("Got Carve with Path %s", carveByID.Path)

	// Create a carve

	newCarve, err := c.CreateCarve(uptycs.Carve{
		Path: "marcus test",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created carve '%s' with id '%s'", newCarve.Path, newCarve.ID))

	// Update a carve by ID
	log.Println(fmt.Sprintf("Attempting to update carve with id '%s': '%s' to 'marcus test updated'", newCarve.ID, newCarve.Path))
	updatedCarve, err := c.UpdateCarve(uptycs.Carve{
		ID:   newCarve.ID,
		Path: "marcus test updated",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated carve '%s' with id '%s'", updatedCarve.Path, updatedCarve.ID))

	// Delete a carve by ID

	_, err = c.DeleteCarve(uptycs.Carve{
		ID: newCarve.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted carve '%s' with id '%s'", updatedCarve.Path, newCarve.ID))

}
