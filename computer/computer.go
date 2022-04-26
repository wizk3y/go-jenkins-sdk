package computer

// ComputersResponse --
type ComputersResponse struct {
	_              struct{}    `json:"-"`
	Class          string      `json:"_class"`
	BusyExecutors  int         `json:"busyExecutors"`
	Computer       []*Computer `json:"computer"`
	DisplayName    string      `json:"displayName"`
	TotalExecutors int         `json:"totalExecutors"`
}

// Actions --
type Actions struct {
	_     struct{} `json:"-"`
	Class string   `json:"_class"`
}

// AssignedLabels --
type AssignedLabels struct {
	_    struct{} `json:"-"`
	Name string   `json:"name"`
}

// SubBuilds --
type SubBuilds struct {
	_                 struct{}    `json:"-"`
	Abort             bool        `json:"abort"`
	Build             interface{} `json:"build"`
	BuildNumber       int         `json:"buildNumber"`
	Duration          string      `json:"duration"`
	Icon              string      `json:"icon"`
	JobName           string      `json:"jobName"`
	ParentBuildNumber int         `json:"parentBuildNumber"`
	ParentJobName     string      `json:"parentJobName"`
	PhaseName         string      `json:"phaseName"`
	Result            string      `json:"result"`
	Retry             bool        `json:"retry"`
	URL               string      `json:"url"`
}

// CurrentExecutable --
type CurrentExecutable struct {
	_         struct{}    `json:"-"`
	Number    int         `json:"number"`
	URL       string      `json:"url"`
	SubBuilds []SubBuilds `json:"subBuilds"`
}

// Executors --
type Executors struct {
	_                 struct{}          `json:"-"`
	CurrentExecutable CurrentExecutable `json:"currentExecutable"`
}

// LoadStatistics --
type LoadStatistics struct {
	_     struct{} `json:"-"`
	Class string   `json:"_class"`
}

// HudsonNodeMonitorsSwapSpaceMonitor --
type HudsonNodeMonitorsSwapSpaceMonitor struct {
	_                       struct{} `json:"-"`
	Class                   string   `json:"_class"`
	AvailablePhysicalMemory int      `json:"availablePhysicalMemory"`
	AvailableSwapSpace      int      `json:"availableSwapSpace"`
	TotalPhysicalMemory     int64    `json:"totalPhysicalMemory"`
	TotalSwapSpace          int      `json:"totalSwapSpace"`
}

// HudsonNodeMonitorsTemporarySpaceMonitor --
type HudsonNodeMonitorsTemporarySpaceMonitor struct {
	_         struct{} `json:"-"`
	Class     string   `json:"_class"`
	Timestamp int64    `json:"timestamp"`
	Path      string   `json:"path"`
	Size      int64    `json:"size"`
}

// HudsonNodeMonitorsDiskSpaceMonitor --
type HudsonNodeMonitorsDiskSpaceMonitor struct {
	_         struct{} `json:"-"`
	Class     string   `json:"_class"`
	Timestamp int64    `json:"timestamp"`
	Path      string   `json:"path"`
	Size      int64    `json:"size"`
}

// HudsonNodeMonitorsResponseTimeMonitor --
type HudsonNodeMonitorsResponseTimeMonitor struct {
	_         struct{} `json:"-"`
	Class     string   `json:"_class"`
	Timestamp int64    `json:"timestamp"`
	Average   int      `json:"average"`
}

// HudsonNodeMonitorsClockMonitor --
type HudsonNodeMonitorsClockMonitor struct {
	_     struct{} `json:"-"`
	Class string   `json:"_class"`
	Diff  int      `json:"diff"`
}

// MonitorData --
type MonitorData struct {
	_                                       struct{}                                `json:"-"`
	HudsonNodeMonitorsSwapSpaceMonitor      HudsonNodeMonitorsSwapSpaceMonitor      `json:"hudson.node_monitors.SwapSpaceMonitor"`
	HudsonNodeMonitorsTemporarySpaceMonitor HudsonNodeMonitorsTemporarySpaceMonitor `json:"hudson.node_monitors.TemporarySpaceMonitor"`
	HudsonNodeMonitorsDiskSpaceMonitor      HudsonNodeMonitorsDiskSpaceMonitor      `json:"hudson.node_monitors.DiskSpaceMonitor"`
	HudsonNodeMonitorsArchitectureMonitor   string                                  `json:"hudson.node_monitors.ArchitectureMonitor"`
	HudsonNodeMonitorsResponseTimeMonitor   HudsonNodeMonitorsResponseTimeMonitor   `json:"hudson.node_monitors.ResponseTimeMonitor"`
	HudsonNodeMonitorsClockMonitor          HudsonNodeMonitorsClockMonitor          `json:"hudson.node_monitors.ClockMonitor"`
}

// OfflineCause --
type OfflineCause struct {
	_     struct{} `json:"-"`
	Class string   `json:"_class"`
}

// OneOffExecutors --
type OneOffExecutors struct {
	_ struct{} `json:"-"`
}

// Computer --
type Computer struct {
	_                   struct{}          `json:"-"`
	Class               string            `json:"_class"`
	Actions             []Actions         `json:"actions"`
	AssignedLabels      []AssignedLabels  `json:"assignedLabels"`
	Description         string            `json:"description"`
	DisplayName         string            `json:"displayName"`
	Executors           []Executors       `json:"executors"`
	Icon                string            `json:"icon"`
	IconClassName       string            `json:"iconClassName"`
	Idle                bool              `json:"idle"`
	JnlpAgent           bool              `json:"jnlpAgent"`
	LaunchSupported     bool              `json:"launchSupported"`
	LoadStatistics      LoadStatistics    `json:"loadStatistics"`
	ManualLaunchAllowed bool              `json:"manualLaunchAllowed"`
	MonitorData         MonitorData       `json:"monitorData"`
	NumExecutors        int               `json:"numExecutors"`
	Offline             bool              `json:"offline"`
	OfflineCause        *OfflineCause     `json:"offlineCause"`
	OfflineCauseReason  string            `json:"offlineCauseReason"`
	OneOffExecutors     []OneOffExecutors `json:"oneOffExecutors"`
	TemporarilyOffline  bool              `json:"temporarilyOffline"`
	AbsoluteRemotePath  interface{}       `json:"absoluteRemotePath,omitempty"`
}

// ComputerRequest --
type ComputerRequest struct {
	_                 struct{}          `json:"-"`
	Name              string            `json:"name"`
	NodeDescription   string            `json:"nodeDescription"`
	NumExecutors      string            `json:"numExecutors"`
	RemoteFS          string            `json:"remoteFS"`
	LabelString       string            `json:"labelString"`
	Mode              string            `json:"mode"`
	Launcher          Launcher          `json:"launcher"`
	RetentionStrategy RetentionStrategy `json:"retentionStrategy"`
	NodeProperties    NodeProperties    `json:"nodeProperties"`
	Type              string            `json:"type"`
}

// SSHHostKeyVerificationStrategy --
type SSHHostKeyVerificationStrategy struct {
	_            struct{} `json:"-"`
	StaplerClass string   `json:"stapler-class"`
	Class        string   `json:"$class"`
}

// Launcher --
type Launcher struct {
	_                              struct{}                       `json:"-"`
	StaplerClass                   string                         `json:"stapler-class"`
	Class                          string                         `json:"$class"`
	Host                           string                         `json:"host"`
	IncludeUser                    string                         `json:"includeUser"`
	CredentialsID                  string                         `json:"credentialsId"`
	SSHHostKeyVerificationStrategy SSHHostKeyVerificationStrategy `json:"sshHostKeyVerificationStrategy"`
	Port                           string                         `json:"port"`
	JavaPath                       string                         `json:"javaPath"`
	JvmOptions                     string                         `json:"jvmOptions"`
	PrefixStartSlaveCmd            string                         `json:"prefixStartSlaveCmd"`
	SuffixStartSlaveCmd            string                         `json:"suffixStartSlaveCmd"`
	LaunchTimeoutSeconds           string                         `json:"launchTimeoutSeconds"`
	MaxNumRetries                  string                         `json:"maxNumRetries"`
	RetryWaitTime                  string                         `json:"retryWaitTime"`
	TCPNoDelay                     bool                           `json:"tcpNoDelay"`
	WorkDir                        string                         `json:"workDir"`
}

// RetentionStrategy --
type RetentionStrategy struct {
	_            struct{} `json:"-"`
	StaplerClass string   `json:"stapler-class"`
	Class        string   `json:"$class"`
}

// NodeProperties --
type NodeProperties struct {
	_               struct{} `json:"-"`
	StaplerClassBag string   `json:"stapler-class-bag"`
}
