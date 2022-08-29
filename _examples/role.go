package main

import (
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

	// Get Role by ID
	roleByID, err := c.GetRole(uptycs.Role{
		ID: "7747f7ab-c859-4d04-ab16-442be2d445cb",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Got Role by ID with Name '%s'\n", roleByID.Name)

	// Get Role by Name
	roleByName, err := c.GetRole(uptycs.Role{
		Name: "user",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Got Role by name with id '%s'\n", roleByName.ID)

	// Create a role
	newRole, err := c.CreateRole(uptycs.Role{
		Name:             "marcus test",
		Permissions:      []string{},
		RoleObjectGroups: []uptycs.ObjectGroup{},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Created new role with id '%s'\n", newRole.ID)

	// Update a role
	updatedRole, err := c.UpdateRole(uptycs.Role{
		ID:                   newRole.ID,
		Name:                 "not marcus test",
		Description:          "test description",
		Permissions:          []string{},
		Custom:               false,
		Hidden:               false,
		NoMinimalPermissions: false,
		RoleObjectGroups:     []uptycs.ObjectGroup{},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Updated role with id '%s'. Name is now '%s'\n", updatedRole.ID, updatedRole.Name)

	// Delete a role
	_, err = c.DeleteRole(uptycs.Role{
		ID: newRole.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Deleted role with id '%s'\n", newRole.ID)
}
