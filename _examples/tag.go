package main

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
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

	tagByID, _ := c.GetTag(uptycs.Tag{
		ID: "bfc1f017-7bd1-4af8-a512-f90e9b2a41ce",
	})
	log.Printf("Got Tag with ID '%s'\n", tagByID.ID)

	tagByKeyValue, _ := c.GetTag(uptycs.Tag{
		Key:   "asset-group",
		Value: "enrolling",
	})
	log.Printf("Got Tag with ID '%s' by Key '%s' and Value '%s'\n", tagByKeyValue.ID, tagByKeyValue.Key, tagByKeyValue.Value)

	// Create a tag
	newTag, err := c.CreateTag(uptycs.Tag{
		Key:           "somekey",
		FlagProfileID: "ed4c1240-ffcb-492a-b95f-82976c8bbab5",
		Value:         "aws",
		FilePathGroups: []uptycs.TagConfigurationObject{
			{
				ID: "a7335d0e-bf70-4ec1-b422-feefedb6dcd9",
			},
			{
				Name: "FIM - USB Storage File Events",
			},
		},
		RegistryPaths:        []uptycs.TagConfigurationObject{},
		Querypacks:           []uptycs.TagConfigurationObject{},
		YaraGroupRules:       []uptycs.TagConfigurationObject{},
		AuditConfigurations:  []uptycs.TagConfigurationObject{},
		EventExcludeProfiles: []uptycs.TagConfigurationObject{},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created tag with id '%s' with %d fileGroupPaths and %d registryPaths", newTag.ID, len(newTag.FilePathGroups), len(newTag.RegistryPaths)))

	// Create a customProfile
	newCustomProfile, err := c.CreateCustomProfile(uptycs.CustomProfile{
		Name:        "marcus test",
		Description: "Test",
		QuerySchedules: uptycs.CustomJSONString(heredoc.Doc(`{
		  "processes": 100
		}`)),
		Priority:     2,
		ResourceType: "asset",
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created customProfile '%s' with id '%s'", newCustomProfile.Name, newCustomProfile.ID))

	// Update a tag by ID
	log.Println(fmt.Sprintf("Attempting to update tag with id '%s", newTag.ID))
	updatedTag, err := c.UpdateTag(uptycs.Tag{
		ID:             newTag.ID,
		Key:            "somekey",
		FlagProfileID:  "ed4c1240-ffcb-492a-b95f-82976c8bbab5",
		Value:          "notaws",
		FilePathGroups: []uptycs.TagConfigurationObject{},
		RegistryPaths: []uptycs.TagConfigurationObject{
			{
				ID: "ce064913-0c00-4b14-8df3-b1dd90372f04",
			},
		},
		CustomProfile:        newCustomProfile.Name,
		Querypacks:           []uptycs.TagConfigurationObject{},
		YaraGroupRules:       []uptycs.TagConfigurationObject{},
		AuditConfigurations:  []uptycs.TagConfigurationObject{},
		EventExcludeProfiles: []uptycs.TagConfigurationObject{},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated tag with id '%s' with %d fileGroupPaths and %d registryPaths", updatedTag.ID, len(updatedTag.FilePathGroups), len(updatedTag.RegistryPaths)))

	// Delete a tag by ID
	_, err = c.DeleteTag(uptycs.Tag{
		ID: newTag.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted tag with id '%s'", newTag.ID))

	// Delete a customProfile by ID
	_, err = c.DeleteCustomProfile(uptycs.CustomProfile{
		ID: newCustomProfile.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted customProfile '%s' with id '%s'", newCustomProfile.Name, newCustomProfile.ID))

}
