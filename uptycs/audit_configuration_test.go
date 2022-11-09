package uptycs

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetAuditConfiguration(t *testing.T) {

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
			name:    "TestAuditConfiguration",
			fixture: "fixtures/auditConfiguration.json",
			id:      "b7c9c973-e2a3-4913-a755-919026267679",
			out: AuditConfiguration{
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
				Checks:      0,
				AuditEntry: []AuditEntry{
					{
						ID:                   "00ce43ae-e106-498e-a6f9-8a30ae845e9c",
						AuditConfigurationID: "7d51a844-f28e-4dbf-8831-e4a063e16156",
						AuditName: []string{
							"CIS",
						},
						Standard:            "AWS",
						Version:             "1.4.0",
						Section:             "1.1",
						Title:               "Maintain current contact details",
						Scored:              false,
						Level:               "Level 1",
						Description:         "Ensure contact email and telephone details for AWS accounts are current and map to more\nthan one individual in your organization.\nAn AWS account supports a number of contact details, and AWS will use these to contact\nthe account owner if activity judged to be in breach of Acceptable Use Policy or indicative\nof likely security compromise is observed by the AWS Abuse team. Contact details should\nnot be for a single individual, as circumstances may arise where that individual is\nunavailable. Email contact details should point to a mail alias which forwards email to\nmultiple individuals within the organization; where feasible, phone contact details should\npoint to a PABX hunt group or other call-forwarding system.",
						Rationale:           "If an AWS account is observed to be behaving in a prohibited or suspicious manner, AWS\nwill attempt to contact the account owner by email and phone using the contact details\nlisted. If this is unsuccessful and the account behavior needs urgent mitigation, proactive\nmeasures may be taken, including throttling of traffic between the account exhibiting\nsuspicious behavior and the AWS API endpoints and the Internet. This will result in\nimpaired service to and from the account in question, so it is in both the customers' and\nAWS' best interests that prompt contact can be established. This is best achieved by setting\nAWS account contact details to point to resources which have multiple individuals as\nrecipients, such as email aliases and PABX hunt groups.",
						Command:             "This activity can only be performed via the AWS Console, with a user who has permission\nto read and write Billing information (aws-portal:*Billing )\n1. Sign in to the AWS Management Console and open the Billing and Cost Management\nconsole at https://console.aws.amazon.com/billing/home#/.\n2. On the navigation bar, choose your account name, and then choose My Account.\n3. On the Account Settings page, review and verify the current details.\n4. Under Contact Information, review and verify the current details.",
						Remediation:         "This activity can only be performed via the AWS Console, with a user who has permission\nto read and write Billing information (aws-portal:*Billing ).\n1. Sign in to the AWS Management Console and open the Billing and Cost Management\nconsole at https://console.aws.amazon.com/billing/home#/.\n2. On the navigation bar, choose your account name, and then choose My Account.\n3. On the Account Settings page, next to Account Settings, choose Edit.\n4. Next to the field that you need to update, choose Edit.\n5. After you have entered your changes, choose Save changes.\n6. After you have made your changes, choose Done.\n7. To edit your contact information, under Contact Information, choose Edit.\n8. For the fields that you want to change, type your updated information, and then\nchoose Update.",
						ExpectedValue:       "This activity can only be performed via the AWS Console, with a user who has permission\nto read and write Billing information (aws-portal:*Billing )\n1. Sign in to the AWS Management Console and open the Billing and Cost Management\nconsole at https://console.aws.amazon.com/billing/home#/.\n2. On the navigation bar, choose your account name, and then choose My Account.\n3. On the Account Settings page, review and verify the current details.\n4. Under Contact Information, review and verify the current details.",
						AuthoritativeSource: "CIS Amazon Web Services Foundations",
						Exception:           "Manual",
						Chapter:             "1 Identity and Access Management",
						CheckID:             "",
						Enabled:             true,
						Service:             "IAM",
						CreatedBy:           "61b98805-54ea-40d9-89b7-f8bf7780666c",
						//Score: null,
						UpdatedBy:   "61b98805-54ea-40d9-89b7-f8bf7780666c",
						RunCategory: 1,
						Timeout:     3600,
						CreatedAt:   "2022-05-09T16:42:48.432Z",
						UpdatedAt:   "2022-05-09T16:42:48.432Z",
						IsManual:    true,
						//Parameters: null,
						//RemediationAction: null
					}},
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Compliance configuration information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/auditConfigurations/7d51a844-f28e-4dbf-8831-e4a063e16156"},
					LinkItem{Rel: "parent", Title: "Compliance configurations information", Href: "/api/customers/111111111111-111111-11111-111111-111111111/auditConfigurations"},
				},
			},
		},
	}

	for _, theT := range theTests {
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()

		t.Run(theT.name, func(t *testing.T) {
			httpmock.RegisterResponder("GET", fmt.Sprintf("https://uptycs.foo/public/api/customers/d/auditConfigurations/%v", theT.id),
				func(req *http.Request) (*http.Response, error) {
					fixture, err := RespFromFixture(theT.fixture)
					if err != nil {
						t.Errorf(err.Error())
					}
					return fixture, err
				},
			)

			auditConfigurationResp, err := c.GetAuditConfiguration(AuditConfiguration{
				ID: theT.id,
			})

			if err != nil {
				t.Errorf(err.Error())
			}

			if !reflect.DeepEqual(auditConfigurationResp, theT.out) {
				t.Log("Output does not match expected")
				t.Logf("Expected: %v", theT.out)
				t.Logf("Actual:   %v", auditConfigurationResp)
				t.Fail()
			}
		})
	}
}
