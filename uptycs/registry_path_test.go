package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetRegistryPath(t *testing.T) {

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
			name:    "TestRegistryPath",
			fixture: "fixtures/registryPath.json",
			id:      "b7c9c973-e2a3-4913-a755-919026267679",
			out: RegistryPath{
				ID:          "ce064913-0c00-4b14-8df3-b1dd90372f04",
				Name:        "upt_lsa_provider",
				Description: "LSA Providers",
				Grouping:    "ATTACK",
				IncludeRegistryPaths: []string{
					"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Lsa\\%",
					"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\SecurityProviders\\SecurityProviders\\%%",
				},
				RegAccesses:          false,
				ExcludeRegistryPaths: []string{},
				Custom:               true,
				CreatedBy:            "82c8c71f-12a3-40c3-ba35-b59c08ab4412",
				UpdatedBy:            "82c8c71f-12a3-40c3-ba35-b59c08ab4412",
				CreatedAt:            "2021-06-10T11:05:01.943Z",
				UpdatedAt:            "2021-06-10T11:05:01.943Z",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Registry path information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/registryPaths/ce064913-0c00-4b14-8df3-b1dd90372f04"},
					LinkItem{Rel: "parent", Title: "Registry paths information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/registryPaths"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/registryPaths/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			registryPathResp, err := c.GetRegistryPath(RegistryPath{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(registryPathResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", registryPathResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteRegistryPath(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   RegistryPath
	}

	theTests := []convTest{
		{
			name: "TestRegistryPath",
			in: RegistryPath{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/registryPaths/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteRegistryPath(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/registryPaths/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutRegistryPath(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      RegistryPath
	}

	theTests := []convTest{
		{
			name:    "TestRegistryPath",
			fixture: "fixtures/registryPath.json",
			in: RegistryPath{
				ID:          "ce064913-0c00-4b14-8df3-b1dd90372f04",
				Name:        "upt_lsa_provider",
				Description: "LSA Providers",
				Grouping:    "ATTACK",
				IncludeRegistryPaths: []string{
					"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Lsa\\%",
					"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\SecurityProviders\\SecurityProviders\\%%",
				},
				RegAccesses:          false,
				ExcludeRegistryPaths: []string{},
				Custom:               true,
				CreatedBy:            "82c8c71f-12a3-40c3-ba35-b59c08ab4412",
				UpdatedBy:            "82c8c71f-12a3-40c3-ba35-b59c08ab4412",
				CreatedAt:            "2021-06-10T11:05:01.943Z",
				UpdatedAt:            "2021-06-10T11:05:01.943Z",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Registry path information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/registryPaths/ce064913-0c00-4b14-8df3-b1dd90372f04"},
					LinkItem{Rel: "parent", Title: "Registry paths information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/registryPaths"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/registryPaths/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateRegistryPath(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/registryPaths/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateRegistryPath(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      RegistryPath
	}

	theTests := []convTest{
		{
			name:    "TestRegistryPath",
			fixture: "fixtures/registryPath.json",
			in: RegistryPath{
				ID:          "ce064913-0c00-4b14-8df3-b1dd90372f04",
				Name:        "upt_lsa_provider",
				Description: "LSA Providers",
				Grouping:    "ATTACK",
				IncludeRegistryPaths: []string{
					"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\Lsa\\%",
					"HKEY_LOCAL_MACHINE\\SYSTEM\\CurrentControlSet\\Control\\SecurityProviders\\SecurityProviders\\%%",
				},
				RegAccesses:          false,
				ExcludeRegistryPaths: []string{},
				Custom:               true,
				CreatedBy:            "82c8c71f-12a3-40c3-ba35-b59c08ab4412",
				UpdatedBy:            "82c8c71f-12a3-40c3-ba35-b59c08ab4412",
				CreatedAt:            "2021-06-10T11:05:01.943Z",
				UpdatedAt:            "2021-06-10T11:05:01.943Z",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Registry path information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/registryPaths/ce064913-0c00-4b14-8df3-b1dd90372f04"},
					LinkItem{Rel: "parent", Title: "Registry paths information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/registryPaths"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/registryPaths",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateRegistryPath(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/registryPaths"], 1)
		})
	}
}
