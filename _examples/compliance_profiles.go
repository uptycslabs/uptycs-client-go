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

	// Get Compliance profiles
	complianceProfiles, err := c.GetComplianceProfiles()
	if err != nil {
		panic(err)
	}
	for _, item := range complianceProfiles.Items{
		log.Printf("Name: %v, ID: %v \n", item.Name, item.ID)
	}
		
	complianceProfilesByID, err2 := c.GetComplianceProfile(uptycs.ComplianceProfile{
			ID: complianceProfiles.Items[0].ID,
		},
	)
	if err2 != nil {
		panic(err2)
	}
	log.Printf("Got ComplianceProfile by ID with Name %s\n", complianceProfilesByID.Name)

	// Below works, but without perms to delete doesn't make sense to run
	// newComplianceProfile, err3 := c.CreateComplianceProfile(uptycs.ComplianceProfile{
	// 	Name: "Test Compliance profile",
	// 	Description: "test compliance profile",
	// 	Priority: 10,
	// })
	// if err3 != nil {
	// 	panic(err3)
	// }
	// log.Printf("Created new ComplianceProfile with ID %s\n", newComplianceProfile.ID)

		// No permissions to delete compliance profile
	// _, err4 := c.DeleteComplianceProfile(uptycs.ComplianceProfile{
	// 		ID: "24829f9b-ce8e-4c6b-89ee-0d3889da3c34",
	// 	},
	// )
	// if err4 != nil {
	// 	panic(err4)
	// }

}
