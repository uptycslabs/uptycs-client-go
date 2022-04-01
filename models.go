package uptycs

type AlertRules struct {
	Links  []LinkItem  `json:"link_item"`
	Items  []AlertRule `json:"items"`
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
}

type AlertRule struct {
	ID                     string        `json:"id"`
	CustomerID             string        `json:"id"`
	SeedID                 string        `json:"id"`
	Name                   string        `json:"name"`
	Description            string        `json:"description"`
	Code                   string        `json:"code"`
	Type                   string        `json:"type"`
	Rule                   string        `json:"rule"`
	Grouping               string        `json:"grouping"`
	Enabled                bool          `json:"enabled"`
	Custom                 bool          `json:"custom"`
	Throttled              bool          `json:"throttled"`
	CreatedAt              string        `json:"created_at"`
	IsInternal             bool          `json:"is_internal"`
	AlertTags              []string      `json:"link_item"`
	CreatedBy              string        `json:"created_by"`
	UpdatedAt              string        `json:"updated_at"`
	TimeSuppresionStart    string        `json:"time_suppresion_start"`
	TimeSuppresionDuration int           `json:"time_suppresion_duration"`
	UpdatedBy              string        `json:"updated_by"`
	GroupingL2             string        `json:"grouping_l2"`
	GroupingL3             string        `json:"grouping_l3"`
	Lock                   bool          `json:"lock"`
	AlertNotifyInterval    int           `json:"alert_notify_interval"`
	AlertNotifyCount       int           `json:"alert_notify_count"`
	Destinations           []Destination `json:"destinations"`
	SqlConfig              string        `json:"sql_config"`
	ScriptConfig           string        `json:"script_config"`
	AlertRuleExceptions    []string      `json:"alert_rule_exceptions"`
	AlertRuleQueries       []string      `json:"alert_rule_queries"`
	Links                  []LinkItem    `json:"link_item"`
}

type LinkItem struct {
	Rel   string `json:"rel"`
	Title string `json:"title"`
	Href  string `json:"href"`
}

type Destination struct {
	ID                 string `json:"id"`
	CustomerID         string `json:"customerId"`
	RuleID             string `json:"ruleId"`
	Severity           string `json:"severity"`
	DestinationID      string `json:"destinationId"`
	NotifyEveryAlert   bool   `json:"notifyEveryAlert"`
	CloseAfterDelivery bool   `json:"closeAfterDelivery"`
	CreatedAt          string `json:"createdAt"`
}
