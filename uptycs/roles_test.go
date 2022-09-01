package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetRoles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/roles",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/roles.json")
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

	rolesResp, err := c.GetRoles()
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
			in:   rolesResp,
			out: Roles{
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/11111111-1111-1111-1111-111111111111/roles"},
					LinkItem{Rel: "parent", Href: "/api/customers/11111111-1111-1111-1111-111111111111"},
				},
				Items: []Role{
					Role{
						ID:                   "ac6ef928-36a5-4388-86d4-575ba1085e7d",
						Name:                 "Monitoring Profile",
						Description:          "Role for Uptycs monitoring team to perform necessary tasks",
						Permissions:          []string{"EXCEPTION:READ", "EXCEPTION:CREATE"},
						Custom:               true,
						Hidden:               false,
						CreatedBy:            "61b98805-54ea-40d9-89b7-f8bf7780666c",
						UpdatedBy:            "61b98805-54ea-40d9-89b7-f8bf7780666c",
						CreatedAt:            "2022-08-25T13:54:35.768Z",
						UpdatedAt:            "2022-08-25T14:46:48.409Z",
						NoMinimalPermissions: true,
					},
					Role{
						ID:                   "baeb925d-ea1f-44ab-a92a-cc0a5a985cb9",
						Name:                 "admin",
						Description:          "Default admin role",
						Permissions:          []string{"OSQUERY:DOWNLOAD", "OSQUERY:READ"},
						Custom:               false,
						Hidden:               false,
						CreatedAt:            "2021-06-01T17:39:09.109Z",
						UpdatedAt:            "2022-08-16T07:35:44.443Z",
						NoMinimalPermissions: false,
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
				t.Logf("Input: %v", theT.in)
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", rolesResp)
				t.Fail()
			}
		})
	}
}
