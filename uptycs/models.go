package uptycs

import (
	"encoding/json"
	"errors"
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
	ID            string        `json:"id,omitempty"`
	CustomerID    string        `json:"customerId,omitempty"`
	SeedID        string        `json:"seedId,omitempty"`
	Name          string        `json:"name,omitempty"`
	Description   string        `json:"description,omitempty"`
	Code          string        `json:"code,omitempty"`
	Type          string        `json:"type,omitempty"`
	Rule          string        `json:"rule,omitempty"`
	Grouping      string        `json:"grouping,omitempty"`
	Enabled       bool          `json:"enabled,omitempty"`
	Custom        bool          `json:"custom,omitempty"`
	CreatedAt     string        `json:"createdAt,omitempty"`
	IsInternal    bool          `json:"isInternal,omitempty"`
	EventTags     []string      `json:"eventTags,omitempty"`
	CreatedBy     string        `json:"createdBy,omitempty"`
	UpdatedAt     string        `json:"updatedAt,omitempty"`
	UpdatedBy     string        `json:"updatedBy,omitempty"`
	GroupingL2    string        `json:"groupingL2,omitempty"`
	GroupingL3    string        `json:"groupingL3,omitempty"`
	Score         string        `json:"score,omitempty"`
	Lock          bool          `json:"lock,omitempty"`
	ScriptConfig  ScriptConfig  `json:"scriptConfig,omitempty"`
	SQLConfig     SQLConfig     `json:"sqlConfig,omitempty"`
	BuilderConfig BuilderConfig `json:"builderConfig"`
	Links         []LinkItem    `json:"links"`
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
	if raw[0] != '[' {
		raw = []byte("[" + string(raw) + "]")
	}
	return json.Unmarshal(raw, (*[]string)(aos))
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
	RaiseAlert   bool `json:"raiseAlert,omitempty"`
	DisableAlert bool `json:"disableAlert,omitempty"`
}

type BuilderConfig struct {
	ID              string              `json:"id,omitempty"`
	CustomerID      string              `json:"customerId,omitempty"`
	TableName       string              `json:"tableName,omitempty"`
	Added           bool                `json:"added,omitempty"`
	MatchesFilter   bool                `json:"matchesFilter,omitempty"`
	Filters         BuilderConfigFilter `json:"filters,omitempty"`
	Severity        string              `json:"severity,omitempty"`
	Key             string              `json:"key,omitempty"`
	ValueField      string              `json:"valueField,omitempty"`
	AutoAlertConfig AutoAlertConfig     `json:"autoAlertConfig,omitempty"`
}

type AlertRules struct {
	Links  []LinkItem  `json:"links,omitempty"`
	Items  []AlertRule `json:"items,omitempty"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type AlertRule struct {
	ID                     string          `json:"id,omitempty"`
	CustomerID             string          `json:"customerId,omitempty"`
	SeedID                 string          `json:"seedId,omitempty"`
	Name                   string          `json:"name,omitempty"`
	Description            string          `json:"description,omitempty"`
	Code                   string          `json:"code,omitempty"`
	Type                   string          `json:"type,omitempty"`
	Rule                   string          `json:"rule,omitempty"`
	Grouping               string          `json:"grouping,omitempty"`
	Enabled                bool            `json:"enabled,omitempty"`
	Custom                 bool            `json:"custom,omitempty"`
	Throttled              bool            `json:"throttled,omitempty"`
	CreatedAt              string          `json:"createdAt,omitempty"`
	IsInternal             bool            `json:"isInternal,omitempty"`
	AlertTags              []string        `json:"alertTags,omitempty"`
	CreatedBy              string          `json:"createdBy,omitempty"`
	UpdatedAt              string          `json:"updatedAt,omitempty"`
	TimeSuppresionStart    string          `json:"timeSuppresionStart,omitempty"`
	TimeSuppresionDuration int             `json:"timeSuppresionDuration,omitempty"`
	UpdatedBy              string          `json:"updatedBy,omitempty"`
	GroupingL2             string          `json:"groupingL2,omitempty"`
	GroupingL3             string          `json:"groupingL3,omitempty"`
	Lock                   bool            `json:"lock,omitempty"`
	AlertNotifyInterval    int             `json:"alertNotifyInterval,omitempty"`
	AlertNotifyCount       int             `json:"alertNotifyCount,omitempty"`
	AlertRuleExceptions    []RuleException `json:"alertRuleExceptions,omitempty"`
	Destinations           []Destination   `json:"destinations,omitempty"`
	SQLConfig              SQLConfig       `json:"sqlConfig,omitempty"`
	ScriptConfig           ScriptConfig    `json:"scriptConfig,omitempty"`
	Links                  []LinkItem      `json:"links,omitempty"`
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

type Destination struct {
	ID                 string `json:"id,omitempty"`
	CustomerID         string `json:"customerId,omitempty"`
	RuleID             string `json:"ruleId,omitempty"`
	Severity           string `json:"severity,omitempty"`
	DestinationID      string `json:"destinationId,omitempty"`
	NotifyEveryAlert   bool   `json:"notifyEveryAlert,omitempty"`
	CloseAfterDelivery bool   `json:"closeAfterDelivery,omitempty"`
	CreatedAt          string `json:"createdAt,omitempty"`
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
