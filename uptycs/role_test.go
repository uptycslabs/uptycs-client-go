package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetRole(t *testing.T) {

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
			name:    "TestRole",
			fixture: "fixtures/role.json",
			id:      "b7c9c973-e2a3-4913-a755-919026267679",
			out: Role{
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
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/roles/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			roleResp, err := c.GetRole(Role{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(roleResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", roleResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteRole(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   Role
	}

	theTests := []convTest{
		{
			name: "TestRole",
			in: Role{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/roles/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteRole(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/roles/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutRole(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Role
	}

	theTests := []convTest{
		{
			name:    "TestRole",
			fixture: "fixtures/roleCreate.json",
			in: Role{
				ID:               "ac6ef928-36a5-4388-86d4-575ba1085e7d",
				Name:             "Monitoring Profile",
				Description:      "Role for Uptycs monitoring team to perform necessary tasks",
				Permissions:      []string{},
				RoleObjectGroups: []ObjectGroup{},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/roles/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateRole(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/roles/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateRole(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Role
	}

	theTests := []convTest{
		{
			name:    "TestRole",
			fixture: "fixtures/roleCreate.json",
			in: Role{
				ID:                   "ac6ef928-36a5-4388-86d4-575ba1085e7d",
				Name:                 "Monitoring Profile",
				Description:          "Role for Uptycs monitoring team to perform necessary tasks",
				Permissions:          []string{"EXCEPTION:READ", "EXCEPTION:CREATE"},
				RoleObjectGroups:     []ObjectGroup{},
				Custom:               true,
				Hidden:               false,
				CreatedBy:            "61b98805-54ea-40d9-89b7-f8bf7780666c",
				UpdatedBy:            "61b98805-54ea-40d9-89b7-f8bf7780666c",
				CreatedAt:            "2022-08-25T13:54:35.768Z",
				UpdatedAt:            "2022-08-25T14:46:48.409Z",
				NoMinimalPermissions: true,
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/roles",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateRole(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/roles"], 1)
		})
	}
}
