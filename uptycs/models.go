package uptycs

type EventRules struct {
	Links  []LinkItem  `json:"links,omitempty"`
	Items  []EventRule `json:"items,omitempty"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type ScriptConfig struct {
	ID          string `json:"id,omitempty"`
	QueryPackID string `json:"querypackId,omitempty"`
	TableName   string `json:"tableName,omitempty"`
	Added       bool   `json:"added,omitempty"`
}

type EventRule struct {
	ID            string        `json:"id,omitempty"`
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
	ID                  string        `json:"id,omitempty"`
	Name                string        `json:"name"`
	Email               string        `json:"email,omitempty"`
	Phone               string        `json:"phone,omitempty"`
	Active              bool          `json:"active"`
	SuperAdmin          bool          `json:"superAdmin"`
	Bot                 bool          `json:"bot"`
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
	Links  []LinkItem `json:"links,omitempty"`
	Items  []User     `json:"items,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type Role struct {
	ID                   string        `json:"id,omitempty"`
	Name                 string        `json:"name"`
	Description          string        `json:"description,omitempty"`
	Permissions          []string      `json:"permissions" validate:"required,min=0"`
	Custom               bool          `json:"custom"`
	Hidden               bool          `json:"hidden"`
	CreatedBy            string        `json:"createdBy"`
	UpdatedBy            string        `json:"updatedBy"`
	CreatedAt            string        `json:"createdAt"`
	UpdatedAt            string        `json:"updatedAt"`
	NoMinimalPermissions bool          `json:"noMinimalPermissions"`
	RoleObjectGroups     []ObjectGroup `json:"roleObjectGroups" validate:"required,min=0"`
}
type Roles struct {
	Links  []LinkItem `json:"links,omitempty"`
	Items  []Role     `json:"items,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type ComplianceProfile struct {
	ID          string     `json:"id"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Custom      bool       `json:"custom"`
	Priority    int        `json:"priority,omitempty"`
	CreatedBy   string     `json:"createdBy"`
	UpdatedBy   string     `json:"updatedBy"`
	CreatedAt   string     `json:"createdAt"`
	UpdatedAt   string     `json:"updatedAt"`
	Links       []LinkItem `json:"links"`
}

type ComplianceProfiles struct {
	Links      []LinkItem          `json:"links,omitempty"`
	Items      []ComplianceProfile `json:"items,omitempty"`
	Offset     int                 `json:"offset,omitempty"`
	Limit      int                 `json:"limit,omitempty"`
	Decorators []string            `json:"decorators,omitempty"`
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

type TagConfigurations struct {
	Links  []LinkItem         `json:"links,omitempty"`
	Items  []TagConfiguration `json:"items,omitempty"`
	Offset int                `json:"offset,omitempty"`
	Limit  int                `json:"limit,omitempty"`
}
type TagConfiguration Tag

type TagRules struct {
	Links  []LinkItem `json:"links,omitempty"`
	Items  []TagRule  `json:"items,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type TagRule struct {
	ID             string `json:"id,omitempty"`
	Name           string `json:"name,omitempty" validate:"required,max=255,min=1"`
	Description    string `json:"description,omitempty"`
	Query          string `json:"query,omitempty" validate:"required"`
	Source         string `json:"source,omitempty" validate:"required,oneof=global realtime"`
	RunOnce        bool   `json:"runOnce,omitempty"`
	Interval       int    `json:"interval,omitempty" validate:"min=30,excluded_if=RunOnce false"`
	OSqueryVersion string `json:"osqueryVersion,omitempty"`
	Platform       string `json:"platform" validate:"required_if=Source realtime"`
	Enabled        bool   `json:"enabled,omitempty"`
	System         bool   `json:"system,omitempty"`
	LastRunAt      string `json:"lastRunAt,omitempty"`
	CreatedBy      string `json:"createdBy,omitempty"`
	UpdatedBy      string `json:"updatedBy,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	UpdatedAt      string `json:"updatedAt,omitempty"`
	ResourceType   string `json:"resourceType,omitempty"`
}

type Tags struct {
	Links  []LinkItem `json:"links,omitempty"`
	Items  []Tag      `json:"items,omitempty"`
	Offset int        `json:"offset,omitempty"`
	Limit  int        `json:"limit,omitempty"`
}

type Tag struct {
	ID                          string                   `json:"id,omitempty"`
	Name                        string                   `json:"-"` // There is no name, so just throw the ID in there
	Value                       string                   `json:"value"`
	Key                         string                   `json:"key,omitempty"`
	CreatedBy                   string                   `json:"createdBy,omitempty"`
	UpdatedBy                   string                   `json:"updatedBy,omitempty"`
	FlagProfileID               string                   `json:"flagProfileId,omitempty"`
	CustomProfileID             string                   `json:"customProfileId,omitempty"`
	ComplianceProfileID         string                   `json:"complianceProfileId,omitempty"`
	ProcessBlockRuleID          string                   `json:"processBlockRuleId,omitempty"`
	DNSBlockRuleID              string                   `json:"dnsBlockRuleId,omitempty"`
	WindowsDefenderPreferenceID string                   `json:"windowsDefenderPreferenceId,omitempty"`
	Tag                         string                   `json:"tag,omitempty"`
	Custom                      bool                     `json:"custom,omitempty"`
	System                      bool                     `json:"system,omitempty"`
	CreatedAt                   string                   `json:"createdAt,omitempty"`
	TagRuleID                   string                   `json:"tagRuleId,omitempty"`
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
	Links  []LinkItem      `json:"links,omitempty"`
	Items  []FilePathGroup `json:"items,omitempty"`
	Offset int             `json:"offset,omitempty"`
	Limit  int             `json:"limit,omitempty"`
}

type FilePathGroup struct {
	ID                    string   `json:"id,omitempty"`
	Name                  string   `json:"name,omitempty"`
	Description           string   `json:"description,omitempty"`
	Grouping              string   `json:"grouping,omitempty"`
	IncludePaths          []string `json:"includePaths"`
	IncludePathExtensions []string `json:"includePathExtensions"`
	ExcludePaths          []string `json:"excludePaths"`
	Custom                bool     `json:"custom" validate:"required"`
	CheckSignature        bool     `json:"checkSignature"`
	FileAccesses          bool     `json:"fileAccesses"`
	//ExcludeProcessPaths   []string                 `json:"excludeProcessPaths"` //TODO this seems broken in the API. returns null or {}
	ExcludeProcessNames []string                 `json:"excludeProcessNames"`
	PriorityPaths       []string                 `json:"priorityPaths"`
	CreatedBy           string                   `json:"createdBy,omitempty"`
	UpdatedBy           string                   `json:"updatedBy,omitempty"`
	CreatedAt           string                   `json:"createdAt,omitempty"`
	UpdatedAt           string                   `json:"updatedAt,omitempty"`
	Signatures          []FilePathGroupSignature `json:"signatures" validate:"required,min=0"`
	YaraGroupRules      []YaraGroupRule          `json:"yaraGroupRules" validate:"required,min=0"`
	Links               []LinkItem               `json:"links"`
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
	Links  []LinkItem      `json:"links,omitempty"`
	Items  []YaraGroupRule `json:"items,omitempty"`
	Offset int             `json:"offset,omitempty"`
	Limit  int             `json:"limit,omitempty"`
}

type YaraGroupRule struct {
	ID          string     `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description string     `json:"description,omitempty"`
	Rules       string     `json:"rules,omitempty"`
	Custom      bool       `json:"custom,omitempty"`
	CreatedBy   string     `json:"createdBy,omitempty"`
	UpdatedBy   string     `json:"updatedBy,omitempty"`
	CreatedAt   string     `json:"createdAt,omitempty"`
	UpdatedAt   string     `json:"updatedAt,omitempty"`
	Links       []LinkItem `json:"links,omitempty"`
}

type RegistryPaths struct {
	Links  []LinkItem     `json:"links,omitempty"`
	Items  []RegistryPath `json:"items,omitempty"`
	Offset int            `json:"offset,omitempty"`
	Limit  int            `json:"limit,omitempty"`
}

type RegistryPath struct {
	ID                   string     `json:"id"`
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
	Links  []LinkItem  `json:"links,omitempty"`
	Items  []Querypack `json:"items,omitempty"`
	Offset int         `json:"offset,omitempty"`
	Limit  int         `json:"limit,omitempty"`
}

type Querypack struct {
	ID string `json:"id"`
	//Sha              string           `json:"sha" validate:"required,max=40,min=1"` //TODO Does not work
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
	Queries          []Query          `json:"queries,omitempty"`
	Conf             CustomJSONString `json:"conf,omitempty"`
	Links            []LinkItem       `json:"links"`
}

type Query struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Query       string           `json:"query"`
	Removed     bool             `json:"removed"`
	Version     string           `json:"version"`
	Interval    int              `json:"interval"`
	Platform    string           `json:"platform"`
	Snapshot    bool             `json:"snapshot"`
	RunNow      bool             `json:"runNow"`
	Value       string           `json:"value"`
	QuerypackID string           `json:"querypackId"`
	TableName   string           `json:"tableName"`
	DataTypes   CustomJSONString `json:"dataTypes,omitempty"` //This is super ephemeral
	Verified    bool             `json:"verified"`
	CreatedBy   string           `json:"createdBy"`
	UpdatedBy   string           `json:"updatedBy"`
	CreatedAt   string           `json:"createdAt"`
	UpdatedAt   string           `json:"updatedAt"`
}

type AuditConfigurations struct {
	Links  []LinkItem           `json:"links,omitempty"`
	Items  []AuditConfiguration `json:"items,omitempty"`
	Offset int                  `json:"offset,omitempty"`
	Limit  int                  `json:"limit,omitempty"`
}

type AuditConfiguration struct {
	ID          string       `json:"id"`
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
	Links       []LinkItem   `json:"links"`
}

type AuditEntry struct {
	ID                   string   `json:"id"`
	AuditConfigurationID string   `json:"auditConfigurationId"`
	AuditName            []string `json:"auditName"`
	Standard             string   `json:"standard"`
	Version              string   `json:"version"`
	Section              string   `json:"section"`
	Title                string   `json:"title"`
	Scored               bool     `json:"scored"`
	Level                string   `json:"level"`
	Description          string   `json:"description"`
	Rationale            string   `json:"rationale"`
	Command              string   `json:"command"`
	Remediation          string   `json:"remediation"`
	ExpectedValue        string   `json:"expectedValue"`
	AuthoritativeSource  string   `json:"authoritativeSource"`
	Exception            string   `json:"exception"`
	Chapter              string   `json:"chapter"`
	CheckID              string   `json:"checkId"`
	Enabled              bool     `json:"enabled"`
	Service              string   `json:"service"`
	CreatedBy            string   `json:"createdBy"`
	Score                float64  `json:"score"`
	UpdatedBy            string   `json:"updatedBy"`
	RunCategory          int      `json:"runCategory"`
	Timeout              int      `json:"timeout"`
	CreatedAt            string   `json:"createdAt"`
	UpdatedAt            string   `json:"updatedAt"`
	IsManual             bool     `json:"isManual"`
	//RemediationAction    interface{} `json:"remediationAction"`// TODO there are no examples of this
	//Parameters           interface{} `json:"parameters"` // TODO there are no examples of this
}

type iAPIType interface {
	AlertRule | Destination | EventExcludeProfile | EventRule | User | Role | ObjectGroup | TagConfiguration | TagRule | Tag | FilePathGroup | YaraGroupRule | RegistryPath | Querypack | AuditConfiguration | ComplianceProfile
	GetID() string
	GetName() string
	KeysToDelete() []string
}

type iAPITypes interface {
	AlertRules | Destinations | EventExcludeProfiles | EventRules | Users | Roles | ObjectGroups | TagConfigurations | TagRules | Tags | FilePathGroups | YaraGroupRules | RegistryPaths | Querypacks | AuditConfigurations | ComplianceProfiles
}
