package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetAssetGroupRule(t *testing.T) {

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
			name:    "TestAssetGroupRule",
			fixture: "fixtures/assetGroupRule.json",
			id:      "d774ac13-ad82-4fb2-8bc9-893a1b957264",
			out: AssetGroupRule{
				ID:        "d774ac13-ad82-4fb2-8bc9-893a1b957264",
				Name:      "servers",
				Query:     "select 'servers' as value from ec2_instance_metadata\nwhere instance_id is not null\nunion\nselect 'servers' as value from gce_instance_metadata\nwhere id is not null",
				Interval:  3600,
				Platform:  "all",
				Enabled:   true,
				CreatedBy: "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				UpdatedBy: "69a61265-f2fa-40a5-a5c9-48f354533c9a",
				CreatedAt: "2021-10-05T17:35:18.425Z",
				UpdatedAt: "2022-03-30T20:18:23.016Z",
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/assetGroupRules/d774ac13-ad82-4fb2-8bc9-893a1b957264"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111/assetGroupRules"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/assetGroupRules/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			assetGroupRuleResp, err := c.GetAssetGroupRule(AssetGroupRule{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(assetGroupRuleResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", assetGroupRuleResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteAssetGroupRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   AssetGroupRule
	}

	theTests := []convTest{
		{
			name: "TestAssetGroupRule",
			in: AssetGroupRule{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/assetGroupRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteAssetGroupRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/assetGroupRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutAssetGroupRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      AssetGroupRule
	}

	theTests := []convTest{
		{
			name:    "TestAssetGroupRule",
			fixture: "fixtures/assetGroupRule.json",
			in: AssetGroupRule{
				ID:        "d774ac13-ad82-4fb2-8bc9-893a1b957264",
				Name:      "servers",
				Query:     "select 'servers' as value from ec2_instance_metadata\nwhere instance_id is not null\nunion\nselect 'servers' as value from gce_instance_metadata\nwhere id is not null",
				Interval:  3600,
				Platform:  "all",
				Enabled:   true,
				CreatedBy: "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
				UpdatedBy: "69a61265-f2fa-40a5-a5c9-48f354533c9a",
				CreatedAt: "2021-10-05T17:35:18.425Z",
				UpdatedAt: "2022-03-30T20:18:23.016Z",
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/assetGroupRules/d774ac13-ad82-4fb2-8bc9-893a1b957264"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111/assetGroupRules"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/assetGroupRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateAssetGroupRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/assetGroupRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateAssetGroupRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      AssetGroupRule
	}

	theTests := []convTest{
		{
			name:    "TestAssetGroupRule",
			fixture: "fixtures/assetGroupRule.json",
			in: AssetGroupRule{
				ID: "d774ac13-ad82-4fb2-8bc9-893a1b957264",
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/assetGroupRules/d774ac13-ad82-4fb2-8bc9-893a1b957264"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111/assetGroupRules"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/assetGroupRules",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateAssetGroupRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/assetGroupRules"], 1)
		})
	}
}
