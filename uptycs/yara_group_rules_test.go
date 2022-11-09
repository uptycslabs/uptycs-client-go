package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetYaraGroupRules(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/yaraGroupRules",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/yaraGroupRules.json")
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

	yaraGroupRulesResp, err := c.GetYaraGroupRules()
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
			in:   yaraGroupRulesResp,
			out: YaraGroupRules{
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "YaraGroupRules information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/yaraGroupRules"},
					LinkItem{Rel: "parent", Title: "Customer information", Href: "/api/customers/111111111111-111111-11111-111111-111111111"},
				},
				Items: []YaraGroupRule{
					YaraGroupRule{
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
				t.Logf("  Actual:   %v", yaraGroupRulesResp)
				t.Fail()
			}
		})
	}
}
