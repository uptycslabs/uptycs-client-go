package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetTags(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/tags",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/tags.json")
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

	tagsResp, err := c.GetTags()
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
			in:   tagsResp,
			out: Tags{
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Tags information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/tags"},
					LinkItem{Rel: "parent", Title: "Customer information", Href: "/api/customers/111111111111-111111-11111-111111-111111111"},
				},
				Items: []Tag{
					Tag{
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
					Tag{
						ID:                   "925dd20d-ec90-4c27-9f95-afc7ec6bdc14",
						Key:                  "ads-1",
						CreatedBy:            "00000000-0000-0000-0000-000000000000",
						Tag:                  "ads-1",
						Custom:               false,
						System:               false,
						CreatedAt:            "2022-08-01T16:01:08.853Z",
						Status:               "active",
						Source:               "direct",
						UpdatedAt:            "2022-08-01T16:01:08.857Z",
						ResourceType:         "kubernetes",
						FilePathGroups:       []TagConfigurationObject{},
						EventExcludeProfiles: []TagConfigurationObject{},
						RegistryPaths:        []TagConfigurationObject{},
						Querypacks:           []TagConfigurationObject{},
						YaraGroupRules:       []TagConfigurationObject{},
						AuditConfigurations:  []TagConfigurationObject{},
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
				t.Logf("Actual:   %v", tagsResp)
				t.Fail()
			}
		})
	}
}
