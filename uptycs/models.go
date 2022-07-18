package uptycs

import (
	"encoding/json"
	"errors"
	"strconv"
)

var (
	// ErrUnsupportedType is returned if the type is not implemented
	ErrUnsupportedType = errors.New("unsupported type")
)

type EventRules struct {
	Links  []LinkItem  `json:"links,omitempty"`
	Items  []EventRule `json:"items,omitempty"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type ScriptConfig struct {
	ID          string `json:"id,omitempty"`
	CustomerID  string `json:"customerId,omitempty"`
	QueryPackID string `json:"querypackId,omitempty"`
	TableName   string `json:"tableName,omitempty"`
	Added       bool   `json:"added,omitempty"`
}

type EventRule struct {
	ID                string        `json:"id,omitempty"`
	CustomerID        string        `json:"customerId,omitempty"`
	SeedID            string        `json:"seedId,omitempty"`
	Name              string        `json:"name,omitempty"`
	Description       string        `json:"description,omitempty"`
	Code              string        `json:"code,omitempty"`
	Type              string        `json:"type,omitempty"`
	Rule              string        `json:"rule,omitempty"`
	Grouping          string        `json:"grouping,omitempty"`
	Enabled           bool          `json:"enabled,omitempty"`
	Custom            bool          `json:"custom,omitempty"`
	CreatedAt         string        `json:"createdAt,omitempty"`
	IsInternal        bool          `json:"isInternal,omitempty"`
	EventTags         []string      `json:"eventTags,omitempty"`
	CreatedBy         string        `json:"createdBy,omitempty"`
	UpdatedAt         string        `json:"updatedAt,omitempty"`
	UpdatedBy         string        `json:"updatedBy,omitempty"`
	GroupingL2        string        `json:"groupingL2,omitempty"`
	GroupingL3        string        `json:"groupingL3,omitempty"`
	Score             string        `json:"score,omitempty"`
	Lock              bool          `json:"lock,omitempty"`
	ScriptConfig      *ScriptConfig `json:"scriptConfig,omitempty"`
	SQLConfig         *SQLConfig    `json:"sqlConfig,omitempty"`
	BuilderConfig     BuilderConfig `json:"builderConfig"`
	BuilderConfigJson string        `json:"builderConfigJson,omitempty"`
	Links             []LinkItem    `json:"links"`
}

type BuilderConfigFilter struct {
	And             []BuilderConfigFilter `json:"and,omitempty"`
	Or              []BuilderConfigFilter `json:"or,omitempty"`
	Not             bool                  `json:"not,omitempty"`
	Name            string                `json:"name,omitempty"`
	Value           ArrayOrString         `json:"value,omitempty"`
	Operator        string                `json:"operator,omitempty"`
	IsDate          bool                  `json:"isDate,omitempty"`
	IsVersion       bool                  `json:"isVersion,omitempty"`
	IsWordMatch     bool                  `json:"isWordMatch,omitempty"`
	CaseInsensitive bool                  `json:"caseInsensitive,omitempty"`
}

type ArrayOrString []string

func (aos *ArrayOrString) UnmarshalJSON(raw []byte) error {
	var dyval interface{}
	err := json.Unmarshal(raw, &dyval)
	if err != nil {
		return err
	}

	// Have to figure out the type and parse appropriately
	// For now everything is a string, so turn all values
	// (float64, bool, string, array of [float64, bool, string])
	// to simply an array of strings
	// TODO: Perhaps preserve value but currently there's no use-case
	switch conval := dyval.(type) {
	case float64:
		*aos = ArrayOrString{strconv.FormatFloat(conval, 'f', -1, 64)}
	case bool:
		if conval {
			*aos = ArrayOrString{"true"}
		} else {
			*aos = ArrayOrString{"false"}
		}
	case string:
		*aos = ArrayOrString{conval}
	case []interface{}:
		// A slice of these can...yet again.. be anything (except another array fortunately)
		*aos = ArrayOrString(make([]string, 0, len(conval)))
		for _, dynX := range conval {
			var finalVal string
			switch conXval := dynX.(type) {
			case float64:
				finalVal = strconv.FormatFloat(conXval, 'f', -1, 64)
			case bool:
				if conXval {
					finalVal = "true"
				} else {
					finalVal = "false"
				}
			case string:
				finalVal = conXval
			}
			*aos = append(*aos, finalVal)
		}
	}

	return nil
}

func (aos ArrayOrString) MarshalJSON() ([]byte, error) {
	switch len(aos) {
	case 0:
		return json.Marshal([]byte(""))
	case 1:
		return json.Marshal(string(aos[0]))
	default:
		return json.Marshal([]string(aos))
	}
}

type AutoAlertConfig struct {
	RaiseAlert   bool `json:"raiseAlert"`
	DisableAlert bool `json:"disableAlert"`
}

type BuilderConfig struct {
	ID              string              `json:"id,omitempty"`
	CustomerID      string              `json:"customerId,omitempty"`
	TableName       string              `json:"tableName,omitempty"`
	Added           bool                `json:"added"`
	MatchesFilter   bool                `json:"matchesFilter"`
	Filters         BuilderConfigFilter `json:"filters,omitempty"`
	FiltersJson     string              `json:"filtersjson,omitempty"`
	Severity        string              `json:"severity,omitempty"`
	Key             string              `json:"key,omitempty"`
	ValueField      string              `json:"valueField,omitempty"`
	AutoAlertConfig AutoAlertConfig     `json:"autoAlertConfig"`
}

type AlertRules struct {
	Links  []LinkItem  `json:"links,omitempty"`
	Items  []AlertRule `json:"items,omitempty"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type AlertRule struct {
	ID                     string                 `json:"id,omitempty"`
	CustomerID             string                 `json:"customerId,omitempty"`
	SeedID                 string                 `json:"seedId,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Description            string                 `json:"description,omitempty"`
	Code                   string                 `json:"code,omitempty"`
	Type                   string                 `json:"type,omitempty"`
	Rule                   string                 `json:"rule,omitempty"`
	Grouping               string                 `json:"grouping,omitempty"`
	Enabled                bool                   `json:"enabled,omitempty"`
	Custom                 bool                   `json:"custom,omitempty"`
	Throttled              bool                   `json:"throttled,omitempty"`
	CreatedAt              string                 `json:"createdAt,omitempty"`
	IsInternal             bool                   `json:"isInternal,omitempty"`
	AlertTags              []string               `json:"alertTags,omitempty"`
	CreatedBy              string                 `json:"createdBy,omitempty"`
	UpdatedAt              string                 `json:"updatedAt,omitempty"`
	TimeSuppresionStart    string                 `json:"timeSuppresionStart,omitempty"`
	TimeSuppresionDuration int                    `json:"timeSuppresionDuration,omitempty"`
	UpdatedBy              string                 `json:"updatedBy,omitempty"`
	GroupingL2             string                 `json:"groupingL2,omitempty"`
	GroupingL3             string                 `json:"groupingL3,omitempty"`
	Lock                   bool                   `json:"lock,omitempty"`
	AlertNotifyInterval    int                    `json:"alertNotifyInterval,omitempty"`
	AlertNotifyCount       int                    `json:"alertNotifyCount,omitempty"`
	AlertRuleExceptions    []RuleException        `json:"alertRuleExceptions,omitempty"`
	Destinations           []AlertRuleDestination `json:"destinations,omitempty"`
	SQLConfig              *SQLConfig             `json:"sqlConfig,omitempty"`
	ScriptConfig           *ScriptConfig          `json:"scriptConfig,omitempty"`
	Links                  []LinkItem             `json:"links,omitempty"`
}

type AlertRuleDestination struct {
	ID                 string `json:"id,omitempty"`
	RuleID             string `json:"ruleId,omitempty"`
	Severity           string `json:"severity,omitempty"`
	DestinationID      string `json:"destinationId,omitempty"`
	NotifyEveryAlert   bool   `json:"notifyEveryAlert,omitempty"`
	CloseAfterDelivery bool   `json:"closeAfterDelivery,omitempty"`
	CreatedAt          string `json:"createdAt,omitempty"`
}

type RuleException struct {
	ID          string `json:"id,omitempty"`
	CustomerID  string `json:"customerId,omitempty"`
	RuleID      string `json:"ruleId,omitempty"`
	ExceptionID string `json:"exceptionId,omitempty"`
	CreatedAt   string `json:"createdAt,omitempty"`
	UpdatedAt   string `json:"updatedAt,omitempty"`
}

type LinkItem struct {
	Rel   string `json:"rel,omitempty"`
	Title string `json:"title,omitempty"`
	Href  string `json:"href,omitempty"`
}

type Destinations struct {
	Links  []LinkItem    `json:"links,omitempty"`
	Items  []Destination `json:"items,omitempty"`
	Offset int           `json:"offset,omitempty"`
	Limit  int           `json:"limit,omitempty"`
}

type Destination struct {
	ID         string `json:"id,omitempty"`
	CustomerID string `json:"customerId,omitempty"`
	Name       string `json:"name,omitempty"`
	Type       string `json:"type,omitempty"`
	Address    string `json:"address,omitempty"`
	//Config TODO
	//"config": {
	//  "sender": null
	//},
	CreatedAt string `json:"createdAt,omitempty"`
	CreatedBy string `json:"createdBy,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	UpdatedBy string `json:"updatedBy,omitempty"`
	Enabled   bool   `json:"enabled,omitempty"`
	Default   bool   `json:"default,omitempty"`
	//Template TODO
	Links []LinkItem `json:"links,omitempty"`
}

type SQLConfig struct {
	IntervalSeconds int `json:"intervalSeconds,omitempty"`
}

type AlertRuleQuery struct {
	RuleID    string `json:"ruleId,omitempty"`
	Sequence  int    `json:"sequence,omitempty"`
	QueryID   string `json:"queryId,omitempty"`
	LastRanAt string `json:"lastRanAt,omitempty"`
}

type EventExcludeProfiles struct {
	Links  []LinkItem            `json:"links,omitempty"`
	Items  []EventExcludeProfile `json:"items,omitempty"`
	Offset int                   `json:"offset,omitempty"`
	Limit  int                   `json:"limit,omitempty"`
}

type EventExcludeProfile struct {
	ID           string                      `json:"id,omitempty"`
	CustomerID   string                      `json:"customerId,omitempty"`
	Name         string                      `json:"name,omitempty"`
	Description  string                      `json:"description,omitempty"`
	Priority     int                         `json:"priority,omitempty"`
	Metadata     EventExcludeProfileMetadata `json:"metadata,omitempty"`
	MetadataJson string                      `json:"metadataJson,omitempty"`
	ResourceType string                      `json:"resourceType,omitempty"`
	Platform     string                      `json:"platform,omitempty"`
	CreatedAt    string                      `json:"createdAt,omitempty"`
	CreatedBy    string                      `json:"createdBy,omitempty"`
	UpdatedAt    string                      `json:"updatedAt,omitempty"`
	UpdatedBy    string                      `json:"updatedBy,omitempty"`
	Links        []LinkItem                  `json:"links,omitempty"`
}

type EventExcludeProfileMetadata struct {
	DnsLookupEvents     DnsLookupEvents     `json:"dns_lookup_events,omitempty"`
	UserEvents          UserEvents          `json:"user_events,omitempty"`
	SocketEvents        SocketEvents        `json:"socket_events,omitempty"`
	ProcessEvents       ProcessEvents       `json:"process_events,omitempty"`
	RegistryEvents      RegistryEvents      `json:"registry_events,omitempty"`
	ProcessFileEvents   ProcessFileEvents   `json:"process_file_events,omitempty"`
	HttpEvents          HttpEvents          `json:"http_events,omitempty"`
	EbpfDnsLookupEvents EbpfDnsLookupEvents `json:"ebpf_dns_lookup_events,omitempty"`
}

type EbpfDnsLookupEvents struct {
	Answer   []string `json:"answer,omitempty"`
	Question []string `json:"question,omitempty"`
}

type DnsLookupEvents struct {
	Answer   []string `json:"answer,omitempty"`
	Question []string `json:"question,omitempty"`
}

type HttpEvents struct {
	Host []string `json:"host,omitempty"`
}

type RegistryEvents struct {
	Action []string `json:"action,omitempty"`
}

type UserEvents struct {
	Message []string `json:"message,omitempty"`
}

type SocketEvents struct {
	RemoteAddress []string `json:"remote_address,omitempty"`
}

type ProcessEvents struct {
	Path []string `json:"path,omitempty"`
}

type ProcessFileEvents struct {
	Path       []string `json:"path,omitempty"`
	Operation  []string `json:"operation,omitempty"`
	Executable []string `json:"executable,omitempty"`
}

type iApiType interface {
	AlertRule | Destination | EventExcludeProfile | EventRule | User
	GetID() string
	GetName() string
	KeysToDelete() []string
}

type iApiTypes interface {
	AlertRules | Destinations | EventExcludeProfiles | EventRules | Users
}

type User struct {
	ID                  string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	Email               string   `json:"email,omitempty"`
	Phone               string   `json:"phone,omitempty"`
	Active              bool     `json:"active"`
	SuperAdmin          bool     `json:"superAdmin,omitempty"`
	Bot                 bool     `json:"bot,omitempty"`
	Support             bool     `json:"support,omitempty"`
	PriorLogin          bool     `json:"priorLogin,omitempty"`
	ImageURL            string   `json:"imageUrl,omitempty"`
	Password            string   `json:"password,omitempty"`
	CreatedAt           string   `json:"createdAt,omitempty"`
	MaxIdleTimeMins     int      `json:"maxIdleTimeMins,omitempty"`
	AlertHiddenColumns  []string `json:"alertHiddenColumns,omitempty"`
	UpdatedAt           string   `json:"updatedAt,omitempty"`
	LastUpdatedByUptycs string   `json:"lastUpdatedByUptycs,omitempty"`
	//TODO: Unknown type
	// DetectionHiddenColumns interface{} `json:"detectionHiddenColumns,omitempty"`
	// RangerID               interface{} `json:"rangerId,omitempty"`
	// LastSyncedWithRanger   interface{} `json:"lastSyncedWithRanger,omitempty"`
}

type Users struct {
	Links  []LinkItem `json:"links,omitempty"`
	Items  []User     `json:"items,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}
