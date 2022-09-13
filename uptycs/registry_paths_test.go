package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetRegistryPaths(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/registryPaths",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/registryPaths.json")
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

	registryPathsResp, err := c.GetRegistryPaths()
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
			in:   registryPathsResp,
			out: RegistryPaths{
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Registry path information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/registryPaths"},
					LinkItem{Rel: "parent", Title: "Customer information", Href: "/api/customers/111111111111-111111-11111-111111-111111111"},
				},
				Items: []RegistryPath{
					RegistryPath{
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
				t.Logf("  Actual: %v", registryPathsResp)
				t.Fail()
			}
		})
	}
}
