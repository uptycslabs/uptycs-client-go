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

	exceptionByID, _ := c.GetException(uptycs.Exception{
		ID: "ce67f12a-91d1-4a79-b0ee-a60501c5990b",
	})
	log.Println("Got Exception with Path %s", exceptionByID.Name)

	// Create a exception

	newException, err := c.CreateException(uptycs.Exception{
		Name:        "marcus test",
		Description: "marc test",
		TableName:   "aws_cloudtrail_events",
		Rule:        "{\"and\":[{\"caseInsensitive\":true,\"isDate\":false,\"isVersion\":false,\"isWordMatch\":false,\"name\":\"account_id\",\"not\":false,\"operator\":\"EQUALS\",\"value\":\"11111111111\"}]}",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created exception '%s' with id '%s'", newException.Name, newException.ID))

	// Update a exception by ID
	log.Println(fmt.Sprintf("Attempting to update exception with id '%s': '%s' to 'marcus test updated'", newException.ID, newException.Name))
	updatedException, err := c.UpdateException(uptycs.Exception{
		ID:   newException.ID,
		Name: "new marcus test",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated exception '%s' with id '%s'", updatedException.Name, updatedException.ID))

	// Delete a exception by ID
	_, err = c.DeleteException(uptycs.Exception{
		ID: newException.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted exception '%s' with id '%s'", updatedException.Name, newException.ID))

}
