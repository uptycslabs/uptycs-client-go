package uptycs

import (
	"github.com/jarcoal/httpmock"
	"net/http"
	"reflect"
	"testing"
)

func TestGetEventExcludeProfiles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://uptycs.foo/public/api/customers/d/eventExcludeProfiles",
		func(req *http.Request) (*http.Response, error) {
			fixture, err := RespFromFixture("fixtures/eventExcludeProfiles.json")
			if err != nil {
				t.Errorf(err.Error())
			}
			return fixture, err
		},
	)

	c, _ := NewClient(UptycsConfig{
		Host:       "https://uptycs.foo",
		ApiKey:     "b",
		ApiSecret:  "c",
		CustomerID: "d",
	})

	eventExcludeProfilesResp, err := c.GetEventExcludeProfiles()
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
			in:   eventExcludeProfilesResp,
			out: EventExcludeProfiles{
				Items: []EventExcludeProfile{
					EventExcludeProfile{
						ID:          "13da8bc4-3c70-4bb9-a4d4-7ca320860926",
						CustomerID:  "11111111-1111-1111-1111-111111111111",
						Name:        "Uptycs default event exclude profile",
						Description: "Filters known paths, IPs and domains",
						Priority:    20200227,
						Metadata: EventExcludeProfileMetadata{
							SocketEvents: SocketEvents{
								RemoteAddress: []string{
									"^S",
									"^::S",
									"^::1S",
									"^0000:0000:0000:0000:0000:0000:0000:.*S",
									"^0000:0000:0000:0000:0000:ffff:0a.*S",
									"^0000:0000:0000:0000:0000:ffff:7f00:.*S",
									"^0000:0000:0000:0000:0000:ffff:ac10:.*",
									"^2001:4860:4860::8888S",
									"^2001:4860:4860::8844S",
									"^::ffff:127\\.0\\.0\\.1S",
									"^0S",
									"^0\\..*S",
									"^8\\.8\\.8\\.8S",
									"^8\\.8\\.4\\.4S",
									"^1\\.1\\.1\\.1S",
									"^1\\.0\\.0\\.1S",
									"^0\\.",
									"^10\\..*S",
									"^100\\.6[4-9]\\..*S",
									"^100\\.[7-9]\\d\\..*S",
									"^100\\.1[0-1]\\d\\..*S",
									"^100\\.12[0-7]\\..*S",
									"^127\\..*S",
									"^169\\.254\\..*S",
									"^172\\.1[6-9]\\..*S",
									"^172\\.2[0-9]\\..*S",
									"^172\\.3[0-1]\\..*S",
									"^192\\.168\\..*S",
									"^22[4-9]\\..*S",
									"^23[0-9]\\..*S",
									"^24[0-9]\\..*S",
									"^25[0-5]\\..*S",
									"^-1S",
								},
							},
							ProcessEvents: ProcessEvents{
								Path: []string{
									"^.*ntp\\.orgS",
									"^/bin/bashS",
									"^/bin/catS",
									"^/bin/dashS",
									"^/bin/dateS",
									"^/bin/grepS",
									"^/bin/hostnameS",
									"^/bin/lsS",
									"^/bin/mkdirS",
									"^/bin/pgrepS",
									"^/bin/psS",
									"^/bin/sedS",
									"^/bin/sleepS",
									"^/bin/systemctlS",
									"^/bin/unameS",
									"^/usr/bin/basenameS",
									"^/usr/bin/curlS",
									"^/usr/bin/cutS",
									"^/usr/bin/duS",
									"^/usr/bin/gawkS",
									"^/usr/bin/perlS",
									"^/usr/bin/pgrepS",
									"^/usr/bin/pingS",
									"^/usr/bin/python2.7S",
									"^/usr/bin/python3.5S",
									"^/usr/bin/sleepS",
									"^/usr/bin/statS",
									"^/usr/bin/tailS",
									"^/usr/bin/teeS",
									"^/usr/bin/topS",
									"^/usr/bin/trS",
									"^/usr/bin/wcS",
									"^/bin/busyboxS",
									"^/usr/bin/runcS",
									"^/usr/lib/jvm/java-1.8-openjdk/bin/javaS",
								},
							},
							ProcessFileEvents: ProcessFileEvents{
								Path: []string{
									"\\.uptycs.ioS",
									".*(cache|notification|localstate|resource|safety|automaticdestination|packages|tempstate|((edb|cf|gthr|dir|crwl|exd|db-journal|aodl|evtx|json|dat|log|tmp|etl|db|ini|xml|chk|jfm|pf|temp))S)",
								},
								Operation: []string{
									"^open\\+readS",
									"attributes_modified",
									"unlink",
								},
								Executable: []string{
									"^.*osqueryd\\.exeS|^.*collectguestlogs\\.exeS|^.*MsMpEng\\.exeS",
								},
							},
							UserEvents: UserEvents{
								Message: []string{
									"^cwd=.*",
									"^op=PAM:accountingS",
									"^op=PAM:session_closeS",
									"^op=PAM:session_openS",
									"^op=PAM:setcredS",
								},
							},
							RegistryEvents: RegistryEvents{
								Action: []string{
									"SET_INFORMATION",
									"CREATED",
								},
							},
							DnsLookupEvents: DnsLookupEvents{
								Answer: []string{
									"^(?![\\s\\S])",
								},
								Question: []string{
									"^0\\.pool\\.ntp\\.orgS",
									"^1\\.pool\\.ntp\\.orgS",
									"^2\\.pool\\.ntp\\.orgS",
									"^3\\.pool\\.ntp\\.orgS",
									"^localhostS",
									"^.*\\.compute\\.internalS",
									"^.*\\.amazonaws\\.comS",
									"^.*\\.googleapis\\.comS",
									"^.*\\.datadoghq\\.comS",
									"^.*\\.internal\\.mxS",
									"^.*\\.ntp\\.orgS",
									"^.*\\.uptycs\\.ioS",
									"^encrypted\\.debug\\.opendns\\.comS",
								},
							},
						},
						CreatedAt:    "2022-06-27T04:09:05.342Z",
						CreatedBy:    "00000000-0000-0000-0000-000000000000",
						UpdatedAt:    "2022-06-27T04:09:05.342Z",
						UpdatedBy:    "00000000-0000-0000-0000-000000000000",
						ResourceType: "asset",
						Platform:     "all",
						Links: []LinkItem{
							LinkItem{Rel: "self", Title: "Event exlude profile information", Href: "/api/customers/11111111-1111-1111-1111-111111111111/eventExcludeProfiles/13da8bc4-3c70-4bb9-a4d4-7ca320860926"},
							LinkItem{Rel: "parent", Title: "Event exlude profiles information", Href: "/api/customers/11111111-1111-1111-1111-111111111111/eventExcludeProfiles"},
						},
					},
				},
				Links: []LinkItem{
					LinkItem{Rel: "self", Title: "Event exclude profile information", Href: "/api/customers/11111111-1111-1111-1111-111111111111/eventExcludeProfiles"},
					LinkItem{Rel: "parent", Title: "Customer information", Href: "/api/customers/11111111-1111-1111-1111-111111111111"},
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
				t.Logf("  Actual: %v", eventExcludeProfilesResp)
				t.Fail()
			}
		})
	}
}
