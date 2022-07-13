package main

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
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

	// Get all event exclude profiles
	excludeProfiles, _ := c.GetEventExcludeProfiles()
	for _, excludeProfile := range excludeProfiles.Items {
		if excludeProfile.Name == "marcus test" {
			log.Println(fmt.Sprintf("%s has id %s", excludeProfile.Name, excludeProfile.ID))
		}
	}

	// Get an exclude profile by ID
	rule, _ := c.GetEventExcludeProfile(uptycs.EventExcludeProfile{
		ID: "2a86d4ad-3aa3-42f1-8430-6da238c82b11",
	})
	log.Println(fmt.Sprintf("Found profile by ID with name %s as %s", rule.Name, rule.ID))

	eventExcludeProfile, err := c.CreateEventExcludeProfile(uptycs.EventExcludeProfile{
		Name:        "marc test",
		Description: "marcs test",
		Priority:    999999999,
		Metadata: uptycs.EventExcludeProfileMetadata{
			SocketEvents: uptycs.SocketEvents{
				RemoteAddress: []string{
					"^foo",
				},
			},
			ProcessEvents: uptycs.ProcessEvents{
				Path: []string{
					"^.*test\\.orgS",
				},
			},
			ProcessFileEvents: uptycs.ProcessFileEvents{
				Path: []string{
					"\\.test.ioS",
				},
				Operation: []string{
					"attributes_modified",
				},
				Executable: []string{
					"^.*foo\\.exeS",
				},
			},
			UserEvents: uptycs.UserEvents{
				Message: []string{
					"^op=PAM:fooS",
				},
			},
			RegistryEvents: uptycs.RegistryEvents{
				Action: []string{
					"SET_INFORMATION",
					"CREATED",
				},
			},
			DnsLookupEvents: uptycs.DnsLookupEvents{
				Question: []string{
					"^0\\.foo\\.ntp\\.orgS",
					"^1\\.pool\\.ntp\\.orgS",
				},
			},
		},
		ResourceType: "asset",
		Platform:     "all",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("Created Event Exclude Profile '%s' with id %s", eventExcludeProfile.Name, eventExcludeProfile.ID))

	// Update but use JSON for the metadata
	_, err = c.UpdateEventExcludeProfile(uptycs.EventExcludeProfile{
		ID: eventExcludeProfile.ID,
		MetadataJson: heredoc.Doc(`
  {
    "process_events": {
      "path": [
        "^/Library/Developer/Xcode$"
      ]
    },
    "process_file_events": {
      "path": [
        "^/Library/Developer/Xcode$",
        "^/Library/Application Support/JAMF$"
      ],
      "executable": [
        "^.*osqueryd\\.exe$|^.*collectguestlogs\\.exe$|^.*MsMpEng\\.exe$"
      ]
    }
  }
	    `),
		Name:         "marc test",
		Description:  "marcs test",
		Priority:     76,
		ResourceType: "asset",
		Platform:     "all",
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fmt.Sprintf("Updated Event Exclude Profile '%s' with id %s", eventExcludeProfile.Name, eventExcludeProfile.ID))

	// Delete an event exclude profile by ID
	_, err = c.DeleteEventExcludeProfile(uptycs.EventExcludeProfile{
		ID: eventExcludeProfile.ID,
	})
	log.Println(fmt.Sprintf("Deleted Event Exclude Profile '%s' with id %s", eventExcludeProfile.Name, eventExcludeProfile.ID))

}
