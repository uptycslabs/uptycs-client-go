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

	lookupTableByID, _ := c.GetLookupTable(uptycs.LookupTable{
		ID: "7f042e5a-e4e8-47e7-8077-a0581411d4fd",
	})
	log.Println("Got LookupTable with ID ", lookupTableByID.ID)
	_dataRow, _ := c.GetLookupTableDataRow(lookupTableByID, lookupTableByID.DataRows[0])
	log.Println(fmt.Sprintf("%+v", _dataRow.Data))

	// Create a lookupTable
	newLookupTable, err := c.CreateLookupTable(uptycs.LookupTable{
		Name:        "test",
		Description: "a test",
		IDField:     "id",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created lookupTable '%s' with id '%s'", newLookupTable.Name, newLookupTable.ID))

	// Add some data rows
	_, err = c.CreateLookupTableDataRow(
		newLookupTable,
		uptycs.LookupTableDataRow{
			Data: "[{\"id\": \"foo\", \"exception\":\"bar\"}]",
		},
	)
	_, err = c.CreateLookupTableDataRow(
		newLookupTable,
		uptycs.LookupTableDataRow{
			Data: "[{\"id\": \"asdf\", \"exception\":\"wut\"}]",
		},
	)

	//Update Some data
	_, err = c.UpdateLookupTableDataRow(
		newLookupTable,
		uptycs.LookupTableDataRow{
			IDFieldValue: "foo",
			Data:         "[{\"id\": \"foo\", \"exception\":\"new bar\"}]",
		},
	)

	_updatedLookupTable, err := c.UpdateLookupTable(uptycs.LookupTable{
		ID:   newLookupTable.ID,
		Name: "test updated",
	})
	if err != nil {
		panic(err)
		log.Println(fmt.Sprintf("Error updating lookupTable '%s' with id '%s'", "", "54c7e31e-02a9-4b58-aaaf-e4a3b09e1980"))
	}
	log.Println(fmt.Sprintf("Updated lookupTable '%s' with id '%s'", _updatedLookupTable.Name, newLookupTable.ID))

	//Delete Some data
	_, err = c.DeleteLookupTableDataRow(
		newLookupTable,
		uptycs.LookupTableDataRow{
			IDFieldValue: "foo",
		},
	)
	_test, _ := c.GetLookupTable(uptycs.LookupTable{
		ID: newLookupTable.ID,
	})
	log.Println(fmt.Sprintf("%+v", _test.DataRows[0].Data))

	// Delete a lookupTable by ID
	_, err = c.DeleteLookupTable(uptycs.LookupTable{
		ID: newLookupTable.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted lookupTable '%s' with id '%s'", newLookupTable.Name, newLookupTable.ID))

}
