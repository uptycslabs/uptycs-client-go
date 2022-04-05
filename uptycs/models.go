package uptycs

type AlertRules struct {
	Links  []LinkItem  `json:"links"`
	Items  []AlertRule `json:"items"`
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
}

type AlertRule struct {
	ID                     string        `json:"id"`
	CustomerID             string        `json:"customerId"`
	SeedID                 string        `json:"seedId"`
	Name                   string        `json:"name"`
	Description            string        `json:"description"`
	Code                   string        `json:"code"`
	Type                   string        `json:"type"`
	Rule                   string        `json:"rule"`
	Grouping               string        `json:"grouping"`
	Enabled                bool          `json:"enabled"`
	Custom                 bool          `json:"custom"`
	Throttled              bool          `json:"throttled"`
	CreatedAt              string        `json:"createdAt"`
	IsInternal             bool          `json:"isInternal"`
	AlertTags              []string      `json:"alertTags"`
	CreatedBy              string        `json:"createdBy"`
	UpdatedAt              string        `json:"updatedAt"`
	TimeSuppresionStart    string        `json:"timeSuppresionStart"`
	TimeSuppresionDuration int           `json:"timeSuppresionDuration"`
	UpdatedBy              string        `json:"updatedBy"`
	GroupingL2             string        `json:"groupingL2"`
	GroupingL3             string        `json:"groupingL3"`
	Lock                   bool          `json:"lock"`
	AlertNotifyInterval    int           `json:"alertNotifyInterval"`
	AlertNotifyCount       int           `json:"alertNotifyCount"`
	Destinations           []Destination `json:"destinations"`
	SQLConfig              SQLConfig     `json:"sqlConfig"`
	ScriptConfig           string        `json:"script_config"`
	//AlertRuleExceptions []AlertRuleException       `json:"alertRuleExceptions"`
	//AlertRuleQueries    []AlertRuleQuery `json:"alertRuleQueries"`
	Links []LinkItem `json:"links"`
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

type SQLConfig struct {
	IntervalSeconds int `json:"intervalSeconds"`
}

type AlertRuleQuery struct {
	RuleID    string `json:"ruleId"`
	Sequence  int    `json:"sequence"`
	QueryID   string `json:"queryId"`
	LastRanAt string `json:"lastRanAt"`
}
