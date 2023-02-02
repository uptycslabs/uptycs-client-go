package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAsset(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
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
			name:    "TestAsset",
			fixture: "fixtures/asset.json",
			id:      "93e08dcb-4454-48dd-bfb6-95b004246450",
			out: Asset{
				Gateway:               "50.35.120.181",
				CityID:                "93e08dcb-4454-48dd-bfb6-95b004246450",
				CreatedAt:             "2021-08-03T16:57:26.027Z",
				Disabled:              true,
				HostName:              "C02F50H1MD6V",
				ID:                    "ed5750fb-4588-59b6-b15f-35cc7dc2fbb9",
				LastEnrolledAt:        "2022-09-15T06:47:16.479Z",
				Os:                    "macOS",
				OsFlavor:              "darwin",
				OsqueryVersion:        "5.0.1.16-Uptycs",
				OsVersion:             "12.6",
				Status:                "active",
				UpgradeState:          false,
				ObjectGroupID:         "6f49e4cb-ae9e-4a8a-a80b-4f747731b802",
				Live:                  false,
				Location:              "Snohomish, Washington, United States",
				ManualSlackAssignment: false,
				HardwareVendor:        "Apple Inc.",
				AssetObjectGroupID:    "6f49e4cb-ae9e-4a8a-a80b-4f747731b802",
				PackageObjectGroupID:  "6f49e4cb-ae9e-4a8a-a80b-4f747731b802",
				ObjectGroup: ObjectGroup{
					ID:            "6f49e4cb-ae9e-4a8a-a80b-4f747731b802",
					Name:          "assets",
					Description:   "Default asset group",
					Secret:        "67066459-0d4f-4ed8-a7cf-2ca0549ccdee",
					ObjectType:    "ASSET",
					Custom:        false,
					RetentionDays: 0,
					RangerID:      296,
					CreatedAt:     "2021-06-01T17:39:15.969Z",
					UpdatedAt:     "2021-10-18T15:57:04.023Z",
				},
				City: City{
					ID:                 "93e08dcb-4454-48dd-bfb6-95b004246450",
					Name:               "Snohomish",
					SubdivisionIsoCode: "WA",
					SubdivisionName:    "Washington",
					CountryIsoCode:     "US",
					CountryName:        "United States",
				},
				LastActivityAt: "2022-07-01T17:21:44.056Z",
				Tags: []string{
					"darwin",
					"asset-group=assets",
					"all",
				},
				OsDisplay:                 "macOS 12.6",
				Latitude:                  47.8581,
				Longitude:                 -122.0872,
				CPUBrand:                  "Intel(R) Core(TM) i9-9980HK CPU @ 2.40GHz",
				HardwareModel:             "MacBookPro16,1",
				HardwareSerial:            "C02F50H1MD6V",
				Cores:                     8,
				LogicalCores:              16,
				MemoryMb:                  32768,
				OsKey:                     "darwin_12.6",
				OsVersionSortable:         "0000000012.0000000006",
				ProfileID:                 "darwin_seed",
				OsqueryVersionSortable:    "0000000005.0000000000.0000000001.016-Uptycs",
				NewEpoch:                  false,
				LastEpoch:                 "2022-06-30T16:38:08.739Z",
				Flags:                     "{\"disable_audit\":false,\"disable_carver\":false,\"enable_wmi\":true,\"enable_yara_process_events\":true,\"generate_process_hash_in_process_event\":true,\"watchdog_utilization_limit\":\"0\",\"win_allow_sockets\":true}",
				UpdatedBy:                 "3d56cffc-1c2d-42ff-aedd-06d6ca18b463",
				UpdatedAt:                 "2022-09-15T06:47:16.484Z",
				FlagsFile:                 "--add_container_image_to_events=true",
				FlagsFileChecksum:         "f9c64898ea0c7699430d5f23f612794e7083513cb0f548cab9aad03254cd161a",
				ActiveFlagProfileID:       "9c78cc0a-7cac-42ca-92ae-ceac261b838d",
				ActiveFlagProfileSource:   "tag",
				ActiveCustomProfileSource: "default",
				Protect:                   false,
				AgentVersion:              "5.0.1.16-Uptycs",
				AgentVersionSortable:      "0000000005.0000000000.0000000001.016-Uptycs",
				AgentID:                   "ed5750fb-4588-59b6-b15f-35cc7dc2fbb9",
				AgentType:                 "asset",
				ResourceType:              "asset",
				Arch:                      "x86_64h",
				Platform:                  "darwin",
				PlatformLike:              "darwin",
				EventExcludeProfiles:      []EventExcludeProfiles{},
				AssetCapabilities: []AssetCapabilities{
					{
						ID:      "334e7585-e09d-400f-80ba-198bf83b4387",
						AssetID: "ed5750fb-4588-59b6-b15f-35cc7dc2fbb9",
						Name:    "Yara",
						Status:  "PARTIAL",
						IndividualStatus: struct {
							ConfigurationStatus  string `json:"configurationStatus,omitempty"`
							FlagStatus           string `json:"flagStatus,omitempty"`
							ScheduledQueryStatus string `json:"scheduledQueryStatus,omitempty"`
						}{
							FlagStatus:           "CONFIGURED",
							ConfigurationStatus:  "PARTIAL",
							ScheduledQueryStatus: "CONFIGURED",
						},
						CreatedAt: "2022-04-11T16:29:00.362Z",
						UpdatedAt: "2022-04-30T07:47:42.063Z",
					},
					{
						ID:        "4830e4a7-a434-4385-b5bc-7659bcde2b8f",
						AssetID:   "ed5750fb-4588-59b6-b15f-35cc7dc2fbb9",
						Name:      "Blocking",
						Status:    "NOT-CONFIGURED",
						CreatedAt: "2022-04-11T16:29:00.398Z",
						UpdatedAt: "2022-04-30T07:47:42.080Z",
					},
				},
				Interfaces: []AssetInterface{
					{
						Name:      "lo0",
						Mac:       "00:00:00:00:00:00",
						IP:        "127.0.0.1",
						Mask:      "255.0.0.0",
						IsPrimary: false,
					},
					{
						Name:      "lo0",
						Mac:       "00:00:00:00:00:00",
						IP:        "::1",
						Mask:      "ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",
						IsPrimary: false,
					},
				},
				Links: []LinkItem{
					{
						Rel:   "self",
						Title: "Asset information",
						Href:  "/api/customers/11111111-1111-1111-1111-111111111111/assets/ed5750fb-4588-59b6-b15f-35cc7dc2fbb9",
					},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/assets/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			assetResp, err := c.GetAsset(Asset{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(assetResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("  Actual: %v", assetResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteAsset(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   Asset
	}

	theTests := []convTest{
		{
			name: "TestAsset",
			in: Asset{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/assets/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteAsset(theT.in)
			assert.Equal(t, err.Error(), "DELETE is not supported for assets")
			countInfo := httpmock.GetCallCountInfo()
			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/assets/%v", theT.in.ID)], 0)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutAsset(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Asset
	}

	theTests := []convTest{
		{
			name:    "TestAsset",
			fixture: "fixtures/assetCreate.json",
			in: Asset{
				ID: "b7c9c973-e2a3-4913-a755-919026267679",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/assets/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateAsset(theT.in)
			assert.Equal(t, err.Error(), "PUT is not supported for assets")
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/assets/%v", theT.in.ID)], 0)
		})
	}
}

func TestCreateAsset(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Asset
	}

	theTests := []convTest{
		{
			name:    "TestAsset",
			fixture: "fixtures/assetCreate.json",
			in: Asset{
				ID: "b7c9c973-e2a3-4913-a755-919026267679",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/assets",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateAsset(theT.in)
			assert.Equal(t, err.Error(), "POST is not supported for assets")
			countInfo := httpmock.GetCallCountInfo()
			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/assets"], 0)
		})
	}
}
