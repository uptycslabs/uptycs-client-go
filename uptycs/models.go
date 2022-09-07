package uptycs

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
	ScriptConfig  *ScriptConfig `json:"scriptConfig,omitempty"`
	SQLConfig     *SQLConfig    `json:"sqlConfig,omitempty"`
	BuilderConfig BuilderConfig `json:"builderConfig,omitempty"`
	Links         []LinkItem    `json:"links"`
}

type BuilderConfig struct {
	ID              string                    `json:"id,omitempty"`
	CustomerID      string                    `json:"customerId,omitempty"`
	TableName       string                    `json:"tableName,omitempty"`
	Added           bool                      `json:"added"`
	MatchesFilter   bool                      `json:"matchesFilter"`
	Filters         BuilderConfigFilterString `json:"filters,omitempty"`
	Severity        string                    `json:"severity,omitempty"`
	Key             string                    `json:"key,omitempty"`
	ValueField      string                    `json:"valueField,omitempty"`
	AutoAlertConfig AutoAlertConfig           `json:"autoAlertConfig"`
}

type BuilderConfigFilterString string

func (bcfs *BuilderConfigFilterString) UnmarshalJSON(raw []byte) error {
	if string(raw) == "null" {
		*bcfs = BuilderConfigFilterString("{}")
	} else {
		*bcfs = BuilderConfigFilterString(string(raw))
	}
	return nil
}

func (bcfs BuilderConfigFilterString) MarshalJSON() ([]byte, error) {
	switch len(bcfs) {
	case 0:
		return []byte("{}"), nil
	default:
		return []byte(bcfs), nil
	}
}

type AutoAlertConfig struct {
	RaiseAlert   bool `json:"raiseAlert"`
	DisableAlert bool `json:"disableAlert"`
}

type AlertRules struct {
	Links  []LinkItem  `json:"links,omitempty"`
	Items  []AlertRule `json:"items,omitempty"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type AlertRule struct {
	ID                     string                 `json:"id,omitempty"`
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
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Type    string `json:"type,omitempty"`
	Address string `json:"address,omitempty"`
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
	Name         string                      `json:"name,omitempty"`
	Description  string                      `json:"description,omitempty"`
	Priority     int                         `json:"priority,omitempty"`
	Metadata     EventExcludeProfileMetadata `json:"metadata,omitempty"`
	MetadataJSON string                      `json:"metadataJson,omitempty"`
	ResourceType string                      `json:"resourceType,omitempty"`
	Platform     string                      `json:"platform,omitempty"`
	CreatedAt    string                      `json:"createdAt,omitempty"`
	CreatedBy    string                      `json:"createdBy,omitempty"`
	UpdatedAt    string                      `json:"updatedAt,omitempty"`
	UpdatedBy    string                      `json:"updatedBy,omitempty"`
	Links        []LinkItem                  `json:"links,omitempty"`
}

type EventExcludeProfileMetadata struct {
	DNSLookupEvents     DNSLookupEvents     `json:"dns_lookup_events,omitempty"`
	UserEvents          UserEvents          `json:"user_events,omitempty"`
	SocketEvents        SocketEvents        `json:"socket_events,omitempty"`
	ProcessEvents       ProcessEvents       `json:"process_events,omitempty"`
	RegistryEvents      RegistryEvents      `json:"registry_events,omitempty"`
	ProcessFileEvents   ProcessFileEvents   `json:"process_file_events,omitempty"`
	HTTPEvents          HTTPEvents          `json:"http_events,omitempty"`
	EbpfDNSLookupEvents EbpfDNSLookupEvents `json:"ebpf_dns_lookup_events,omitempty"`
}

type EbpfDNSLookupEvents struct {
	Answer   []string `json:"answer,omitempty"`
	Question []string `json:"question,omitempty"`
}

type DNSLookupEvents struct {
	Answer   []string `json:"answer,omitempty"`
	Question []string `json:"question,omitempty"`
}

type HTTPEvents struct {
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

type User struct {
	ID                  string        `json:"id"`
	Name                string        `json:"name"`
	Email               string        `json:"email" validate:"required,max=512,min=1"`
	Phone               string        `json:"phone"`
	Active              bool          `json:"active"`
	SuperAdmin          bool          `json:"superAdmin"`
	Bot                 bool          `json:"bot"`
	Support             bool          `json:"support"`
	PriorLogin          bool          `json:"priorLogin"`
	ImageURL            string        `json:"imageUrl" validate:"required,max=512,min=1"`
	Password            string        `json:"password"`
	CreatedAt           string        `json:"createdAt"`
	MaxIdleTimeMins     int           `json:"maxIdleTimeMins" validate:"required,max=360,min=1"`
	AlertHiddenColumns  []string      `json:"alertHiddenColumns" validate:"required,min=0"`
	UpdatedAt           string        `json:"updatedAt"`
	LastUpdatedByUptycs string        `json:"lastUpdatedByUptycs"`
	Roles               []Role        `json:"roles" validate:"required,min=0"`
	UserObjectGroups    []ObjectGroup `json:"userObjectGroups"`
	//DetectionHiddenColumns interface{} `json:"detectionHiddenColumns"`
	//RangerID               interface{} `json:"rangerId"`
	//LastSyncedWithRanger   interface{} `json:"lastSyncedWithRanger"`
}

type Users struct {
	Links  []LinkItem `json:"links,omitempty"`
	Items  []User     `json:"items,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type Role struct {
	ID                   string        `json:"id,omitempty"`
	Name                 string        `json:"name"`
	Description          string        `json:"description,omitempty"`
	Permissions          []string      `json:"permissions"`
	Custom               bool          `json:"custom"`
	Hidden               bool          `json:"hidden"`
	CreatedBy            string        `json:"createdBy"`
	UpdatedBy            string        `json:"updatedBy"`
	CreatedAt            string        `json:"createdAt"`
	UpdatedAt            string        `json:"updatedAt"`
	NoMinimalPermissions bool          `json:"noMinimalPermissions"`
	RoleObjectGroups     []ObjectGroup `json:"roleObjectGroups"`
}
type Roles struct {
	Links  []LinkItem `json:"links,omitempty"`
	Items  []Role     `json:"items,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type ComplianceProfile struct {
	ID                   			  string        `json:"id"`
	Name                 			  string        `json:"name,omitempty"`
	Description          			  string        `json:"description,omitempty"`
	SeedId	             			  string        `json:"seedId"`
	CustomerId           			  string        `json:"customerId,omitempty"`
	Custom               			  bool          `json:"custom"`
	Priority             			  int           `json:"priority,omitempty"`
	CreatedBy            			  string        `json:"createdBy"`
	UpdatedBy            			  string        `json:"updatedBy"`
	CreatedAt            			  string        `json:"createdAt"`
	UpdatedAt             			  string        `json:"updatedAt"`
	Links						   	  []LinkItem    `json:"links"`
	ComplianceProfileObjectGroups     []ObjectGroup `json:"complianceProfileObjectGroups"`
}

type ComplianceProfiles struct {
	Links  		[]LinkItem 		     `json:"links,omitempty"`
	Items  		[]ComplianceProfile  `json:"items,omitempty"`
	Offset 		int        			 `json:"offset,omitempty"`
	Limit  		int        			 `json:"limit,omitempty"`
	Decorators  []string             `json:"decorators,omitempty"`
}

type ObjectGroup struct {
	ID               string        `json:"id,omitempty"`
	Name             string        `json:"name,omitempty"`
	Key              string        `json:"key,omitempty"`
	Value            string        `json:"value,omitempty"`
	AssetGroupRuleID string        `json:"assetGroupRuleId,omitempty"`
	ObjectGroupID    string        `json:"objectGroupId,omitempty"`
	UserID           string        `json:"userid,omitempty"`
	RoleID           string        `json:"roleid,omitempty"`
	Description      string        `json:"description,omitempty"`
	Secret           string        `json:"secret,omitempty"`
	ObjectType       string        `json:"objectType,omitempty"`
	Custom           bool          `json:"custom,omitempty"`
	RetentionDays    int           `json:"retentionDays,omitempty"`
	RangerID         int           `json:"rangerId,omitempty"`
	CreatedBy        string        `json:"createdBy,omitempty"`
	UpdatedBy        string        `json:"updatedBy,omitempty"`
	CreatedAt        string        `json:"createdAt,omitempty"`
	UpdatedAt        string        `json:"updatedAt,omitempty"`
	Destinations     []Destination `json:"destinations,omitempty"`
}

type ObjectGroups struct {
	Links  []LinkItem    `json:"links,omitempty"`
	Items  []ObjectGroup `json:"items,omitempty"`
	Offset int           `json:"offset,omitempty"`
	Limit  int           `json:"limit,omitempty"`
}

type iAPIType interface {
	AlertRule | Destination | EventExcludeProfile | EventRule | User | Role | ObjectGroup |ComplianceProfile
	GetID() string
	GetName() string
	KeysToDelete() []string
}

type iAPITypes interface {
	AlertRules | Destinations | EventExcludeProfiles | EventRules | Users | Roles | ObjectGroups | ComplianceProfiles 
}
