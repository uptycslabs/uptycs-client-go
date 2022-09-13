package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTag(t *testing.T) {

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
			name:    "TestTag",
			fixture: "fixtures/tag.json",
			id:      "9dfc53a4-bf1e-4efe-8f04-e1fd8802e9e3",
			out: Tag{
				ID:                   "9dfc53a4-bf1e-4efe-8f04-e1fd8802e9e3",
				Key:                  "636776063332",
				CreatedBy:            "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				Tag:                  "636776063332",
				Custom:               true,
				System:               false,
				CreatedAt:            "2022-02-03T05:06:53.918Z",
				Status:               "active",
				Source:               "direct",
				UpdatedAt:            "2022-02-03T05:06:53.918Z",
				ResourceType:         "asset",
				FilePathGroups:       []TagConfigurationObject{},
				EventExcludeProfiles: []TagConfigurationObject{},
				RegistryPaths:        []TagConfigurationObject{},
				Querypacks:           []TagConfigurationObject{},
				YaraGroupRules:       []TagConfigurationObject{},
				AuditConfigurations:  []TagConfigurationObject{},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/tags/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			tagResp, err := c.GetTag(Tag{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(tagResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", tagResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteTag(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   Tag
	}

	theTests := []convTest{
		{
			name: "TestTag",
			in: Tag{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/tags/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteTag(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/tags/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutTag(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Tag
	}

	theTests := []convTest{
		{
			name:    "TestTag",
			fixture: "fixtures/tag.json",
			in: Tag{
				ID:                   "9dfc53a4-bf1e-4efe-8f04-e1fd8802e9e3",
				Key:                  "636776063332",
				CreatedBy:            "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				Tag:                  "636776063332",
				Custom:               true,
				System:               false,
				CreatedAt:            "2022-02-03T05:06:53.918Z",
				Status:               "active",
				Source:               "direct",
				UpdatedAt:            "2022-02-03T05:06:53.918Z",
				ResourceType:         "asset",
				FilePathGroups:       []TagConfigurationObject{},
				EventExcludeProfiles: []TagConfigurationObject{},
				RegistryPaths:        []TagConfigurationObject{},
				Querypacks:           []TagConfigurationObject{},
				YaraGroupRules:       []TagConfigurationObject{},
				AuditConfigurations:  []TagConfigurationObject{},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/tags/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateTag(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/tags/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateTag(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      Tag
	}

	theTests := []convTest{
		{
			name:    "TestTag",
			fixture: "fixtures/tag.json",
			in: Tag{
				ID:                   "9dfc53a4-bf1e-4efe-8f04-e1fd8802e9e3",
				Key:                  "636776063332",
				CreatedBy:            "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				Tag:                  "636776063332",
				Custom:               true,
				System:               false,
				CreatedAt:            "2022-02-03T05:06:53.918Z",
				Status:               "active",
				Source:               "direct",
				UpdatedAt:            "2022-02-03T05:06:53.918Z",
				ResourceType:         "asset",
				FilePathGroups:       []TagConfigurationObject{},
				EventExcludeProfiles: []TagConfigurationObject{},
				RegistryPaths:        []TagConfigurationObject{},
				Querypacks:           []TagConfigurationObject{},
				YaraGroupRules:       []TagConfigurationObject{},
				AuditConfigurations:  []TagConfigurationObject{},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/tags",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateTag(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/tags"], 1)
		})
	}
}
