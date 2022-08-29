package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetTagRules(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/tagRules",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/tagRules.json")
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

	tagRulesResp, err := c.GetTagRules()
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
			in:   tagRulesResp,
			out: TagRules{
				Links: []LinkItem{
					LinkItem{Rel: "self", Href: "/api/customers/111111111111-111111-11111-111111-111111111/tagRules"},
					LinkItem{Rel: "parent", Href: "/api/customers/111111111111-111111-11111-111111-111111111"},
				},
				Items: []TagRule{
					TagRule{
						ID:           "54755880-2f1a-430f-9c52-4787c3b24fb7",
						Name:         "AWS Account ID Tag",
						Description:  "Tag assets with their AWS account ID",
						Query:        "select owner_id as tag, upt_asset_id from upt_cloud_instance_inventory;",
						Source:       "realtime",
						RunOnce:      false,
						Interval:     60,
						Platform:     "linux",
						SeedID:       "d4004ffc-74a9-436c-bf25-070c32032db2",
						Enabled:      false,
						System:       false,
						CreatedBy:    "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
						UpdatedBy:    "66a9a52c-5fa0-4cf4-abe7-da5504f67950",
						CreatedAt:    "2022-02-03T04:48:09.945Z",
						UpdatedAt:    "2022-07-11T19:28:16.532Z",
						ResourceType: "asset",
					},
					TagRule{
						ID:           "3b231f3a-f852-4d86-8611-2d3a1a368851",
						Name:         "Log4jshell tag rule",
						Description:  "Auto tag all assets",
						Query:        "SELECT\n\tCASE\n  \tWHEN vercmp(split(osquery_version, '-Uptycs')[1], '5.0.1.22') < 0 THEN 'upt-log4shell-exploit-tracker-osq-pre-5.0.1.22'\n  \tWHEN vercmp(split(osquery_version, '-Uptycs')[1], '5.0.1.22') >= 0 THEN 'upt-log4shell-exploit-tracker-osq-post-5.0.1.22'\n\tEND\n\t\tAS tag,\n\tupt_asset_id\nFROM\n\tupt_assets;",
						Source:       "global",
						RunOnce:      false,
						Interval:     3600,
						Platform:     "all",
						SeedID:       "614aba88-e7b5-4eaf-bcac-c0b39e98f77d",
						Enabled:      false,
						System:       false,
						CreatedBy:    "00000000-0000-0000-0000-000000000000",
						UpdatedBy:    "00000000-0000-0000-0000-000000000000",
						CreatedAt:    "2022-01-03T05:34:17.964Z",
						UpdatedAt:    "2022-04-09T05:59:46.315Z",
						ResourceType: "asset",
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
				t.Logf("  Actual: %v", tagRulesResp)
				t.Fail()
			}
		})
	}
}
