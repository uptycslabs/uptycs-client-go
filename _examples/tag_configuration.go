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

	tagConfigurationByID, _ := c.GetTagConfiguration(uptycs.TagConfiguration{
		ID: "4fa318d5-e483-4207-b34f-d770339c7bc8",
	})
	log.Printf("Got TagConfiguration with ID '%s'\n", tagConfigurationByID.ID)

	// Create a tagConfiguration
	newTagConfiguration, err := c.CreateTagConfiguration(uptycs.TagConfiguration{
		Key:           "somekey",
		FlagProfileID: "ed4c1240-ffcb-492a-b95f-82976c8bbab5",
		Value:         "aws",
		FilePathGroups: []uptycs.TagConfigurationObject{{
			ID: "a7335d0e-bf70-4ec1-b422-feefedb6dcd9",
		}},
		RegistryPaths:        []uptycs.TagConfigurationObject{},
		Querypacks:           []uptycs.TagConfigurationObject{},
		YaraGroupRules:       []uptycs.TagConfigurationObject{},
		AuditConfigurations:  []uptycs.TagConfigurationObject{},
		EventExcludeProfiles: []uptycs.TagConfigurationObject{},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created tagConfiguration with id '%s' with %d fileGroupPaths and %d registryPaths", newTagConfiguration.ID, len(newTagConfiguration.FilePathGroups), len(newTagConfiguration.RegistryPaths)))

	// Update a tagConfiguration by by ID
	log.Println(fmt.Sprintf("Attempting to update tagConfiguration with id '%s", newTagConfiguration.ID))
	updatedTagConfiguration, err := c.UpdateTagConfiguration(uptycs.TagConfiguration{
		ID:             newTagConfiguration.ID,
		Key:            "somekey",
		FlagProfileID:  "ed4c1240-ffcb-492a-b95f-82976c8bbab5",
		Value:          "notaws",
		FilePathGroups: []uptycs.TagConfigurationObject{},
		RegistryPaths: []uptycs.TagConfigurationObject{
			{
				ID: "ce064913-0c00-4b14-8df3-b1dd90372f04",
			},
		},
		Querypacks:           []uptycs.TagConfigurationObject{},
		YaraGroupRules:       []uptycs.TagConfigurationObject{},
		AuditConfigurations:  []uptycs.TagConfigurationObject{},
		EventExcludeProfiles: []uptycs.TagConfigurationObject{},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated tagConfiguration with id '%s' with %d fileGroupPaths and %d registryPaths", updatedTagConfiguration.ID, len(updatedTagConfiguration.FilePathGroups), len(updatedTagConfiguration.RegistryPaths)))

	// Delete a tagConfiguration by ID
	_, err = c.DeleteTagConfiguration(uptycs.TagConfiguration{
		ID: newTagConfiguration.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted tagConfiguration with id '%s'", newTagConfiguration.ID))

}
