package uptycs

type EventRules struct {
	Links  []LinkItem  `json:"links"`
	Items  []EventRule `json:"items"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type EventRule struct {
	ID            string          `json:"id,omitempty"`
	Name          string          `json:"name,omitempty"`
	Description   string          `json:"description,omitempty"`
	Code          string          `json:"code,omitempty"`
	Type          string          `json:"type,omitempty"`
	Rule          string          `json:"rule,omitempty"`
	Grouping      string          `json:"grouping,omitempty"`
	Enabled       bool            `json:"enabled"`
	Custom        bool            `json:"custom"`
	CreatedAt     string          `json:"createdAt,omitempty"`
	IsInternal    bool            `json:"isInternal"`
	EventTags     []string        `json:"eventTags"`
	CreatedBy     string          `json:"createdBy,omitempty"`
	UpdatedAt     string          `json:"updatedAt,omitempty"`
	UpdatedBy     string          `json:"updatedBy,omitempty"`
	GroupingL2    string          `json:"groupingL2,omitempty"`
	GroupingL3    string          `json:"groupingL3,omitempty"`
	Score         string          `json:"score,omitempty"`
	Lock          bool            `json:"lock"`
	Exceptions    []RuleException `json:"exceptions"`
	SQLConfig     *SQLConfig      `json:"sqlConfig,omitempty"`
	BuilderConfig *BuilderConfig  `json:"builderConfig"`
	Links         []LinkItem      `json:"links,omitempty"`
}

type BuilderConfigLite struct {
	ID string `json:"id"`
}

type BuilderConfig struct {
	ID              string           `json:"id,omitempty"`
	TableName       string           `json:"tableName,omitempty"`
	Added           bool             `json:"added"`
	MatchesFilter   bool             `json:"matchesFilter"`
	Filters         CustomJSONString `json:"filters,omitempty"`
	Severity        string           `json:"severity,omitempty"`
	Key             string           `json:"key,omitempty"`
	ValueField      string           `json:"valueField,omitempty"`
	AutoAlertConfig AutoAlertConfig  `json:"autoAlertConfig"`
}

type CustomJSONString string

func (bcfs *CustomJSONString) UnmarshalJSON(raw []byte) error {
	if string(raw) == "null" {
		*bcfs = CustomJSONString("{}")
	} else {
		*bcfs = CustomJSONString(string(raw))
	}
	return nil
}

func (bcfs CustomJSONString) MarshalJSON() ([]byte, error) {
	switch len(bcfs) {
	case 0:
		return []byte("{}"), nil
	default:
		return []byte(bcfs), nil
	}
}

type AutoAlertConfig struct {
	RaiseAlert      bool             `json:"raiseAlert"`
	DisableAlert    bool             `json:"disableAlert"`
	MetadataSources CustomJSONString `json:"metadataSources,omitempty"`
}

type AlertRules struct {
	Links  []LinkItem  `json:"links"`
	Items  []AlertRule `json:"items"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type AlertRule struct {
	ID                     string                 `json:"id,omitempty"`
	Name                   string                 `json:"name,omitempty"`
	Description            string                 `json:"description,omitempty"`
	Code                   string                 `json:"code,omitempty"`
	Type                   string                 `json:"type,omitempty" validate:"required,oneof=sql builder"`
	Rule                   string                 `json:"rule,omitempty"`
	Grouping               string                 `json:"grouping,omitempty"`
	Enabled                bool                   `json:"enabled"`
	Custom                 bool                   `json:"custom"`
	Throttled              bool                   `json:"throttled"`
	CreatedAt              string                 `json:"createdAt,omitempty"`
	IsInternal             bool                   `json:"isInternal"`
	AlertTags              []string               `json:"alertTags" validate:"required_if=Type sql"`
	CreatedBy              string                 `json:"createdBy,omitempty"`
	UpdatedAt              string                 `json:"updatedAt,omitempty"`
	TimeSuppresionStart    string                 `json:"timeSuppresionStart,omitempty"`
	TimeSuppresionDuration int                    `json:"timeSuppresionDuration,omitempty"`
	UpdatedBy              string                 `json:"updatedBy,omitempty"`
	GroupingL2             string                 `json:"groupingL2,omitempty"`
	GroupingL3             string                 `json:"groupingL3,omitempty"`
	Lock                   bool                   `json:"lock"`
	AlertNotifyInterval    int                    `json:"alertNotifyInterval,omitempty"`
	AlertNotifyCount       int                    `json:"alertNotifyCount,omitempty"`
	AlertRuleExceptions    []RuleException        `json:"alertRuleExceptions"`
	Destinations           []AlertRuleDestination `json:"destinations"`
	SQLConfig              *SQLConfig             `json:"sqlConfig,omitempty"`
	Links                  []LinkItem             `json:"links,omitempty"`
	BuilderConfig          *BuilderConfigLite     `json:"builderConfig"`
}

type AlertRuleDestination struct {
	ID                 string `json:"id,omitempty"`
	RuleID             string `json:"ruleId,omitempty"`
	Severity           string `json:"severity,omitempty"`
	DestinationID      string `json:"destinationId,omitempty"`
	NotifyEveryAlert   bool   `json:"notifyEveryAlert"`
	CloseAfterDelivery bool   `json:"closeAfterDelivery"`
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

type Exceptions struct {
	Links  []LinkItem  `json:"links"`
	Items  []Exception `json:"items"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type Exception struct {
	ID              string           `json:"id"`
	Name            string           `json:"name,omitempty"`
	Description     string           `json:"description,omitempty"`
	ExceptionType   string           `json:"exceptionType,omitempty"`
	CreatedBy       string           `json:"createdBy,omitempty"`
	CreatedAt       string           `json:"createdAt,omitempty"`
	UpdatedAt       string           `json:"updatedAt,omitempty"`
	UpdatedBy       string           `json:"updatedBy,omitempty"`
	TableName       string           `json:"tableName,omitempty"`
	IsGlobal        bool             `json:"isGlobal"`
	Custom          bool             `json:"custom"`
	Disabled        bool             `json:"disabled"`
	CloseOpenAlerts bool             `json:"closeOpenAlerts"`
	Rule            CustomJSONString `json:"rule,omitempty"`
	Links           []LinkItem       `json:"links"`
}

type Destinations struct {
	Links  []LinkItem    `json:"links"`
	Items  []Destination `json:"items"`
	Offset int           `json:"offset,omitempty"`
	Limit  int           `json:"limit,omitempty"`
}

type Destination struct {
	ID        string     `json:"id,omitempty"`
	Name      string     `json:"name,omitempty"`
	Type      string     `json:"type,omitempty"`
	Address   string     `json:"address,omitempty"`
	CreatedAt string     `json:"createdAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedAt string     `json:"updatedAt,omitempty"`
	UpdatedBy string     `json:"updatedBy,omitempty"`
	Enabled   bool       `json:"enabled"`
	Default   bool       `json:"default"`
	Links     []LinkItem `json:"links,omitempty"`
	//Template TODO
	//Config TODO
	//"config": {
	//  "sender": null
	//},
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

type AlertRuleCategories struct {
	Links  []LinkItem          `json:"links"`
	Items  []AlertRuleCategory `json:"items"`
	Offset int                 `json:"offset,omitempty"`
	Limit  int                 `json:"limit,omitempty"`
}

type AlertRuleCategory struct {
	ID        string     `json:"id,omitempty"`
	RuleID    string     `json:"ruleId,omitempty"`
	Name      string     `json:"name,omitempty"`
	CreatedAt string     `json:"createdAt,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	Links     []LinkItem `json:"links,omitempty"`
}

type EventExcludeProfiles struct {
	Links  []LinkItem            `json:"links"`
	Items  []EventExcludeProfile `json:"items"`
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
	Answer   []string `json:"answer"`
	Question []string `json:"question"`
}

type DNSLookupEvents struct {
	Answer   []string `json:"answer"`
	Question []string `json:"question"`
}

type HTTPEvents struct {
	Host []string `json:"host"`
}

type RegistryEvents struct {
	Action []string `json:"action"`
}

type UserEvents struct {
	Message []string `json:"message"`
}

type SocketEvents struct {
	RemoteAddress []string `json:"remote_address"`
}

type ProcessEvents struct {
	Path []string `json:"path"`
}

type ProcessFileEvents struct {
	Path       []string `json:"path"`
	Operation  []string `json:"operation"`
	Executable []string `json:"executable"`
}

type User struct {
	ID                  string        `json:"id,omitempty"`
	Name                string        `json:"name"`
	Email               string        `json:"email,omitempty" validate:"required_if=Bot false"`
	Phone               string        `json:"phone,omitempty" validate:"required_if=Bot false"`
	Active              bool          `json:"active"`
	SuperAdmin          bool          `json:"superAdmin"`
	Bot                 bool          `json:"bot" validate:"excluded_with=Email"`
	Support             bool          `json:"support"`
	PriorLogin          bool          `json:"priorLogin"`
	ImageURL            string        `json:"imageUrl,omitempty"`
	Password            string        `json:"password,omitempty"`
	MaxIdleTimeMins     int           `json:"maxIdleTimeMins" validate:"required,max=360,min=1"`
	AlertHiddenColumns  []string      `json:"alertHiddenColumns" validate:"required,min=0"`
	UpdatedAt           string        `json:"updatedAt,omitempty"`
	CreatedAt           string        `json:"createdAt,omitempty"`
	LastUpdatedByUptycs string        `json:"lastUpdatedByUptycs"`
	Roles               []Role        `json:"roles" validate:"required,min=0"`
	UserObjectGroups    []ObjectGroup `json:"userObjectGroups" validate:"required,min=0"`
	//DetectionHiddenColumns interface{} `json:"detectionHiddenColumns"`
	//RangerID               interface{} `json:"rangerId"`
	//LastSyncedWithRanger   interface{} `json:"lastSyncedWithRanger"`
}

type Users struct {
	Links  []LinkItem `json:"links"`
	Items  []User     `json:"items"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type Role struct {
	ID                   string        `json:"id,omitempty"`
	Name                 string        `json:"name,omitempty"`
	Description          string        `json:"description,omitempty"`
	Permissions          []string      `json:"permissions" validate:"required,min=0"`
	Custom               bool          `json:"custom"`
	Hidden               bool          `json:"hidden"`
	CreatedBy            string        `json:"createdBy,omitempty"`
	UpdatedBy            string        `json:"updatedBy,omitempty"`
	CreatedAt            string        `json:"createdAt,omitempty"`
	UpdatedAt            string        `json:"updatedAt,omitempty"`
	NoMinimalPermissions bool          `json:"noMinimalPermissions"`
	RoleObjectGroups     []ObjectGroup `json:"roleObjectGroups" validate:"required,min=0"`
}
type Roles struct {
	Links  []LinkItem `json:"links"`
	Items  []Role     `json:"items"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type ComplianceProfile struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Custom      bool       `json:"custom"`
	Priority    int        `json:"priority,omitempty"`
	CreatedBy   string     `json:"createdBy,omitempty"`
	UpdatedBy   string     `json:"updatedBy,omitempty"`
	CreatedAt   string     `json:"createdAt,omitempty"`
	UpdatedAt   string     `json:"updatedAt,omitempty"`
	Links       []LinkItem `json:"links,omitempty"`
}

type ComplianceProfiles struct {
	Links      []LinkItem          `json:"links"`
	Items      []ComplianceProfile `json:"items"`
	Offset     int                 `json:"offset,omitempty"`
	Limit      int                 `json:"limit,omitempty"`
	Decorators []string            `json:"decorators"`
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
	Custom           bool          `json:"custom"`
	RetentionDays    int           `json:"retentionDays,omitempty"`
	RangerID         int           `json:"rangerId,omitempty"`
	CreatedBy        string        `json:"createdBy,omitempty"`
	UpdatedBy        string        `json:"updatedBy,omitempty"`
	CreatedAt        string        `json:"createdAt,omitempty"`
	UpdatedAt        string        `json:"updatedAt,omitempty"`
	Destinations     []Destination `json:"destinations"`
}

type ObjectGroups struct {
	Links  []LinkItem    `json:"links"`
	Items  []ObjectGroup `json:"items"`
	Offset int           `json:"offset,omitempty"`
	Limit  int           `json:"limit,omitempty"`
}

type TagConfigurations struct {
	Links  []LinkItem         `json:"links"`
	Items  []TagConfiguration `json:"items"`
	Offset int                `json:"offset,omitempty"`
	Limit  int                `json:"limit,omitempty"`
}
type TagConfiguration Tag

type TagRules struct {
	Links  []LinkItem `json:"links"`
	Items  []TagRule  `json:"items"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type TagRule struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty" validate:"required,max=255,min=1"`
	Description    string `json:"description,omitempty"`
	Query          string `json:"query,omitempty" validate:"required"`
	Source         string `json:"source,omitempty" validate:"required,oneof=global realtime"`
	RunOnce        bool   `json:"runOnce"`
	Interval       int    `json:"interval,omitempty" validate:"required_if=RunOnce false"`
	OSqueryVersion string `json:"osqueryVersion,omitempty"`
	Platform       string `json:"platform,omitempty" validate:"required_if=Source realtime"`
	Enabled        bool   `json:"enabled"`
	System         bool   `json:"system"`
	LastRunAt      string `json:"lastRunAt,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	UpdatedBy      string `json:"updatedBy,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	UpdatedAt      string `json:"updatedAt,omitempty"`
	ResourceType   string `json:"resourceType,omitempty"`
}

type Tags struct {
	Links  []LinkItem `json:"links"`
	Items  []Tag      `json:"items"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type Tag struct {
	ID                          string                   `json:"id,omitempty"`
	Name                        string                   `json:"-"` // There is no name
	Value                       string                   `json:"value,omitempty"`
	Key                         string                   `json:"key"`
	CreatedBy                   string                   `json:"createdBy,omitempty"`
	UpdatedBy                   string                   `json:"updatedBy,omitempty"`
	FlagProfileID               string                   `json:"flagProfileId,omitempty"`
	CustomProfileID             string                   `json:"customProfileId,omitempty"`
	ComplianceProfileID         string                   `json:"complianceProfileId,omitempty"`
	ProcessBlockRuleID          string                   `json:"processBlockRuleId,omitempty"`
	DNSBlockRuleID              string                   `json:"dnsBlockRuleId,omitempty"`
	WindowsDefenderPreferenceID string                   `json:"windowsDefenderPreferenceId,omitempty"`
	TagRuleID                   string                   `json:"tagRuleId,omitempty"`
	Tag                         string                   `json:"tag,omitempty"`
	Custom                      bool                     `json:"custom"`
	System                      bool                     `json:"system"`
	CreatedAt                   string                   `json:"createdAt,omitempty"`
	ExpireAt                    string                   `json:"expireAt,omitempty"`
	Status                      string                   `json:"status,omitempty"`
	Source                      string                   `json:"source,omitempty"`
	UpdatedAt                   string                   `json:"updatedAt,omitempty"`
	ResourceType                string                   `json:"resourceType,omitempty"`
	FilePathGroups              []TagConfigurationObject `json:"filePathGroups" validate:"required,min=0"`
	EventExcludeProfiles        []TagConfigurationObject `json:"eventExcludeProfiles" validate:"required,min=0"`
	RegistryPaths               []TagConfigurationObject `json:"registryPaths" validate:"required,min=0"`
	Querypacks                  []TagConfigurationObject `json:"querypacks" validate:"required,min=0"`
	YaraGroupRules              []TagConfigurationObject `json:"yaraGroupRules" validate:"required,min=0"`
	AuditConfigurations         []TagConfigurationObject `json:"auditConfigurations" validate:"required,min=0"`
	//ImageLoadExclusions []interface{} `json:"imageLoadExclusions"` # TODO: cant find any examples of this
	//AuditGroups         []interface{} `json:"auditGroups"` # TODO: cant find any examples of this
	//Destinations        []interface{} `json:"destinations"` # TODO: cant find any examples of this
	//Redactions          []interface{} `json:"redactions"` # TODO: cant find any examples of this
	//AuditRules          []interface{} `json:"auditRules"` # TODO: cant find any examples of this
	//PrometheusTargets   []interface{} `json:"prometheusTargets"` # TODO: cant find any examples of this
	//AtcQueries          []interface{} `json:"atcQueries"` # TODO: cant find any examples of this
}

type TagConfigurationObjectDetails struct {
	ID                   string `json:"id,omitempty"`
	AuditConfigurationID string `json:"auditConfigurationId,omitempty"`
	YaraGroupRuleID      string `json:"yaraGroupRuleId,omitempty"`
	QuerypackID          string `json:"querypackId,omitempty"`
	RegistryPathID       string `json:"registryPathId,omitempty"`
	EventExcludeProfile  string `json:"eventExcludeProfile,omitempty"`
	FilePathGroupID      string `json:"filePathGroupId,omitempty"`
	TagID                string `json:"tagId,omitempty"`
	CreatedBy            string `json:"createdBy,omitempty"`
	CreatedAt            string `json:"createdAt,omitempty"`
}

type TagConfigurationObject struct {
	ID                     string                         `json:"id,omitempty"`
	Name                   string                         `json:"name,omitempty"`
	AuditConfigurationTag  *TagConfigurationObjectDetails `json:"AuditConfigurationTag,omitempty"`
	YaraGroupRuleTag       *TagConfigurationObjectDetails `json:"YaraGroupRuleTag,omitempty"`
	QuerypackTag           *TagConfigurationObjectDetails `json:"QuerypackTag,omitempty"`
	RegistryPathTag        *TagConfigurationObjectDetails `json:"RegistryPathTag,omitempty"`
	EventExcludeProfileTag *TagConfigurationObjectDetails `json:"EventExcludeProfileTag,omitempty"`
	FilePathGroupTag       *TagConfigurationObjectDetails `json:"FilePathGroupTag,omitempty"`
	Links                  []LinkItem                     `json:"links,omitempty"`
}
type FilePathGroups struct {
	Links  []LinkItem      `json:"links"`
	Items  []FilePathGroup `json:"items"`
	Offset int             `json:"offset,omitempty"`
	Limit  int             `json:"limit,omitempty"`
}

type FilePathGroup struct {
	ID                    string                   `json:"id,omitempty"`
	Name                  string                   `json:"name,omitempty"`
	Description           string                   `json:"description,omitempty"`
	Grouping              string                   `json:"grouping,omitempty"`
	IncludePaths          []string                 `json:"includePaths"`
	IncludePathExtensions []string                 `json:"includePathExtensions"`
	ExcludePaths          []string                 `json:"excludePaths"`
	Custom                bool                     `json:"custom" validate:"required"`
	CheckSignature        bool                     `json:"checkSignature"`
	FileAccesses          bool                     `json:"fileAccesses"`
	ExcludeProcessNames   []string                 `json:"excludeProcessNames"`
	PriorityPaths         []string                 `json:"priorityPaths"`
	CreatedBy             string                   `json:"createdBy,omitempty"`
	UpdatedBy             string                   `json:"updatedBy,omitempty"`
	CreatedAt             string                   `json:"createdAt,omitempty"`
	UpdatedAt             string                   `json:"updatedAt,omitempty"`
	Signatures            []FilePathGroupSignature `json:"signatures" validate:"required,min=0"`
	YaraGroupRules        []YaraGroupRule          `json:"yaraGroupRules" validate:"required,min=0"`
	Links                 []LinkItem               `json:"links,omitempty"`
	//ExcludeProcessPaths   []string                 `json:"excludeProcessPaths"` //TODO this seems broken in the API. returns null or {}
}

type FilePathGroupSignature struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Paths       []string `json:"paths"`
	CreatedBy   string   `json:"createdBy,omitempty"`
	UpdatedBy   string   `json:"updatedBy,omitempty"`
	CreatedAt   string   `json:"createdAt,omitempty"`
	UpdatedAt   string   `json:"updatedAt,omitempty"`
}

type YaraGroupRules struct {
	Links  []LinkItem      `json:"links"`
	Items  []YaraGroupRule `json:"items"`
	Offset int             `json:"offset,omitempty"`
	Limit  int             `json:"limit,omitempty"`
}

type YaraGroupRule struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Rules       string     `json:"rules,omitempty"`
	Custom      bool       `json:"custom"`
	CreatedBy   string     `json:"createdBy,omitempty"`
	UpdatedBy   string     `json:"updatedBy,omitempty"`
	CreatedAt   string     `json:"createdAt,omitempty"`
	UpdatedAt   string     `json:"updatedAt,omitempty"`
	Links       []LinkItem `json:"links,omitempty"`
}

type RegistryPaths struct {
	Links  []LinkItem     `json:"links"`
	Items  []RegistryPath `json:"items"`
	Offset int            `json:"offset,omitempty"`
	Limit  int            `json:"limit,omitempty"`
}

type RegistryPath struct {
	ID                   string     `json:"id,omitempty"`
	Name                 string     `json:"name,omitempty"`
	Description          string     `json:"description,omitempty"`
	Grouping             string     `json:"grouping,omitempty"`
	IncludeRegistryPaths []string   `json:"includeRegistryPaths" validate:"required,min=1"`
	RegAccesses          bool       `json:"regAccesses"`
	ExcludeRegistryPaths []string   `json:"excludeRegistryPaths"`
	Custom               bool       `json:"custom"`
	CreatedBy            string     `json:"createdBy,omitempty"`
	UpdatedBy            string     `json:"updatedBy,omitempty"`
	CreatedAt            string     `json:"createdAt,omitempty"`
	UpdatedAt            string     `json:"updatedAt,omitempty"`
	Links                []LinkItem `json:"links,omitempty"`
}

type Querypacks struct {
	Links  []LinkItem  `json:"links"`
	Items  []Querypack `json:"items"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type Querypack struct {
	ID               string           `json:"id,omitempty"`
	Name             string           `json:"name,omitempty"`
	Description      string           `json:"description" validate:"required,max=1024,min=1"`
	Type             string           `json:"type" validate:"required,oneof=compliance default hardware incident system vulnerability"`
	AdditionalLogger bool             `json:"additionalLogger"`
	Custom           bool             `json:"custom"`
	CreatedBy        string           `json:"createdBy,omitempty"`
	UpdatedBy        string           `json:"updatedBy,omitempty"`
	CreatedAt        string           `json:"createdAt,omitempty"`
	UpdatedAt        string           `json:"updatedAt,omitempty"`
	IsInternal       bool             `json:"isInternal"`
	ResourceType     string           `json:"resourceType"`
	Queries          []Query          `json:"queries"`
	Conf             CustomJSONString `json:"conf,omitempty"`
	Links            []LinkItem       `json:"links,omitempty"`
	//Sha              string           `json:"sha" validate:"required,max=40,min=1"` //TODO Does not work
}

type Query struct {
	ID          string           `json:"id,omitempty"`
	Name        string           `json:"name,omitempty"`
	Description string           `json:"description,omitempty"`
	Query       string           `json:"query,omitempty"`
	Removed     bool             `json:"removed"`
	Version     string           `json:"version,omitempty"`
	Interval    int              `json:"interval"`
	Platform    string           `json:"platform,omitempty"`
	Snapshot    bool             `json:"snapshot"`
	RunNow      bool             `json:"runNow"`
	Value       string           `json:"value,omitempty"`
	QuerypackID string           `json:"querypackId,omitempty"`
	TableName   string           `json:"tableName,omitempty"`
	DataTypes   CustomJSONString `json:"dataTypes,omitempty"` //This is super ephemeral
	Verified    bool             `json:"verified"`
	CreatedBy   string           `json:"createdBy,omitempty"`
	UpdatedBy   string           `json:"updatedBy,omitempty"`
	CreatedAt   string           `json:"createdAt,omitempty"`
	UpdatedAt   string           `json:"updatedAt,omitempty"`
}

type AuditConfigurations struct {
	Links  []LinkItem           `json:"links"`
	Items  []AuditConfiguration `json:"items"`
	Offset int                  `json:"offset,omitempty"`
	Limit  int                  `json:"limit,omitempty"`
}

type AuditConfiguration struct {
	ID          string       `json:"id,omitempty"`
	Name        string       `json:"name,omitempty"`
	Description string       `json:"description,omitempty"`
	Framework   string       `json:"framework" validate:"required,oneof=CIS PCI CUSTOM FEDRAMP SOC2 HIPAA STIG NIST ISO VDA-TISAX"`
	Version     string       `json:"version" validate:"required,max=256,min=1"`
	OsVersion   string       `json:"osVersion" validate:"required,min=1"`
	Platform    string       `json:"platform" validate:"required,min=1"`
	TableName   string       `json:"tableName" validate:"required,min=1"`
	Sha256      string       `json:"sha256,omitempty"`
	CreatedBy   string       `json:"createdBy,omitempty"`
	UpdatedBy   string       `json:"updatedBy,omitempty"`
	CreatedAt   string       `json:"createdAt,omitempty"`
	UpdatedAt   string       `json:"updatedAt,omitempty"`
	Type        string       `json:"type,omitempty"`
	Checks      int          `json:"checks,omitempty"`
	AuditEntry  []AuditEntry `json:"auditEntities"`
	Links       []LinkItem   `json:"links,omitempty"`
}

type AuditEntry struct {
	ID                   string   `json:"id,omitempty"`
	AuditConfigurationID string   `json:"auditConfigurationId,omitempty"`
	AuditName            []string `json:"auditName"`
	Standard             string   `json:"standard,omitempty"`
	Version              string   `json:"version,omitempty"`
	Section              string   `json:"section,omitempty"`
	Title                string   `json:"title,omitempty"`
	Scored               bool     `json:"scored"`
	Level                string   `json:"level,omitempty"`
	Description          string   `json:"description,omitempty"`
	Rationale            string   `json:"rationale,omitempty"`
	Command              string   `json:"command,omitempty"`
	Remediation          string   `json:"remediation,omitempty"`
	ExpectedValue        string   `json:"expectedValue,omitempty"`
	AuthoritativeSource  string   `json:"authoritativeSource,omitempty"`
	Exception            string   `json:"exception,omitempty"`
	Chapter              string   `json:"chapter,omitempty"`
	CheckID              string   `json:"checkId,omitempty"`
	Enabled              bool     `json:"enabled"`
	Service              string   `json:"service,omitempty"`
	CreatedBy            string   `json:"createdBy,omitempty"`
	Score                float64  `json:"score"`
	UpdatedBy            string   `json:"updatedBy,omitempty"`
	RunCategory          int      `json:"runCategory"`
	Timeout              int      `json:"timeout"`
	CreatedAt            string   `json:"createdAt,omitempty"`
	UpdatedAt            string   `json:"updatedAt,omitempty"`
	IsManual             bool     `json:"isManual"`
	//RemediationAction    interface{} `json:"remediationAction"`// TODO there are no examples of this
	//Parameters           interface{} `json:"parameters"` // TODO there are no examples of this
}

type AssetGroupRules struct {
	Links  []LinkItem       `json:"links"`
	Items  []AssetGroupRule `json:"items"`
	Offset int              `json:"offset,omitempty"`
	Limit  int              `json:"limit,omitempty"`
}

type AssetGroupRule struct {
	ID             string     `json:"id,omitempty"`
	Name           string     `json:"name"`
	Description    string     `json:"description,omitempty"`
	Query          string     `json:"query"`
	Interval       int        `json:"interval,omitempty"`
	OsqueryVersion string     `json:"osqueryVersion,omitempty"`
	Platform       string     `json:"platform,omitempty"`
	Enabled        bool       `json:"enabled"`
	CreatedBy      string     `json:"createdBy,omitempty"`
	UpdatedBy      string     `json:"updatedBy,omitempty"`
	CreatedAt      string     `json:"createdAt,omitempty"`
	UpdatedAt      string     `json:"updatedAt,omitempty"`
	Links          []LinkItem `json:"links,omitempty"`
}

type PathStruct struct {
	Path string `json:"path,omitempty"`
}

type AtcQueries struct {
	Links  []LinkItem `json:"links"`
	Items  []AtcQuery `json:"items"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type AtcQuery struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Query       string `json:"query,omitempty"`
	OsPaths     struct {
		Darwin  []PathStruct `json:"darwin,omitempty"`
		Debian  []PathStruct `json:"debian,omitempty"`
		Windows []PathStruct `json:"windows,omitempty"`
	} `json:"osPaths,omitempty"`
	Columns []struct {
		Name        string `json:"name,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"columns,omitempty"`
	CreatedBy string     `json:"createdBy,omitempty"`
	UpdatedBy string     `json:"updatedBy,omitempty"`
	CreatedAt string     `json:"createdAt,omitempty"`
	UpdatedAt string     `json:"updatedAt,omitempty"`
	Links     []LinkItem `json:"links,omitempty"`
}

type Carves struct {
	Links  []LinkItem `json:"links"`
	Items  []Carve    `json:"items"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type Carve struct {
	ID              string `json:"id,omitempty"`
	Name            string `json:"-"` // Required but not actually in a carve
	AssetID         string `json:"assetId,omitempty"`
	Path            string `json:"path,omitempty"`
	CreatedAt       string `json:"createdAt,omitempty"`
	UpdatedAt       string `json:"updatedAt,omitempty"`
	Status          string `json:"status,omitempty"`
	DeletedUserName string `json:"deletedUserName,omitempty"`
	DeletedAt       string `json:"deletedAt,omitempty"`
	AssetHostName   string `json:"assetHostName,omitempty"`
	Offset          int    `json:"offset,omitempty"`
	Length          int    `json:"length,omitempty"`
	//Error           interface{} `json:"error"` TODO
}
type CustomProfiles struct {
	Links  []LinkItem      `json:"links"`
	Items  []CustomProfile `json:"items"`
	Offset int             `json:"offset,omitempty"`
	Limit  int             `json:"limit,omitempty"`
}

type CustomProfile struct {
	ID             string           `json:"id,omitempty"`
	Name           string           `json:"name,omitempty"`
	Description    string           `json:"description,omitempty"`
	QuerySchedules CustomJSONString `json:"querySchedules,omitempty" validate:"required,min=1"`
	Priority       int              `json:"priority"`
	CreatedBy      string           `json:"createdBy,omitempty"`
	UpdatedBy      string           `json:"updatedBy,omitempty"`
	CreatedAt      string           `json:"createdAt,omitempty"`
	UpdatedAt      string           `json:"updatedAt,omitempty"`
	ResourceType   string           `json:"resourceType,omitempty"`
	Links          []LinkItem       `json:"links,omitempty"`
}

type FlagProfiles struct {
	Links  []LinkItem    `json:"links"`
	Items  []FlagProfile `json:"items"`
	Offset int           `json:"offset,omitempty"`
	Limit  int           `json:"limit,omitempty"`
}

type FlagProfile struct {
	ID           string           `json:"id,omitempty"`
	Custom       bool             `json:"custom"`
	Name         string           `json:"name,omitempty"`
	Description  string           `json:"description,omitempty"`
	Priority     int              `json:"priority"`
	Flags        CustomJSONString `json:"flags,omitempty" validate:"required,min=1"`
	OsFlags      CustomJSONString `json:"osFlags,omitempty" validate:"required,min=1"`
	CreatedBy    string           `json:"createdBy,omitempty"`
	UpdatedBy    string           `json:"updatedBy,omitempty"`
	CreatedAt    string           `json:"createdAt,omitempty"`
	UpdatedAt    string           `json:"updatedAt,omitempty"`
	ResourceType string           `json:"resourceType,omitempty"`
	Links        []LinkItem       `json:"links,omitempty"`
}

type BlockRules struct {
	Links  []LinkItem  `json:"links"`
	Items  []BlockRule `json:"items"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type BlockRule struct {
	ID                        string     `json:"id,omitempty"`
	Name                      string     `json:"name,omitempty"`
	Description               string     `json:"description,omitempty"`
	Platform                  string     `json:"platform,omitempty"`
	Priority                  int        `json:"priority"`
	Type                      string     `json:"type,omitempty"`
	Status                    string     `json:"status,omitempty"`
	Custom                    bool       `json:"custom"`
	EnableLockdown            bool       `json:"enableLockdown"`
	CertificateMode           string     `json:"certificateMode,omitempty"`
	HasLookupTable            bool       `json:"hasLookupTable"`
	CreatedAt                 string     `json:"createdAt,omitempty"`
	CreatedBy                 string     `json:"createdBy,omitempty"`
	UpdatedAt                 string     `json:"updatedAt,omitempty"`
	UpdatedBy                 string     `json:"updatedBy,omitempty"`
	Checks                    int        `json:"checks"`
	AssetsCount               int        `json:"assetsCount"`
	UptycsProtectEnabledCount int        `json:"uptycsProtectEnabledCount"`
	Links                     []LinkItem `json:"links,omitempty"`
	//Tags                      []interface{} `json:"tags"` // TODO cant find any
	//AssociatedAssets          []interface{} `json:"associatedAssets"` // TODO cant find any
}

type WindowsDefenderPreferences struct {
	Links  []LinkItem                  `json:"links"`
	Items  []WindowsDefenderPreference `json:"items"`
	Offset int                         `json:"offset,omitempty"`
	Limit  int                         `json:"limit,omitempty"`
}

type WindowsDefenderPreference struct {
	ID                                            string     `json:"id,omitempty"`
	Name                                          string     `json:"name,omitempty"`
	Description                                   string     `json:"description,omitempty"`
	Priority                                      int        `json:"priority"`
	RealTimeScanDirection                         int        `json:"realTimeScanDirection"`
	CheckForSignaturesBeforeRunningScan           bool       `json:"checkForSignaturesBeforeRunningScan"`
	ScanOnlyIfIdleEnabled                         bool       `json:"scanOnlyIfIdleEnabled"`
	ScanScheduleDay                               int        `json:"scanScheduleDay"`
	SignatureDisableUpdateOnStartupWithoutEngine  bool       `json:"signatureDisableUpdateOnStartupWithoutEngine"`
	SignatureScheduleDay                          int        `json:"signatureScheduleDay"`
	DisablePrivacyMode                            bool       `json:"disablePrivacyMode"`
	RandomizeScheduleTaskTimes                    bool       `json:"randomizeScheduleTaskTimes"`
	DisableBehaviorMonitoring                     bool       `json:"disableBehaviorMonitoring"`
	DisableIntrusionPreventionSystem              bool       `json:"disableIntrusionPreventionSystem"`
	DisableIOAVProtection                         bool       `json:"disableIOAVProtection"`
	DisableRealtimeMonitoring                     bool       `json:"disableRealtimeMonitoring"`
	DisableScriptScanning                         bool       `json:"disableScriptScanning"`
	DisableArchiveScanning                        bool       `json:"disableArchiveScanning"`
	DisableCatchupFullScan                        bool       `json:"disableCatchupFullScan"`
	DisableCatchupQuickScan                       bool       `json:"disableCatchupQuickScan"`
	DisableEmailScanning                          bool       `json:"disableEmailScanning"`
	DisableRemovableDriveScanning                 bool       `json:"disableRemovableDriveScanning"`
	DisableRestorePoint                           bool       `json:"disableRestorePoint"`
	DisableScanningMappedNetworkDrivesForFullScan bool       `json:"disableScanningMappedNetworkDrivesForFullScan"`
	DisableScanningNetworkFiles                   bool       `json:"disableScanningNetworkFiles"`
	UILockdown                                    bool       `json:"uILockdown"`
	Force                                         bool       `json:"force"`
	CreatedBy                                     string     `json:"createdBy,omitempty"`
	UpdatedBy                                     string     `json:"updatedBy,omitempty"`
	CreatedAt                                     string     `json:"createdAt,omitempty"`
	UpdatedAt                                     string     `json:"updatedAt,omitempty"`
	Links                                         []LinkItem `json:"links"`
	//ThreatIDDefaultActionIds                      interface{} `json:"threatIDDefaultActionIds"` // TODO cant find any
	//ThreatIDDefaultActionActions                  interface{} `json:"threatIDDefaultActionActions"` // TODO cant find any
	//UnknownThreatDefaultAction                    interface{} `json:"unknownThreatDefaultAction"` // TODO cant find any
	//LowThreatDefaultAction                        interface{} `json:"lowThreatDefaultAction"` // TODO cant find any
	//ModerateThreatDefaultAction                   interface{} `json:"moderateThreatDefaultAction"` // TODO cant find any
	//HighThreatDefaultAction                       interface{} `json:"highThreatDefaultAction"` // TODO cant find any
	//SevereThreatDefaultAction                     interface{} `json:"severeThreatDefaultAction"` // TODO cant find any
	//SignatureScheduleTime                         interface{} `json:"signatureScheduleTime"` // TODO cant find any
	//SignatureUpdateCatchupInterval                interface{} `json:"signatureUpdateCatchupInterval"` // TODO cant find any
	//SignatureUpdateInterval                       interface{} `json:"signatureUpdateInterval"` // TODO cant find any
	//MAPSReporting                                 interface{} `json:"mAPSReporting"` // TODO cant find any
	//SignatureFallbackOrder                        interface{} `json:"signatureFallbackOrder"` // TODO cant find any
	//ScanParameters                                interface{} `json:"scanParameters"` // TODO cant find any
	//ScanPurgeItemsAfterDelay                      interface{} `json:"scanPurgeItemsAfterDelay"` // TODO cant find any
	//RemediationScheduleDay                        interface{} `json:"remediationScheduleDay"` // TODO cant find any
	//RemediationScheduleTime                       interface{} `json:"remediationScheduleTime"` // TODO cant find any
	//ReportingAdditionalActionTimeOut              interface{} `json:"reportingAdditionalActionTimeOut"` // TODO cant find any
	//ReportingCriticalFailureTimeOut               interface{} `json:"reportingCriticalFailureTimeOut"` // TODO cant find any
	//ReportingNonCriticalTimeOut                   interface{} `json:"reportingNonCriticalTimeOut"` // TODO cant find any
	//ScanAvgCPULoadFactor                          interface{} `json:"scanAvgCPULoadFactor"` // TODO cant find any
	//ExclusionPath                                 interface{} `json:"exclusionPath"` // TODO cant find any
	//ExclusionExtension                            interface{} `json:"exclusionExtension"` // TODO cant find any
	//ExclusionProcess                              interface{} `json:"exclusionProcess"` // TODO cant find any
	//QuarantinePurgeItemsAfterDelay                interface{} `json:"quarantinePurgeItemsAfterDelay"` // TODO cant find any
	//ScanScheduleQuickScanTime                     interface{} `json:"scanScheduleQuickScanTime"` // TODO cant find any
	//ScanScheduleTime                              interface{} `json:"scanScheduleTime"` // TODO cant find any
	//SignatureFirstAuGracePeriod                   interface{} `json:"signatureFirstAuGracePeriod"` // TODO cant find any
	//SignatureAuGracePeriod                        interface{} `json:"signatureAuGracePeriod"` // TODO cant find any
	//SignatureDefinitionUpdateFileSharesSources    interface{} `json:"signatureDefinitionUpdateFileSharesSources"` // TODO cant find any
}

type Assets struct {
	Links  []LinkItem `json:"links"`
	Items  []Asset    `json:"items"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type City struct {
	ID                 string `json:"id"`
	Name               string `json:"name,omitempty"`
	SubdivisionIsoCode string `json:"subdivisionIsoCode,omitempty"`
	SubdivisionName    string `json:"subdivisionName,omitempty"`
	CountryIsoCode     string `json:"countryIsoCode,omitempty"`
	CountryName        string `json:"countryName,omitempty"`
}

type AssetCloudInfo struct {
	CloudID  string `json:"cloudId,omitempty"`
	Provider string `json:"provider,omitempty"`
	Account  string `json:"account,omitempty"`
	Image    string `json:"image,omitempty"`
	Type     string `json:"type,omitempty"`
	Region   string `json:"region,omitempty"`
	Zone     string `json:"zone,omitempty"`
	Service  string `json:"service,omitempty"`
}

type AssetCapabilities struct {
	ID               string `json:"id"`
	AssetID          string `json:"assetId,omitempty"`
	Name             string `json:"name,omitempty"`
	Status           string `json:"status,omitempty"`
	IndividualStatus struct {
		ConfigurationStatus  string `json:"configurationStatus,omitempty"`
		FlagStatus           string `json:"flagStatus,omitempty"`
		ScheduledQueryStatus string `json:"scheduledQueryStatus,omitempty"`
	} `json:"individualStatus,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type AssetInterface struct {
	Name      string `json:"name"`
	Mac       string `json:"mac"`
	IP        string `json:"ip"`
	Mask      string `json:"mask"`
	IsPrimary bool   `json:"isPrimary"`
}

type Asset struct {
	Gateway                   string                 `json:"gateway,omitempty"`
	Name                      string                 `json:"hostname,omitempty"` // there is no name, use hostname for ease
	CityID                    string                 `json:"cityId,omitempty"`
	CreatedAt                 string                 `json:"createdAt,omitempty"`
	Disabled                  bool                   `json:"disabled"`
	HostName                  string                 `json:"hostName,omitempty"`
	ID                        string                 `json:"id,omitempty"`
	LastEnrolledAt            string                 `json:"lastEnrolledAt,omitempty"`
	Os                        string                 `json:"os,omitempty"`
	OsFlavor                  string                 `json:"osFlavor,omitempty"`
	OsqueryVersion            string                 `json:"osqueryVersion,omitempty"`
	OsVersion                 string                 `json:"osVersion,omitempty"`
	Status                    string                 `json:"status,omitempty"`
	UpgradeState              bool                   `json:"upgradeState"`
	ObjectGroupID             string                 `json:"objectGroupId,omitempty"`
	Live                      bool                   `json:"live"`
	Location                  string                 `json:"location,omitempty"`
	ManualSlackAssignment     bool                   `json:"manualSlackAssignment"`
	HardwareVendor            string                 `json:"hardwareVendor,omitempty"`
	AssetObjectGroupID        string                 `json:"assetObjectGroupId,omitempty"`
	PackageObjectGroupID      string                 `json:"packageObjectGroupId,omitempty"`
	ObjectGroup               ObjectGroup            `json:"objectGroup"`
	City                      City                   `json:"city"`
	AssetCloudInfo            AssetCloudInfo         `json:"assetCloudInfo"`
	LastActivityAt            string                 `json:"lastActivityAt,omitempty"`
	LastUpgradedAt            string                 `json:"lastUpgradedAt,omitempty"`
	EventExcludeProfiles      []EventExcludeProfiles `json:"eventExcludeProfiles"`
	Tags                      []string               `json:"tags"`
	Capabilities              []AssetCapabilities    `json:"capabilities"`
	OsDisplay                 string                 `json:"osDisplay,omitempty"`
	Description               string                 `json:"description,omitempty"`
	Latitude                  float64                `json:"latitude"`
	Longitude                 float64                `json:"longitude"`
	CPUBrand                  string                 `json:"cpuBrand,omitempty"`
	HardwareModel             string                 `json:"hardwareModel,omitempty"`
	HardwareSerial            string                 `json:"hardwareSerial,omitempty"`
	Cores                     int                    `json:"cores"`
	LogicalCores              int                    `json:"logicalCores"`
	MemoryMb                  int                    `json:"memoryMb"`
	OsKey                     string                 `json:"osKey,omitempty"`
	OsVersionSortable         string                 `json:"osVersionSortable,omitempty"`
	ProfileID                 string                 `json:"profileId,omitempty"`
	OsqueryVersionSortable    string                 `json:"osqueryVersionSortable,omitempty"`
	UpgradeOsqueryVersion     string                 `json:"upgradeOsqueryVersion,omitempty"`
	NewEpoch                  bool                   `json:"newEpoch"`
	LastEpoch                 string                 `json:"lastEpoch,omitempty"`
	ComplianceProfileID       string                 `json:"complianceProfileId,omitempty"`
	Flags                     CustomJSONString       `json:"flags"`
	UpdatedBy                 string                 `json:"updatedBy,omitempty"`
	UpdatedAt                 string                 `json:"updatedAt,omitempty"`
	DeletedAt                 string                 `json:"deletedAt,omitempty"`
	FlagsFile                 string                 `json:"flagsFile,omitempty"`
	FlagsFileChecksum         string                 `json:"flagsFileChecksum,omitempty"`
	ActiveFlagProfileID       string                 `json:"activeFlagProfileId,omitempty"`
	ActiveFlagProfileSource   string                 `json:"activeFlagProfileSource,omitempty"`
	ActiveCustomProfileID     string                 `json:"activeCustomProfileId,omitempty"`
	ActiveCustomProfileSource string                 `json:"activeCustomProfileSource,omitempty"`
	Protect                   bool                   `json:"protect"`
	AgentVersion              string                 `json:"agentVersion,omitempty"`
	AgentVersionSortable      string                 `json:"agentVersionSortable,omitempty"`
	AgentID                   string                 `json:"agentId,omitempty"`
	AgentType                 string                 `json:"agentType,omitempty"`
	ResourceType              string                 `json:"resourceType,omitempty"`
	Arch                      string                 `json:"arch,omitempty"`
	Platform                  string                 `json:"platform,omitempty"`
	PlatformLike              string                 `json:"platformLike,omitempty"`
	AssetCapabilities         []AssetCapabilities    `json:"assetCapabilities"`
	Interfaces                []AssetInterface       `json:"interfaces"`
	Links                     []LinkItem             `json:"links"`
	//TagQuerypacks                           []interface{}          `json:"tagQuerypacks"`                           // TODO cant find any examples
	//TagEventExcludeProfiles                 []interface{}          `json:"tagEventExcludeProfiles"`                 // TODO cant find any examples
	//TagFilePathGroups                       []interface{}          `json:"tagFilePathGroups"`                       // TODO cant find any examples
	//TagPrometheusTargets                    []interface{}          `json:"tagPrometheusTargets"`                    // TODO cant find any examples
	//TagRegistryPaths                        []interface{}          `json:"tagRegistryPaths"`                        // TODO cant find any examples
	//TagAtcQueries                           []interface{}          `json:"tagAtcQueries"`                           // TODO cant find any examples
	//TagYaraGroupRules                       []interface{}          `json:"tagYaraGroupRules"`                       // TODO cant find any examples
	//TagRedactions                           []interface{}          `json:"tagRedactions"`                           // TODO cant find any examples
	//PrometheusTargets                       []interface{}          `json:"prometheusTargets"`                       // TODO cant find any examples
	//TagImageLoadExclusions                  []interface{}          `json:"tagImageLoadExclusions"`                  // TODO cant find any examples
	//TagAuditGroups                          []interface{}          `json:"tagAuditGroups"`                          // TODO cant find any examples
	//TagAuditRules                           []interface{}          `json:"tagAuditRules"`                           // TODO cant find any examples
	//FinalAuditConfigurations                []interface{}          `json:"finalAuditConfigurations"`                // TODO cant find any examples
	//Cloud                                   string                 `json:"cloud"`                                   // TODO cant find any examples
	//ActiveComplianceProfileID               string                 `json:"activeComplianceProfileId"`               // TODO cant find any examples
	//ActiveComplianceProfileSource           string                 `json:"activeComplianceProfileSource"`           // TODO cant find any examples
	//AssetInfo                               string                 `json:"assetInfo"`                               // TODO cant find any examples
	//ActiveProcessBlockRuleID                string                 `json:"activeProcessBlockRuleId"`                // TODO cant find any examples
	//ActiveProcessBlockRuleSource            string                 `json:"activeProcessBlockRuleSource"`            // TODO cant find any examples
	//ActiveDNSBlockRuleID                    string                 `json:"activeDnsBlockRuleId"`                    // TODO cant find any examples
	//ActiveDNSBlockRuleSource                string                 `json:"activeDnsBlockRuleSource"`                // TODO cant find any examples
	//ActiveWindowsDefenderPreferenceID       string                 `json:"activeWindowsDefenderPreferenceId"`       // TODO cant find any examples
	//ActiveWindowsDefenderPreferenceIDSource string                 `json:"activeWindowsDefenderPreferenceIdSource"` // TODO cant find any examples
	//AssetLastActivityAt                     struct {               // dupe of .lastActivityAt
	//	LastActivityAt string `json:"lastActivityAt"`
	//} `json:"assetLastActivityAt"`
	//FlagProfile struct { // TODO cant find any examples
	//	ID   string `json:"id"`
	//	Name string `json:"name"`
	//} `json:"flagProfile"`
	//CustomProfile struct { // TODO cant find any examples
	//	ID   interface{} `json:"id"`
	//	Name interface{} `json:"name"`
	//} `json:"customProfile"`
	//AssetUpgrade struct {
	//	UpgradeState   bool   `json:"upgradeState"`
	//	LastUpgradedAt string `json:"lastUpgradedAt"`
	//	//UpgradeAfterTime      interface{} `json:"upgradeAfterTime"` //TODO cant find any examples
	//	//UpgradeOsqueryVersion interface{} `json:"upgradeOsqueryVersion"` //TODO cant find any examples
	//} `json:"assetUpgrade"`
	//Querypacks                  []interface{} `json:"querypacks"`                  // TODO cant find any examples
	//FilePathGroups              []interface{} `json:"filePathGroups"`              // TODO cant find any examples
	//AtcQueries                  []interface{} `json:"atcQueries"`                  // TODO cant find any examples
	//RegistryPaths               []interface{} `json:"registryPaths"`               // TODO cant find any examples
	//Redactions                  []interface{} `json:"redactions"`                  // TODO cant find any examples
	//AuditRules                  []interface{} `json:"auditRules"`                  // TODO cant find any examples
	//AuditConfigurations         []interface{} `json:"auditConfigurations"`         // TODO cant find any examples
	//TagAuditConfigurations      []interface{} `json:"tagAuditConfigurations"`      // TODO cant find any examples
	//YaraGroupRules              []interface{} `json:"yaraGroupRules"`              // TODO cant find any examples
	//ImageLoadExclusions         []interface{} `json:"imageLoadExclusions"`         // TODO cant find any examples
	//AuditGroups                 []interface{} `json:"auditGroups"`                 // TODO cant find any examples
	//UpgradeAfterTime            interface{}   `json:"upgradeAfterTime"`            //TODO cant find any examples
	//FlagProfileID               interface{}   `json:"flagProfileId"`               //TODO cant find any examples
	//CustomProfileID             interface{}   `json:"customProfileId"`             //TODO cant find any examples
	//ProcessBlockRuleID          interface{}   `json:"processBlockRuleId"`          //TODO cant find any examples
	//DNSBlockRuleID              interface{}   `json:"dnsBlockRuleId"`              //TODO cant find any examples
	//WindowsDefenderPreferenceID interface{}   `json:"windowsDefenderPreferenceId"` //TODO cant find any examples
	//SlackUserID                 interface{}   `json:"slackUserId"`                 //TODO cant find any examples
	//QuarantinedStatus           interface{}   `json:"quarantinedStatus"`           //TODO cant find any examples

}

// This wont actually exist, cant get assetTags
type AssetTags struct {
	Links  []LinkItem `json:"links"`
	Items  []AssetTag `json:"items"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

// AssetTag only supports PUT/POST/DELETE, not GET
type AssetTag struct {
	ID   string `json:"tagId,omitempty"`
	Name string `json:"-,omitempty"`
}

type iAPIType interface {
	AlertRule | Destination | EventExcludeProfile | EventRule | User | Role | ObjectGroup | TagConfiguration | TagRule | Tag | FilePathGroup | YaraGroupRule | RegistryPath | Querypack | AuditConfiguration | ComplianceProfile | AlertRuleCategory | AssetGroupRule | AtcQuery | Carve | CustomProfile | FlagProfile | BlockRule | WindowsDefenderPreference | Exception | AssetTag | Asset
	GetID() string
	GetName() string
	KeysToDelete() []string
}

type iAPITypes interface {
	AlertRules | Destinations | EventExcludeProfiles | EventRules | Users | Roles | ObjectGroups | TagConfigurations | TagRules | Tags | FilePathGroups | YaraGroupRules | RegistryPaths | Querypacks | AuditConfigurations | ComplianceProfiles | AlertRuleCategories | AssetGroupRules | AtcQueries | Carves | CustomProfiles | FlagProfiles | BlockRules | WindowsDefenderPreferences | Exceptions | AssetTags | Assets
}
