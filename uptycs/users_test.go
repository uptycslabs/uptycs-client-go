package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/users",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/users.json")
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

	usersResp, err := c.GetUsers()
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
			in:   usersResp,
			out: Users{
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Users information", Href: "/api/customers/11111111-1111-1111-1111-111111111111/users"},
				},
				Items: []User{
					User{
						ID:                  "aaadd6ed-b85f-42fe-b0ec-01f89abf7249",
						Name:                "Foo Threat Research",
						Email:               "atest@example.com",
						Active:              false,
						SuperAdmin:          false,
						Bot:                 false,
						Support:             false,
						PriorLogin:          false,
						ImageURL:            "",
						CreatedAt:           "2021-06-09T02:53:53.183Z",
						MaxIdleTimeMins:     30,
						AlertHiddenColumns:  nil,
						UpdatedAt:           "2022-03-24T15:09:57.539Z",
						LastUpdatedByUptycs: "2022-03-24T15:09:57.539Z",
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
				t.Logf("Actual:   %v", usersResp)
				t.Fail()
			}
		})
	}
}
