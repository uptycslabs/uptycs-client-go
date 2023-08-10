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

	// Create all queries
	queries, err := c.GetQueries()
	if err != nil {
		log.Fatal(err)
	}
	query := queries.Items[0]
	log.Println("Got all queries, first with id %s", query.ID)

	// Get a single query
	query, err = c.GetQuery(uptycs.Query{
		ID: query.ID,
	})
	log.Println("Got query by ID with Name %s", query.Name)

	// Create a query
	newQuery, err := c.CreateQuery(uptycs.Query{
		Name:  "marcus test",
		Query: "select * from processes",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created query '%s' with id '%s'", newQuery.Name, newQuery.ID))

	// Update a query by ID
	log.Println(fmt.Sprintf("Attempting to update query with id '%s': '%s' to 'marcus test updated'", newQuery.ID, newQuery.Name))
	updatedQuery, err := c.UpdateQuery(uptycs.Query{
		ID:   newQuery.ID,
		Name: "marcus test updated",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated query '%s' with id '%s'", updatedQuery.Name, updatedQuery.ID))

	// Delete a query by ID
	log.Println(fmt.Sprintf("Attempting to delete query with id '%s': '%s'", updatedQuery.ID, updatedQuery.Name))
	_, err = c.DeleteQuery(uptycs.Query{
		ID: updatedQuery.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted query '%s' with id '%s'", updatedQuery.Name, updatedQuery.ID))
}
