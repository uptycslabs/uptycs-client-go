package main

import (
	"github.com/myoung34/uptycs-client-go/uptycs"
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

	destinationByName, _ := c.GetDestination(uptycs.Destination{
		Name: "#reddiconnect-alerts",
	})
	log.Println("Got Destination with ID %s", destinationByName.ID)

	destinationByID, _ := c.GetDestination(uptycs.Destination{
		ID: "b7c9c973-e2a3-4913-a755-919026267679",
	})
	log.Println("Got Destination with Name %s", destinationByID.Name)
}
