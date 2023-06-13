package main

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/uptycslabs/uptycs-client-go/uptycs"
	"log"
	"os"
	"time"
)

func main() {
	c, _ := uptycs.NewClient(uptycs.Config{
		Host:       os.Getenv("UPTYCS_HOST"),
		APIKey:     os.Getenv("UPTYCS_API_KEY"),
		APISecret:  os.Getenv("UPTYCS_API_SECRET"),
		CustomerID: os.Getenv("UPTYCS_CUSTOMER_ID"),
	})

	newQueryJob, err := c.CreateQueryJob(uptycs.QueryJob{
		Name:  "test",
		Query: heredoc.Doc(`SELECT instance_id,tags from aws_ec2_instance_current limit 1;`),
		Type:  "global",
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("Created Query Job  with id '%s'", newQueryJob.ID))
	time.Sleep(10 * time.Second)

	_foo, err := c.GetQueryJobResults(uptycs.QueryJobResult{
		ID: newQueryJob.ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(fmt.Sprintf("%+v", _foo.Items[0].RowData))

}
