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

	atcQueryByName, _ := c.GetAtcQuery(uptycs.AtcQuery{
		Name: "atc_chrome_downloads",
	})
	log.Println("Got AtcQuery with ID ", atcQueryByName.ID)

	atcQueryByID, _ := c.GetAtcQuery(uptycs.AtcQuery{
		ID: "dc0e9652-ec9a-4baa-9da3-8547333b3628",
	})
	log.Println("Got AtcQuery with Name ", atcQueryByID.Name)

	// Create a atcQuery
	newAtcQuery, err := c.CreateAtcQuery(uptycs.AtcQuery{
		Name:        "marcus_test",
		Description: "Chrome Browser History",
		Query:       "SELECT urls.id id, urls.url url, urls.title title, urls.visit_count visit_count, urls.typed_count typed_count, urls.last_visit_time last_visit_time, urls.hidden hidden, visits.visit_time visit_time, visits.from_visit from_visit, visits.visit_duration visit_duration, visits.transition transition, visit_source.source source FROM urls JOIN visits ON urls.id = visits.url LEFT JOIN visit_source ON visits.id = visit_source.id",
		OsPaths: struct {
			Darwin  []uptycs.PathStruct `json:"darwin,omitempty"`
			Debian  []uptycs.PathStruct `json:"debian,omitempty"`
			Windows []uptycs.PathStruct `json:"windows,omitempty"`
		}{
			Debian: []uptycs.PathStruct{{Path: "/home/%/.config/google-chrome/Default/History"}},
		},
		Columns: []struct {
			Name        string `json:"name,omitempty"`
			Description string `json:"description,omitempty"`
		}{
			{Name: "id", Description: "Id"},
			{Name: "path", Description: "Path"},
			{Name: "url", Description: "Url"},
			{Name: "visit_count", Description: "VisitCount"},
			{Name: "typed_count", Description: "TypedCount"},
			{Name: "last_visit_time", Description: "LastVisitTime"},
			{Name: "hidden", Description: "Hidden"},
			{Name: "visit_time", Description: "VisitTime"},
			{Name: "visit_duration", Description: "VisitDuration"},
			{Name: "source", Description: "Source"},
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created atcQuery '%s' with id '%s'", newAtcQuery.Name, newAtcQuery.ID))

	// Update a atcQuery by ID

	log.Println(fmt.Sprintf("Attempting to update atcQuery with id '%s': '%s' to 'marcus test updated'", newAtcQuery.ID, newAtcQuery.Name))
	updatedAtcQuery, err := c.UpdateAtcQuery(uptycs.AtcQuery{
		ID:   newAtcQuery.ID,
		Name: "marcus test updated",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated atcQuery '%s' with id '%s'", updatedAtcQuery.Name, updatedAtcQuery.ID))

	// Delete a atcQuery by ID

	_, err = c.DeleteAtcQuery(uptycs.AtcQuery{
		ID: newAtcQuery.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted atcQuery '%s' with id '%s'", updatedAtcQuery.Name, newAtcQuery.ID))

}
