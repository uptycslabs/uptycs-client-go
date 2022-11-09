package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAuditConfigurations(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/auditConfigurations",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/auditConfigurations.json")
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

	auditConfigurationsResp, err := c.GetAuditConfigurations()
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
			in:   auditConfigurationsResp,
			out: AuditConfigurations{
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Compliance configuration configurations information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/auditConfigurations"},
					LinkItem{Rel: "parent", Title: "Customer information", Href: "/api/customers/111111111111-111111-11111-111111-111111111"},
				},
				Items: []AuditConfiguration{
					AuditConfiguration{
						ID:          "7d51a844-f28e-4dbf-8831-e4a063e16156",
						Name:        "CIS_AWS_Benchmark_v140-CLONE",
						Description: "",
						Framework:   "CIS",
						Version:     "1.4.0",
						OsVersion:   "any",
						Platform:    "aws",
						TableName:   "cis_aws",
						Sha256:      "3eed2c5ed217c5b0c44159471124e450e4e778643e15892a5e7d1e8db8d27185",
						CreatedBy:   "61b98805-54ea-40d9-89b7-f8bf7780666c",
						UpdatedBy:   "61b98805-54ea-40d9-89b7-f8bf7780666c",
						CreatedAt:   "2022-05-09T16:42:48.423Z",
						UpdatedAt:   "2022-07-18T05:34:08.673Z",
						Type:        "cloud",
						Checks:      58,
						Links: []LinkItem{
							LinkItem{Rel: "self", Title: "Compliance configuration information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/auditConfigurations/7d51a844-f28e-4dbf-8831-e4a063e16156"},
							LinkItem{Rel: "parent", Title: "Compliance configurations information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/auditConfigurations"},
						},
					},
					AuditConfiguration{
						ID:          "66bd39f5-e925-43a5-82be-12c0f3b90ce0",
						Name:        "CIS_Distribution_Independent_Linux_Benchmark_v200-Test",
						Description: "",
						Framework:   "CIS",
						Version:     "2.0.0",
						OsVersion:   "any",
						Platform:    "linux",
						TableName:   "cis_independent_linux",
						Sha256:      "fc02751eebec54570c1152dff9bc0a538e8cd624f41943f9f4be478977d709f3",
						CreatedBy:   "d5934c11-034b-4d8f-a4fd-e3c7aff83abc",
						UpdatedBy:   "d5934c11-034b-4d8f-a4fd-e3c7aff83abc",
						CreatedAt:   "2021-06-22T21:12:41.875Z",
						UpdatedAt:   "2022-06-16T04:48:27.414Z",
						Type:        "host",
						Checks:      237,
						Links: []LinkItem{
							LinkItem{Rel: "self", Title: "Compliance configuration information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/auditConfigurations/66bd39f5-e925-43a5-82be-12c0f3b90ce0"},
							LinkItem{Rel: "parent", Title: "Compliance configurations information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/auditConfigurations"},
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
				t.Logf("  Actual:   %v", auditConfigurationsResp)
				t.Fail()
			}
		})
	}
}
