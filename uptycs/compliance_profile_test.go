package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetComplianceProfile(t *testing.T) {

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
			name:    "TestComplianceProfile",
			fixture: "fixtures/complianceProfile.json",
			id:      "bbcccdddd-eeeeef-fffff",
			out: ComplianceProfile{
				ID:                  "abbcccdddd-eeeeef-fffff",
				Name:                "Foo Threat Research",
				SeedId:              "random-string-of-characters-for-seed",
				CustomerId:          "11111111-1111-1111-1111-111111111111",
				Description:         "Test description",
				Custom:              true,
				Priority:            1,
				CreatedAt:           "2022-03-24T15:09:57.539Z",
				CreatedBy:           "",
				UpdatedAt:           "2022-03-24T15:09:57.539Z",
				UpdatedBy:           "",
				Links:     			 nil,
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/complianceProfiles/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			complianceProfileResp, err := c.GetComplianceProfile(ComplianceProfile{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(complianceProfileResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual: %v", complianceProfileResp)
				t.Fail()
			}
		})
	}
}