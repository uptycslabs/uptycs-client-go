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

	querypackByID, _ := c.GetQuerypack(uptycs.Querypack{
		ID: "65fdf96b-3c4a-44cb-ad9b-06424f81a90a",
	})
	log.Println("Got Querypack with Name ", querypackByID.Name)

	querypackByName, _ := c.GetQuerypack(uptycs.Querypack{
		Name: "spring4shell",
	})
	log.Println("Got Querypack with ID ", querypackByName.ID)

	// Create a querypack
	newQuerypack, err := c.CreateQuerypack(uptycs.Querypack{
		Name:        "marcus_test",
		Description: "marcus test",
		Type:        "vulnerability",
		Conf: uptycs.CustomJSONString(heredoc.Doc(`{
			"queries": {
				"linux_baseline": {
					"description": "",
					"query": "SELECT\n    path,\n    directory,\n    filename,\n    symlink\nFROM\n    file\nWHERE\n    (\n        path like '/usr/lib/%%'\n        OR path like '/lib64/%%'\n        OR path like '/bin/%%'\n        OR path like '/sbin/%%'\n        OR path like '/usr/bin/%%'\n        OR path like '/usr/sbin/%%'\n        OR path like '/usr/local/bin/%%'\n        OR path like '/usr/local/sbin/%%'\n    )\n    and filename != '.'",
					"removed": true,
					"version": null,
					"interval": 86400,
					"platform": "linux",
					"snapshot": true,
					"runNow": false,
					"value": ""
				},
				"linux_baseline_lib_directory": {
					"description": "",
					"query": "SELECT\n    path,\n    directory,\n    filename,\n    symlink\nFROM\n    file\nWHERE path like '/lib/%%'\n  and filename != '.'",
					"removed": true,
					"version": null,
					"interval": 86400,
					"platform": "linux",
					"snapshot": true,
					"runNow": false,
					"value": ""
				}
			}
		}`)),
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Created querypack '%s' with id '%s'", newQuerypack.Name, newQuerypack.ID))

	// Update a querypack by ID
	log.Println(fmt.Sprintf("Attempting to update querypack with id '%s': '%s' to 'marcus test updated'", newQuerypack.ID, newQuerypack.Name))
	updatedQuerypack, err := c.UpdateQuerypack(uptycs.Querypack{
		ID:          newQuerypack.ID,
		Description: "marcus newest",
		Type:        "vulnerability",
		Conf: uptycs.CustomJSONString(heredoc.Doc(`{
			"queries": {
				"linux_baseline": {
					"description": "",
					"query": "SELECT\n    path,\n    directory,\n    filename,\n    symlink\nFROM\n    file\nWHERE\n    (\n        path like '/usr/lib/%%'\n        OR path like '/lib64/%%'\n        OR path like '/bin/%%'\n        OR path like '/sbin/%%'\n        OR path like '/usr/bin/%%'\n        OR path like '/usr/sbin/%%'\n        OR path like '/usr/local/bin/%%'\n        OR path like '/usr/local/sbin/%%'\n    )\n    and filename != '.'",
					"removed": true,
					"version": null,
					"interval": 86400,
					"platform": "linux",
					"snapshot": true,
					"runNow": false,
					"value": ""
				},
				"linux_baseline_lib_directory": {
					"description": "",
					"query": "SELECT\n    path,\n    directory,\n    filename,\n    symlink\nFROM\n    file\nWHERE path like '/lib/%%'\n  and filename != '.'",
					"removed": true,
					"version": null,
					"interval": 86400,
					"platform": "linux",
					"snapshot": true,
					"runNow": false,
					"value": ""
				}
			}
		}`)),
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Updated querypack '%s' with id '%s'", updatedQuerypack.Name, updatedQuerypack.ID))

	// Delete a querypack by ID

	_, err = c.DeleteQuerypack(uptycs.Querypack{
		ID: newQuerypack.ID,
	})
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Deleted querypack '%s' with id '%s'", updatedQuerypack.Name, newQuerypack.ID))

}
