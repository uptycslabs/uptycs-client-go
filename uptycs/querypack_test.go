package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetQuerypack(t *testing.T) {

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
			name:    "TestQuerypack",
			fixture: "fixtures/querypack.json",
			id:      "ac4cea47-17d4-4199-8fa8-f658355c36ae",
			out: Querypack{
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
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/querypacks/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			querypackResp, err := c.GetQuerypack(Querypack{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(querypackResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", querypackResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteQuerypack(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   Querypack
	}

	theTests := []convTest{
		{
			name: "TestQuerypack",
			in: Querypack{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/querypacks/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteQuerypack(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/querypacks/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutQuerypack(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Querypack
	}

	theTests := []convTest{
		{
			name:    "TestQuerypack",
			fixture: "fixtures/querypack.json",
			in: Querypack{
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
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/querypacks/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateQuerypack(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/querypacks/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateQuerypack(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Querypack
	}

	theTests := []convTest{
		{
			name:    "TestQuerypack",
			fixture: "fixtures/querypack.json",
			in: Querypack{
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
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/querypacks",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateQuerypack(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/querypacks"], 1)
		})
	}
}
