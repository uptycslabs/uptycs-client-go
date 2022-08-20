package main

import (
	"github.com/uptycslabs/uptycs-client-go/uptycs"
	"log"
	"os"
)

func main() {
	c, _ := uptycs.NewClient(uptycs.UptycsConfig{
		Host:       os.Getenv("UPTYCS_HOST"),
		ApiKey:     os.Getenv("UPTYCS_API_KEY"),
		ApiSecret:  os.Getenv("UPTYCS_API_SECRET"),
		CustomerID: os.Getenv("UPTYCS_CUSTOMER_ID"),
	})

	// Get User by ID
	userByID, err := c.GetUser(uptycs.User{
		ID: "f48f4c40-9c4a-47bb-9e3f-797d4deca92a",
	})
	if err != nil {
		panic(err)
	}
	log.Println("Got User by ID with Name %s", userByID.Name)

	// Get User by Name
	userByName, err := c.GetUser(uptycs.User{
		Name: "Marcus Young",
	})
	if err != nil {
		panic(err)
	}
	log.Println("Got User by name with id %s", userByName.ID)

	// Create a user
	newUser, err := c.CreateUser(uptycs.User{
		Name:       "Marcus Young",
		Phone:      "866-867-5309",
		Email:      "noone@important.com",
		Active:     true,
		SuperAdmin: true,
		Bot:        true,
		Support:    false,
	})
	if err != nil {
		panic(err)
	}
	log.Println("created new user with id %s", newUser.ID)

	// Update a user
	updatedUser, err := c.UpdateUser(uptycs.User{
		ID:   newUser.ID,
		Name: "Not Marcus Young",
	})
	if err != nil {
		panic(err)
	}
	log.Println("Updated user with id %s. Name is now '%s'", updatedUser.ID)

	// Delete a user
	_, err = c.DeleteUser(uptycs.User{
		ID: newUser.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println("Deleted user with id %s", updatedUser.ID)
}
