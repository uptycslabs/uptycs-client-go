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

	_, err := c.GetAssetTags()
	if err != nil {
		// error is expected
		log.Println(err)
	}

	_, err = c.GetAssetTag(uptycs.AssetTag{
		Name: "foo",
	})
	if err != nil {
		// error is expected
		log.Println(err)
	}

	// Get an asset
	asset, err := c.GetAsset(uptycs.Asset{
		ID: "ed5750fb-4588-59b6-b15f-35cc7dc2fbb9",
	})
	log.Printf("Got asset by id with hostname %s\n  Has tags: %+v", asset.HostName, asset.Tags)

	// Attempt to  an asset
	_, err = c.UpdateAsset(uptycs.Asset{})
	if err != nil {
		// error is expected
		log.Println(err)
	}
}
