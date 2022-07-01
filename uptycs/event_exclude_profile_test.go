package uptycs

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestGetEventExcludeProfile(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		id      string
		out     interface{}
	}

	theTests := []convTest{
		{
			name:    "TestEventExcludeProfile",
			fixture: "fixtures/eventExcludeProfile.json",
			id:      "13da8bc4-3c70-4bb9-a4d4-7ca320860926",
			out: EventExcludeProfile{
				ID:          "13da8bc4-3c70-4bb9-a4d4-7ca320860926",
				CustomerID:  "11111111-1111-1111-1111-111111111111",
				Name:        "Uptycs default event exclude profile",
				Description: "Filters known paths, IPs and domains",
				Priority:    20200227,
				Metadata: EventExcludeProfileMetadata{
          HttpEvents: HttpEvents{
            Host: []string{
              "www.google.com",
            },
          },
					SocketEvents: SocketEvents{
						RemoteAddress: []string{
							"^S",
							"^::S",
						},
					},
					ProcessEvents: ProcessEvents{
						Path: []string{
							"^.*ntp\\.orgS",
							"^/bin/bashS",
						},
					},
					ProcessFileEvents: ProcessFileEvents{
						Path: []string{
							"\\.uptycs.ioS",
							".*(cache|notification|localstate|resource|safety|automaticdestination|packages|tempstate|((edb|cf|gthr|dir|crwl|exd|db-journal|aodl|evtx|json|dat|log|tmp|etl|db|ini|xml|chk|jfm|pf|temp))S)",
						},
						Operation: []string{
							"^open\\+readS",
							"attributes_modified",
							"unlink",
						},
						Executable: []string{
							"^.*osqueryd\\.exeS|^.*collectguestlogs\\.exeS|^.*MsMpEng\\.exeS",
						},
					},
					UserEvents: UserEvents{
						Message: []string{
							"^cwd=.*",
							"^op=PAM:accountingS",
							"^op=PAM:session_closeS",
							"^op=PAM:session_openS",
							"^op=PAM:setcredS",
						},
					},
					RegistryEvents: RegistryEvents{
						Action: []string{
							"SET_INFORMATION",
							"CREATED",
						},
					},
					EbpfDnsLookupEvents: EbpfDnsLookupEvents{
						Answer: []string{
							"foo",
						},
						Question: []string{
							"wut",
						},
					},
					DnsLookupEvents: DnsLookupEvents{
						Answer: []string{
							"^(?![\\s\\S])",
						},
						Question: []string{
							"^0\\.pool\\.ntp\\.orgS",
							"^1\\.pool\\.ntp\\.orgS",
						},
					},
				},
				CreatedAt:    "2022-06-27T04:09:05.342Z",
				CreatedBy:    "00000000-0000-0000-0000-000000000000",
				UpdatedAt:    "2022-06-27T04:09:05.342Z",
				UpdatedBy:    "00000000-0000-0000-0000-000000000000",
				ResourceType: "asset",
				Platform:     "all",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Event exlude profile information", Href: "/api/customers/11111111-1111-1111-1111-111111111111/eventExcludeProfiles/13da8bc4-3c70-4bb9-a4d4-7ca320860926"},
					LinkItem{Rel: "parent", Title: "Event exlude profiles information", Href: "/api/customers/11111111-1111-1111-1111-111111111111/eventExcludeProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/eventExcludeProfiles/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			eventExcludeProfileResp, err := c.GetEventExcludeProfile(EventExcludeProfile{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(eventExcludeProfileResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", eventExcludeProfileResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteEventExcludeProfile(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   EventExcludeProfile
	}

	theTests := []convTest{
		{
			name: "TestEventExcludeProfile",
			in: EventExcludeProfile{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/eventExcludeProfiles/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteEventExcludeProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/eventExcludeProfiles/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutEventExcludeProfile(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   EventExcludeProfile
	}

	theTests := []convTest{
		{
			name: "TestEventExcludeProfile",
			in: EventExcludeProfile{
				ID:          "13da8bc4-3c70-4bb9-a4d4-7ca320860926",
				CustomerID:  "11111111-1111-1111-1111-111111111111",
				Name:        "Uptycs default event exclude profile",
				Description: "Filters known paths, IPs and domains",
				Priority:    20200227,
				Metadata:    EventExcludeProfileMetadata{},
				Platform:    "all",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/eventExcludeProfiles/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.UpdateEventExcludeProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/eventExcludeProfiles/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateEventExcludeProfile(t *testing.T) {

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      EventExcludeProfile
	}

	theTests := []convTest{
		{
			name:    "TestEventExcludeProfile",
			fixture: "fixtures/eventExcludeProfileCreate.json",
			in: EventExcludeProfile{
				ID:          "13da8bc4-3c70-4bb9-a4d4-7ca320860926",
				CustomerID:  "11111111-1111-1111-1111-111111111111",
				Name:        "Uptycs default event exclude profile",
				Description: "Filters known paths, IPs and domains",
				Priority:    20200227,
				Metadata:    EventExcludeProfileMetadata{},
				Platform:    "all",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/eventExcludeProfiles",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateEventExcludeProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/eventExcludeProfiles"], 1)
		})
	}
}
