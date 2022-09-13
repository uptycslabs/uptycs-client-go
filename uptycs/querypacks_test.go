package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetQuerypacks(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/querypacks",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/querypacks.json")
			if err != nil {
				t.Errorf(err.Error())
			}
			return fixture, err
		},
	)

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	querypacksResp, err := c.GetQuerypacks()
	if err != nil {
		t.Errorf(err.Error())
	}

	type convTest struct {
		name string
		in   interface{}
		out  interface{}
	}

	theTests := []convTest{
		{
			name: "thing",
			in:   querypacksResp,
			out: Querypacks{
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Query packs information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/querypacks"},
					LinkItem{Rel: "parent", Title: "Customer information", Href: "/api/customers/111111111111-111111-11111-111111-111111111"},
				},
				Items: []Querypack{
					{
						ID:               "ac4cea47-17d4-4199-8fa8-f658355c36ae",
						Name:             "asset_catalog_os_baseline",
						Description:      "linux os baseline for list binaries",
						Type:             "default",
						AdditionalLogger: false,
						Custom:           false,
						CreatedBy:        "00000000-0000-0000-0000-000000000000",
						UpdatedBy:        "00000000-0000-0000-0000-000000000000",
						CreatedAt:        "2022-05-02T06:29:48.711Z",
						UpdatedAt:        "2022-09-13T04:46:50.098Z",
						IsInternal:       false,
						ResourceType:     "asset",
						Queries: []Query{
							{
								ID:          "a541f9eb-59d9-464d-ab35-854b47dd06f2",
								Name:        "linux_baseline_lib_directory",
								Description: "",
								Query:       "SELECT\n    path,\n    directory,\n    filename,\n    symlink\nFROM\n    file\nWHERE path like '/lib/%%'\n  and filename != '.'",
								Removed:     true,
								//Version: null,
								Interval:    86400,
								Platform:    "linux",
								Snapshot:    true,
								RunNow:      false,
								Value:       "",
								QuerypackID: "ac4cea47-17d4-4199-8fa8-f658355c36ae",
								TableName:   "qp_asset_catalog_os_baseline_q_linux_baseline_lib_directory",
								DataTypes:   "{\"directory\":\"VARCHAR\",\"filename\":\"VARCHAR\",\"path\":\"VARCHAR\",\"symlink\":\"BIGINT\"}",
								Verified:    true,
								CreatedBy:   "00000000-0000-0000-0000-000000000000",
								//UpdatedBy: null,
								CreatedAt: "2022-09-13T04:46:50.152Z",
								UpdatedAt: "2022-09-13T04:46:50.152Z",
							},
						},
						Conf: "{\"queries\":{\"linux_baseline_lib_directory\":{\"description\":\"\"}}}",
					},
				},
				Offset: 0,
				Limit:  1000,
			},
		},
	}

	for _, theT := range theTests {
		t.Run(theT.name, func(t *testing.T) {
			if !reflect.DeepEqual(theT.in, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("  Actual: %v", querypacksResp)
				t.Fail()
			}
		})
	}
}
