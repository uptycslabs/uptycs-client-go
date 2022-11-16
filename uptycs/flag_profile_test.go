package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetFlagProfile(t *testing.T) {

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
			name:    "TestFlagProfile",
			fixture: "fixtures/flagProfile.json",
			id:      "ed4c1240-ffcb-492a-b95f-82976c8bbab5",
			out: FlagProfile{
				ID:           "ed4c1240-ffcb-492a-b95f-82976c8bbab5",
				Custom:       false,
				Name:         "Container Events",
				Description:  "Container Events",
				Priority:     1111144,
				Flags:        "{\"add_container_image_to_events\":true,\"audit_allow_config\":true,\"disable_audit\":false,\"enable_containerd_events\":true,\"enable_docker_events\":true}",
				OsFlags:      "{}",
				CreatedBy:    "00000000-0000-0000-0000-000000000000",
				CreatedAt:    "2021-09-20T16:57:25.926Z",
				UpdatedAt:    "2022-11-07T13:06:59.670Z",
				ResourceType: "asset",
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/flagProfiles/ed4c1240-ffcb-492a-b95f-82976c8bbab5"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111/flagProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/flagProfiles/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			flagProfileResp, err := c.GetFlagProfile(FlagProfile{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(flagProfileResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", flagProfileResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteFlagProfile(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   FlagProfile
	}

	theTests := []convTest{
		{
			name: "TestFlagProfile",
			in: FlagProfile{
				ID:           "9cde7195-ec0c-475e-a208-dbf81a32798a",
				Custom:       false,
				Name:         "Container Events",
				Description:  "Container Events",
				Priority:     1111144,
				Flags:        "{\"add_container_image_to_events\":true,\"audit_allow_config\":true,\"disable_audit\":false,\"enable_containerd_events\":true,\"enable_docker_events\":true}",
				OsFlags:      "{}",
				CreatedBy:    "00000000-0000-0000-0000-000000000000",
				CreatedAt:    "2021-09-20T16:57:25.926Z",
				UpdatedAt:    "2022-11-07T13:06:59.670Z",
				ResourceType: "asset",
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/flagProfiles/ed4c1240-ffcb-492a-b95f-82976c8bbab5"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111/flagProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/flagProfiles/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteFlagProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/flagProfiles/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutFlagProfile(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      FlagProfile
	}

	theTests := []convTest{
		{
			name:    "TestFlagProfile",
			fixture: "fixtures/flagProfile.json",
			in: FlagProfile{
				ID:           "ed4c1240-ffcb-492a-b95f-82976c8bbab5",
				Custom:       false,
				Name:         "Container Events",
				Description:  "Container Events",
				Priority:     1111144,
				Flags:        "{\"add_container_image_to_events\":true,\"audit_allow_config\":true,\"disable_audit\":false,\"enable_containerd_events\":true,\"enable_docker_events\":true}",
				OsFlags:      "{}",
				CreatedBy:    "00000000-0000-0000-0000-000000000000",
				CreatedAt:    "2021-09-20T16:57:25.926Z",
				UpdatedAt:    "2022-11-07T13:06:59.670Z",
				ResourceType: "asset",
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/flagProfiles/ed4c1240-ffcb-492a-b95f-82976c8bbab5"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111/flagProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/flagProfiles/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateFlagProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/flagProfiles/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateFlagProfile(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      FlagProfile
	}

	theTests := []convTest{
		{
			name:    "TestFlagProfile",
			fixture: "fixtures/flagProfile.json",
			in: FlagProfile{
				ID:           "ed4c1240-ffcb-492a-b95f-82976c8bbab5",
				Custom:       false,
				Name:         "Container Events",
				Description:  "Container Events",
				Priority:     1111144,
				Flags:        "{\"add_container_image_to_events\":true,\"audit_allow_config\":true,\"disable_audit\":false,\"enable_containerd_events\":true,\"enable_docker_events\":true}",
				OsFlags:      "{}",
				CreatedBy:    "00000000-0000-0000-0000-000000000000",
				CreatedAt:    "2021-09-20T16:57:25.926Z",
				UpdatedAt:    "2022-11-07T13:06:59.670Z",
				ResourceType: "asset",
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/flagProfiles/ed4c1240-ffcb-492a-b95f-82976c8bbab5"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111/flagProfiles"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/flagProfiles",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateFlagProfile(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/flagProfiles"], 1)
		})
	}
}
