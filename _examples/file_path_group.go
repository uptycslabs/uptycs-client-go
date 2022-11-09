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

	filePathGroupByName, _ := c.GetFilePathGroup(uptycs.FilePathGroup{
		Name: "fim_testing",
	})
	log.Println("Got FilePathGroup with ID ", filePathGroupByName.ID)

	filePathGroupByID, _ := c.GetFilePathGroup(uptycs.FilePathGroup{
		ID: "2c4e7ba8-6456-4f67-8a3f-332043642be8",
	})
	log.Println("Got FilePathGroup with Name ", filePathGroupByID.Name)

	// Create a filePathGroup

	newFilePathGroup, err := c.CreateFilePathGroup(uptycs.FilePathGroup{
		Name:         "marcus test",
		Description:  "test",
		IncludePaths: []string{"/poop"},
		Custom:       true,
		Signatures:   []uptycs.FilePathGroupSignature{},
		YaraGroupRules: []uptycs.YaraGroupRule{
			{ID: "9a5a3262-ee74-417c-ade0-c1948ec8bc27"},
			{Name: "Uptycs Demo Detection"},
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created filePathGroup '%s' with id '%s'", newFilePathGroup.Name, newFilePathGroup.ID))

	// Update a filePathGroup by ID

	log.Println(fmt.Sprintf("Attempting to update filePathGroup with id '%s': '%s' to 'marcus test updated'", newFilePathGroup.ID, newFilePathGroup.Name))
	updatedFilePathGroup, err := c.UpdateFilePathGroup(uptycs.FilePathGroup{
		ID:           newFilePathGroup.ID,
		Name:         "marcus test updated",
		IncludePaths: []string{"/poop"},
		Custom:       true,
		Signatures:   []uptycs.FilePathGroupSignature{},
		YaraGroupRules: []uptycs.YaraGroupRule{
			{ID: "9a5a3262-ee74-417c-ade0-c1948ec8bc27"},
			{Name: "Uptycs Demo Detection"},
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated filePathGroup '%s' with id '%s'", updatedFilePathGroup.Name, updatedFilePathGroup.ID))

	// Delete a filePathGroup by ID

	_, err = c.DeleteFilePathGroup(uptycs.FilePathGroup{
		ID: newFilePathGroup.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted filePathGroup '%s' with id '%s'", updatedFilePathGroup.Name, newFilePathGroup.ID))

}
