package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestGetYaraGroupRule(t *testing.T) {

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
			name:    "TestYaraGroupRule",
			fixture: "fixtures/yaraGroupRule.json",
			id:      "b7c9c973-e2a3-4913-a755-919026267679",
			out: YaraGroupRule{
				ID:          "9a5a3262-ee74-417c-ade0-c1948ec8bc27",
				Name:        "AmazonAccessKeyId",
				Description: "amazon access key id",
				Rules:       "rule AwsAccessKeyIdRule : AWS\n{\n    meta:\n        name = \"AWS Access Key ID\"\n        author = \"github.com/pseudo-security\"\n        date = \"2020-01-01\"\n\n        /* Test Cases */\n        test_match_1 = \"AKIA00TESTAWSIDKEY00\"\n\n    strings:\n        $ = /(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}/ fullword\n\n    condition:\n        any of them\n}",
				Custom:      false,
				CreatedBy:   "00000000-0000-0000-0000-000000000000",
				UpdatedBy:   "00000000-0000-0000-0000-000000000000",
				CreatedAt:   "2022-01-17T08:12:16.882Z",
				UpdatedAt:   "2022-04-09T06:07:16.431Z",
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "YaraGroupRule information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/yaraGroupRules/9a5a3262-ee74-417c-ade0-c1948ec8bc27"},
					LinkItem{Rel: "parent", Title: "YaraGroupRules information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/yaraGroupRules"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/yaraGroupRules/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			yaraGroupRuleResp, err := c.GetYaraGroupRule(YaraGroupRule{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(yaraGroupRuleResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", yaraGroupRuleResp)
				t.Fail()
			}
		})
	}
}

func TestDeleteYaraGroupRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name string
		in   YaraGroupRule
	}

	theTests := []convTest{
		{
			name: "TestYaraGroupRule",
			in: YaraGroupRule{
				ID: "9cde7195-ec0c-475e-a208-dbf81a32798a",
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("DELETE", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/yaraGroupRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					resp, err := httpmock.NewJsonResponse(200, "{}")
					if err != nil {
						t.Errorf(err.Error())
					}
					return resp, err
				},
			)

			_, err := c.DeleteYaraGroupRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("DELETE https://uptycs.foo/public/api/customers/d/yaraGroupRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestPutYaraGroupRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      YaraGroupRule
	}

	theTests := []convTest{
		{
			name:    "TestYaraGroupRule",
			fixture: "fixtures/yaraGroupRule.json",
			in:      YaraGroupRule{},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("PUT", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/yaraGroupRules/%v", theT.in.ID),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.UpdateYaraGroupRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo[fmt.Sprintf("PUT https://uptycs.foo/public/api/customers/d/yaraGroupRules/%v", theT.in.ID)], 1)
			// TODO: assert the body that was intercepted by the mock
		})
	}
}

func TestCreateYaraGroupRule(t *testing.T) {

	c, _ := NewClient(Config{
		Host:       "https://uptycs.foo",
		APIKey:     "b",
		APISecret:  "c",
		CustomerID: "d",
	})

	type convTest struct {
		name    string
		fixture string
		in      YaraGroupRule
	}

	theTests := []convTest{
		{
			name:    "TestYaraGroupRule",
			fixture: "fixtures/yaraGroupRule.json",
			in:      YaraGroupRule{},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("POST", "https://uptycs.foo/public/api/customers/d/yaraGroupRules",
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			_, err := c.CreateYaraGroupRule(theT.in)
			if err != nil {
				t.Errorf(err.Error())
			}
			countInfo := httpmock.GetCallCountInfo()

			assert.Equal(t, countInfo["POST https://uptycs.foo/public/api/customers/d/yaraGroupRules"], 1)
		})
	}
}
