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

	registryPathByName, _ := c.GetRegistryPath(uptycs.RegistryPath{
		Name: "upt_lsa_provider",
	})
	log.Println("Got RegistryPath with ID ", registryPathByName.ID)

	registryPathByID, _ := c.GetRegistryPath(uptycs.RegistryPath{
		ID: "0d52d170-f43a-4705-9c44-16133f31ecf4",
	})
	log.Println("Got RegistryPath with Name ", registryPathByID.Name)

	// Create a registryPath

	newRegistryPath, err := c.CreateRegistryPath(uptycs.RegistryPath{
		Name: "marcus test",
		IncludeRegistryPaths: []string{
			"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Lsa\\%",
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created registryPath '%s' with id '%s'", newRegistryPath.Name, newRegistryPath.ID))

	// Update a registryPath by ID

	log.Println(fmt.Sprintf("Attempting to update registryPath with id '%s': '%s' to 'marcus test updated'", newRegistryPath.ID, newRegistryPath.Name))
	updatedRegistryPath, err := c.UpdateRegistryPath(uptycs.RegistryPath{
		ID:          newRegistryPath.ID,
		Name:        "marcus test updated",
		Description: "my test",
		IncludeRegistryPaths: []string{
			"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Lsa\\%",
		},
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated registryPath '%s' with id '%s'", updatedRegistryPath.Name, updatedRegistryPath.ID))

	// Delete a registryPath by ID

	_, err = c.DeleteRegistryPath(uptycs.RegistryPath{
		ID: newRegistryPath.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted registryPath '%s' with id '%s'", updatedRegistryPath.Name, newRegistryPath.ID))

}
