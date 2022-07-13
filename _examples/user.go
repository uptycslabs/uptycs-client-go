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

	// Get User by ID
	userByID, err := c.GetUser(uptycs.User{
		ID: "f48f4c40-9c4a-47bb-9e3f-797d4deca92a",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Got User by ID with Name %s\n", userByID.Name)

	// Get User by Name
	userByName, err := c.GetUser(uptycs.User{
		Name: "Marcus Young",
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Got User by name with id '%s'\n", userByName.ID)

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

	// get the first objectGroup
	objectGroups, _ := c.GetObjectGroups()

	// Create a role
	newRole, err := c.CreateRole(uptycs.Role{
		Name: "marcus test",
	})

	log.Printf("Created new role with id '%s'\n", newRole.ID)

	// Create a user
	newUser, err := c.CreateUser(uptycs.User{
		Name:               "Marcus Young",
		Phone:              "866-867-5309",
		Email:              "noone@important.com",
		Active:             true,
		AlertHiddenColumns: []string{},
		ImageURL:           "asdf",
		MaxIdleTimeMins:    10,
		SuperAdmin:         true,
		Bot:                true,
		Roles:              []uptycs.Role{newRole},
		Support:            false,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Created new user with id '%s'\n", newUser.ID)

	// Update a role
	newRole.Name = "not marcus test"
	newRole.RoleObjectGroups = []uptycs.ObjectGroup{
		uptycs.ObjectGroup{
			ObjectGroupID: objectGroups.Items[0].ID,
		},
	}
	updatedRole, err := c.UpdateRole(newRole)
	if err != nil {
		panic(err)
	}
	log.Printf("Updated role with id '%s'. Name is now '%s'\n", updatedRole.ID, updatedRole.Name)

	// Update a user
	updatedUser, err := c.UpdateUser(uptycs.User{
		ID:                 newUser.ID,
		Name:               "Not Marcus Young",
		Phone:              "866-867-5309",
		Email:              "noone@important.com",
		Active:             true,
		AlertHiddenColumns: []string{},
		ImageURL:           "asdf",
		MaxIdleTimeMins:    10,
		SuperAdmin:         true,
		Bot:                true,
		Roles:              []uptycs.Role{updatedRole},
		Support:            false,
		UserObjectGroups: []uptycs.ObjectGroup{
			uptycs.ObjectGroup{
				ObjectGroupID: objectGroups.Items[1].ID,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Updated user with id '%s'. Name is now '%s'\n", updatedUser.ID, updatedUser.Name)

	// Delete a user
	_, err = c.DeleteUser(uptycs.User{
		ID: newUser.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Deleted user with id '%s'\n", updatedUser.ID)

	// Delete a role
	_, err = c.DeleteRole(uptycs.Role{
		ID: updatedRole.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Printf("Deleted role with id '%s'\n", updatedRole.ID)
}
