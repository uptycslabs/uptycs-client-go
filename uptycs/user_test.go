package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {

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
			name:    "TestUser",
			fixture: "fixtures/user.json",
			id:      "ac6ef928-36a5-4388-86d4-575ba1085e7d",
			out: User{
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
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/users/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			userResp, err := c.GetUser(User{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(userResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", userResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   User
	}

	theTests := []convTest{
		{
			name: "TestUser",
			in: User{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/users/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteUser(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/users/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutUser(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      User
	}

	theTests := []convTest{
		{
			name:    "TestUser",
			fixture: "fixtures/userCreate.json",
			in: User{
				ID:                  "aaadd6ed-b85f-42fe-b0ec-01f89abf7249",
				Name:                "Foo Threat Research",
				Email:               "atest@example.com",
				Phone:               "111-111-1111",
				Active:              false,
				SuperAdmin:          false,
				Bot:                 false,
				Support:             false,
				PriorLogin:          false,
				ImageURL:            "asdf",
				CreatedAt:           "2021-06-09T02:53:53.183Z",
				MaxIdleTimeMins:     30,
				AlertHiddenColumns:  []string{},
				Roles:               []Role{},
				UpdatedAt:           "2022-03-24T15:09:57.539Z",
				LastUpdatedByUptycs: "2022-03-24T15:09:57.539Z",
				UserObjectGroups:    []ObjectGroup{},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/users/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateUser(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/users/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateUser(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      User
	}

	theTests := []convTest{
		{
			name:    "TestUser",
			fixture: "fixtures/userCreate.json",
			in: User{
				ID:                  "aaadd6ed-b85f-42fe-b0ec-01f89abf7249",
				Name:                "Foo Threat Research",
				Email:               "atest@example.com",
				Phone:               "111-111-1111",
				Active:              false,
				SuperAdmin:          false,
				Bot:                 false,
				Support:             false,
				PriorLogin:          false,
				ImageURL:            "asdf",
				CreatedAt:           "2021-06-09T02:53:53.183Z",
				MaxIdleTimeMins:     30,
				AlertHiddenColumns:  []string{},
				Roles:               []Role{},
				UpdatedAt:           "2022-03-24T15:09:57.539Z",
				LastUpdatedByUptycs: "2022-03-24T15:09:57.539Z",
				UserObjectGroups:    []ObjectGroup{},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/users",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateUser(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/users"], 1)
		})
	}
}
