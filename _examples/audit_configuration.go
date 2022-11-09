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

	auditConfigurationByID, _ := c.GetAuditConfiguration(uptycs.AuditConfiguration{
		ID: "7d51a844-f28e-4dbf-8831-e4a063e16156",
	})
	log.Println("Got AuditConfiguration with Name ", auditConfigurationByID.Name)

	auditConfigurationByName, _ := c.GetAuditConfiguration(uptycs.AuditConfiguration{
		Name: "CIS_Distribution_Independent_Linux_Benchmark_v200-Test",
	})
	log.Println("Got AuditConfiguration with ID ", auditConfigurationByName.ID)
}
