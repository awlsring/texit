// Code generated by ogen, DO NOT EDIT.

package texit

import (
	"github.com/go-faster/errors"
)

// Ref: #/components/schemas/DeprovisionNodeResponseContent
type DeprovisionNodeResponseContent struct {
	// A node's identifier.
	Execution string `json:"execution"`
}

// GetExecution returns the value of Execution.
func (s *DeprovisionNodeResponseContent) GetExecution() string {
	return s.Execution
}

// SetExecution sets the value of Execution.
func (s *DeprovisionNodeResponseContent) SetExecution(val string) {
	s.Execution = val
}

func (*DeprovisionNodeResponseContent) deprovisionNodeRes() {}

// Ref: #/components/schemas/DescribeNodeResponseContent
type DescribeNodeResponseContent struct {
	Summary NodeSummary `json:"summary"`
}

// GetSummary returns the value of Summary.
func (s *DescribeNodeResponseContent) GetSummary() NodeSummary {
	return s.Summary
}

// SetSummary sets the value of Summary.
func (s *DescribeNodeResponseContent) SetSummary(val NodeSummary) {
	s.Summary = val
}

func (*DescribeNodeResponseContent) describeNodeRes() {}

// Ref: #/components/schemas/DescribeProviderResponseContent
type DescribeProviderResponseContent struct {
	Summary ProviderSummary `json:"summary"`
}

// GetSummary returns the value of Summary.
func (s *DescribeProviderResponseContent) GetSummary() ProviderSummary {
	return s.Summary
}

// SetSummary sets the value of Summary.
func (s *DescribeProviderResponseContent) SetSummary(val ProviderSummary) {
	s.Summary = val
}

func (*DescribeProviderResponseContent) describeProviderRes() {}

// Ref: #/components/schemas/DescribeTailnetResponseContent
type DescribeTailnetResponseContent struct {
	Summary TailnetSummary `json:"summary"`
}

// GetSummary returns the value of Summary.
func (s *DescribeTailnetResponseContent) GetSummary() TailnetSummary {
	return s.Summary
}

// SetSummary sets the value of Summary.
func (s *DescribeTailnetResponseContent) SetSummary(val TailnetSummary) {
	s.Summary = val
}

func (*DescribeTailnetResponseContent) describeTailnetRes() {}

// The status of an execution.
// Ref: #/components/schemas/ExecutionStatus
type ExecutionStatus string

const (
	ExecutionStatusPending   ExecutionStatus = "pending"
	ExecutionStatusRunning   ExecutionStatus = "running"
	ExecutionStatusCompleted ExecutionStatus = "completed"
	ExecutionStatusFailed    ExecutionStatus = "failed"
	ExecutionStatusUnknown   ExecutionStatus = "unknown"
)

// AllValues returns all ExecutionStatus values.
func (ExecutionStatus) AllValues() []ExecutionStatus {
	return []ExecutionStatus{
		ExecutionStatusPending,
		ExecutionStatusRunning,
		ExecutionStatusCompleted,
		ExecutionStatusFailed,
		ExecutionStatusUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ExecutionStatus) MarshalText() ([]byte, error) {
	switch s {
	case ExecutionStatusPending:
		return []byte(s), nil
	case ExecutionStatusRunning:
		return []byte(s), nil
	case ExecutionStatusCompleted:
		return []byte(s), nil
	case ExecutionStatusFailed:
		return []byte(s), nil
	case ExecutionStatusUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ExecutionStatus) UnmarshalText(data []byte) error {
	switch ExecutionStatus(data) {
	case ExecutionStatusPending:
		*s = ExecutionStatusPending
		return nil
	case ExecutionStatusRunning:
		*s = ExecutionStatusRunning
		return nil
	case ExecutionStatusCompleted:
		*s = ExecutionStatusCompleted
		return nil
	case ExecutionStatusFailed:
		*s = ExecutionStatusFailed
		return nil
	case ExecutionStatusUnknown:
		*s = ExecutionStatusUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/ExecutionSummary
type ExecutionSummary struct {
	// A node's identifier.
	Identifier string          `json:"identifier"`
	Status     ExecutionStatus `json:"status"`
	Workflow   WorkflowName    `json:"workflow"`
	StartedAt  float64         `json:"startedAt"`
	EndedAt    OptFloat64      `json:"endedAt"`
	Result     OptString       `json:"result"`
}

// GetIdentifier returns the value of Identifier.
func (s *ExecutionSummary) GetIdentifier() string {
	return s.Identifier
}

// GetStatus returns the value of Status.
func (s *ExecutionSummary) GetStatus() ExecutionStatus {
	return s.Status
}

// GetWorkflow returns the value of Workflow.
func (s *ExecutionSummary) GetWorkflow() WorkflowName {
	return s.Workflow
}

// GetStartedAt returns the value of StartedAt.
func (s *ExecutionSummary) GetStartedAt() float64 {
	return s.StartedAt
}

// GetEndedAt returns the value of EndedAt.
func (s *ExecutionSummary) GetEndedAt() OptFloat64 {
	return s.EndedAt
}

// GetResult returns the value of Result.
func (s *ExecutionSummary) GetResult() OptString {
	return s.Result
}

// SetIdentifier sets the value of Identifier.
func (s *ExecutionSummary) SetIdentifier(val string) {
	s.Identifier = val
}

// SetStatus sets the value of Status.
func (s *ExecutionSummary) SetStatus(val ExecutionStatus) {
	s.Status = val
}

// SetWorkflow sets the value of Workflow.
func (s *ExecutionSummary) SetWorkflow(val WorkflowName) {
	s.Workflow = val
}

// SetStartedAt sets the value of StartedAt.
func (s *ExecutionSummary) SetStartedAt(val float64) {
	s.StartedAt = val
}

// SetEndedAt sets the value of EndedAt.
func (s *ExecutionSummary) SetEndedAt(val OptFloat64) {
	s.EndedAt = val
}

// SetResult sets the value of Result.
func (s *ExecutionSummary) SetResult(val OptString) {
	s.Result = val
}

// Ref: #/components/schemas/GetExecutionResponseContent
type GetExecutionResponseContent struct {
	Summary ExecutionSummary `json:"summary"`
}

// GetSummary returns the value of Summary.
func (s *GetExecutionResponseContent) GetSummary() ExecutionSummary {
	return s.Summary
}

// SetSummary sets the value of Summary.
func (s *GetExecutionResponseContent) SetSummary(val ExecutionSummary) {
	s.Summary = val
}

func (*GetExecutionResponseContent) getExecutionRes() {}

// Ref: #/components/schemas/GetNodeStatusResponseContent
type GetNodeStatusResponseContent struct {
	Status NodeStatus `json:"status"`
}

// GetStatus returns the value of Status.
func (s *GetNodeStatusResponseContent) GetStatus() NodeStatus {
	return s.Status
}

// SetStatus sets the value of Status.
func (s *GetNodeStatusResponseContent) SetStatus(val NodeStatus) {
	s.Status = val
}

func (*GetNodeStatusResponseContent) getNodeStatusRes() {}

// Ref: #/components/schemas/HealthResponseContent
type HealthResponseContent struct {
	Healthy bool `json:"healthy"`
}

// GetHealthy returns the value of Healthy.
func (s *HealthResponseContent) GetHealthy() bool {
	return s.Healthy
}

// SetHealthy sets the value of Healthy.
func (s *HealthResponseContent) SetHealthy(val bool) {
	s.Healthy = val
}

// Ref: #/components/schemas/InvalidInputErrorResponseContent
type InvalidInputErrorResponseContent struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *InvalidInputErrorResponseContent) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *InvalidInputErrorResponseContent) SetMessage(val string) {
	s.Message = val
}

func (*InvalidInputErrorResponseContent) deprovisionNodeRes()  {}
func (*InvalidInputErrorResponseContent) describeNodeRes()     {}
func (*InvalidInputErrorResponseContent) describeProviderRes() {}
func (*InvalidInputErrorResponseContent) describeTailnetRes()  {}
func (*InvalidInputErrorResponseContent) getExecutionRes()     {}
func (*InvalidInputErrorResponseContent) getNodeStatusRes()    {}
func (*InvalidInputErrorResponseContent) provisionNodeRes()    {}
func (*InvalidInputErrorResponseContent) startNodeRes()        {}
func (*InvalidInputErrorResponseContent) stopNodeRes()         {}

// Ref: #/components/schemas/ListNodesResponseContent
type ListNodesResponseContent struct {
	Summaries []NodeSummary `json:"summaries"`
}

// GetSummaries returns the value of Summaries.
func (s *ListNodesResponseContent) GetSummaries() []NodeSummary {
	return s.Summaries
}

// SetSummaries sets the value of Summaries.
func (s *ListNodesResponseContent) SetSummaries(val []NodeSummary) {
	s.Summaries = val
}

// Ref: #/components/schemas/ListNotifiersResponseContent
type ListNotifiersResponseContent struct {
	Summaries []NotifierSummary `json:"summaries"`
}

// GetSummaries returns the value of Summaries.
func (s *ListNotifiersResponseContent) GetSummaries() []NotifierSummary {
	return s.Summaries
}

// SetSummaries sets the value of Summaries.
func (s *ListNotifiersResponseContent) SetSummaries(val []NotifierSummary) {
	s.Summaries = val
}

// Ref: #/components/schemas/ListProvidersResponseContent
type ListProvidersResponseContent struct {
	Summaries []ProviderSummary `json:"summaries"`
}

// GetSummaries returns the value of Summaries.
func (s *ListProvidersResponseContent) GetSummaries() []ProviderSummary {
	return s.Summaries
}

// SetSummaries sets the value of Summaries.
func (s *ListProvidersResponseContent) SetSummaries(val []ProviderSummary) {
	s.Summaries = val
}

// Ref: #/components/schemas/ListTailnetsResponseContent
type ListTailnetsResponseContent struct {
	Summaries []TailnetSummary `json:"summaries"`
}

// GetSummaries returns the value of Summaries.
func (s *ListTailnetsResponseContent) GetSummaries() []TailnetSummary {
	return s.Summaries
}

// SetSummaries sets the value of Summaries.
func (s *ListTailnetsResponseContent) SetSummaries(val []TailnetSummary) {
	s.Summaries = val
}

// The size a node. Size are abstracted so that a provider can define what to provision for each.
// Ref: #/components/schemas/NodeSize
type NodeSize string

const (
	NodeSizeSmall   NodeSize = "small"
	NodeSizeMedium  NodeSize = "medium"
	NodeSizeLarge   NodeSize = "large"
	NodeSizeUnknown NodeSize = "unknown"
)

// AllValues returns all NodeSize values.
func (NodeSize) AllValues() []NodeSize {
	return []NodeSize{
		NodeSizeSmall,
		NodeSizeMedium,
		NodeSizeLarge,
		NodeSizeUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s NodeSize) MarshalText() ([]byte, error) {
	switch s {
	case NodeSizeSmall:
		return []byte(s), nil
	case NodeSizeMedium:
		return []byte(s), nil
	case NodeSizeLarge:
		return []byte(s), nil
	case NodeSizeUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *NodeSize) UnmarshalText(data []byte) error {
	switch NodeSize(data) {
	case NodeSizeSmall:
		*s = NodeSizeSmall
		return nil
	case NodeSizeMedium:
		*s = NodeSizeMedium
		return nil
	case NodeSizeLarge:
		*s = NodeSizeLarge
		return nil
	case NodeSizeUnknown:
		*s = NodeSizeUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// The status of a node.
// Ref: #/components/schemas/NodeStatus
type NodeStatus string

const (
	NodeStatusPending  NodeStatus = "pending"
	NodeStatusStarting NodeStatus = "starting"
	NodeStatusRunning  NodeStatus = "running"
	NodeStatusStopping NodeStatus = "stopping"
	NodeStatusStopped  NodeStatus = "stopped"
	NodeStatusUnknown  NodeStatus = "unknown"
)

// AllValues returns all NodeStatus values.
func (NodeStatus) AllValues() []NodeStatus {
	return []NodeStatus{
		NodeStatusPending,
		NodeStatusStarting,
		NodeStatusRunning,
		NodeStatusStopping,
		NodeStatusStopped,
		NodeStatusUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s NodeStatus) MarshalText() ([]byte, error) {
	switch s {
	case NodeStatusPending:
		return []byte(s), nil
	case NodeStatusStarting:
		return []byte(s), nil
	case NodeStatusRunning:
		return []byte(s), nil
	case NodeStatusStopping:
		return []byte(s), nil
	case NodeStatusStopped:
		return []byte(s), nil
	case NodeStatusUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *NodeStatus) UnmarshalText(data []byte) error {
	switch NodeStatus(data) {
	case NodeStatusPending:
		*s = NodeStatusPending
		return nil
	case NodeStatusStarting:
		*s = NodeStatusStarting
		return nil
	case NodeStatusRunning:
		*s = NodeStatusRunning
		return nil
	case NodeStatusStopping:
		*s = NodeStatusStopping
		return nil
	case NodeStatusStopped:
		*s = NodeStatusStopped
		return nil
	case NodeStatusUnknown:
		*s = NodeStatusUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/NodeSummary
type NodeSummary struct {
	// A node's identifier.
	Identifier string `json:"identifier"`
	// The name of the provider.
	Provider string `json:"provider"`
	// A location provided by a provider.
	Location string `json:"location"`
	// The identifier of the node resource in the provider.
	ProviderNodeIdentifier string `json:"providerNodeIdentifier"`
	// .
	Tailnet string `json:"tailnet"`
	// The name of a tailnet device.
	TailnetDeviceName string `json:"tailnetDeviceName"`
	// The identifier of a tailnet device.
	TailnetDeviceIdentifier string   `json:"tailnetDeviceIdentifier"`
	Size                    NodeSize `json:"size"`
	// If a node is ephemeral.
	Ephemeral bool `json:"ephemeral"`
	// When a node was created.
	Created float64 `json:"created"`
	// When a node was last updated.
	Updated            float64            `json:"updated"`
	ProvisioningStatus ProvisioningStatus `json:"provisioningStatus"`
}

// GetIdentifier returns the value of Identifier.
func (s *NodeSummary) GetIdentifier() string {
	return s.Identifier
}

// GetProvider returns the value of Provider.
func (s *NodeSummary) GetProvider() string {
	return s.Provider
}

// GetLocation returns the value of Location.
func (s *NodeSummary) GetLocation() string {
	return s.Location
}

// GetProviderNodeIdentifier returns the value of ProviderNodeIdentifier.
func (s *NodeSummary) GetProviderNodeIdentifier() string {
	return s.ProviderNodeIdentifier
}

// GetTailnet returns the value of Tailnet.
func (s *NodeSummary) GetTailnet() string {
	return s.Tailnet
}

// GetTailnetDeviceName returns the value of TailnetDeviceName.
func (s *NodeSummary) GetTailnetDeviceName() string {
	return s.TailnetDeviceName
}

// GetTailnetDeviceIdentifier returns the value of TailnetDeviceIdentifier.
func (s *NodeSummary) GetTailnetDeviceIdentifier() string {
	return s.TailnetDeviceIdentifier
}

// GetSize returns the value of Size.
func (s *NodeSummary) GetSize() NodeSize {
	return s.Size
}

// GetEphemeral returns the value of Ephemeral.
func (s *NodeSummary) GetEphemeral() bool {
	return s.Ephemeral
}

// GetCreated returns the value of Created.
func (s *NodeSummary) GetCreated() float64 {
	return s.Created
}

// GetUpdated returns the value of Updated.
func (s *NodeSummary) GetUpdated() float64 {
	return s.Updated
}

// GetProvisioningStatus returns the value of ProvisioningStatus.
func (s *NodeSummary) GetProvisioningStatus() ProvisioningStatus {
	return s.ProvisioningStatus
}

// SetIdentifier sets the value of Identifier.
func (s *NodeSummary) SetIdentifier(val string) {
	s.Identifier = val
}

// SetProvider sets the value of Provider.
func (s *NodeSummary) SetProvider(val string) {
	s.Provider = val
}

// SetLocation sets the value of Location.
func (s *NodeSummary) SetLocation(val string) {
	s.Location = val
}

// SetProviderNodeIdentifier sets the value of ProviderNodeIdentifier.
func (s *NodeSummary) SetProviderNodeIdentifier(val string) {
	s.ProviderNodeIdentifier = val
}

// SetTailnet sets the value of Tailnet.
func (s *NodeSummary) SetTailnet(val string) {
	s.Tailnet = val
}

// SetTailnetDeviceName sets the value of TailnetDeviceName.
func (s *NodeSummary) SetTailnetDeviceName(val string) {
	s.TailnetDeviceName = val
}

// SetTailnetDeviceIdentifier sets the value of TailnetDeviceIdentifier.
func (s *NodeSummary) SetTailnetDeviceIdentifier(val string) {
	s.TailnetDeviceIdentifier = val
}

// SetSize sets the value of Size.
func (s *NodeSummary) SetSize(val NodeSize) {
	s.Size = val
}

// SetEphemeral sets the value of Ephemeral.
func (s *NodeSummary) SetEphemeral(val bool) {
	s.Ephemeral = val
}

// SetCreated sets the value of Created.
func (s *NodeSummary) SetCreated(val float64) {
	s.Created = val
}

// SetUpdated sets the value of Updated.
func (s *NodeSummary) SetUpdated(val float64) {
	s.Updated = val
}

// SetProvisioningStatus sets the value of ProvisioningStatus.
func (s *NodeSummary) SetProvisioningStatus(val ProvisioningStatus) {
	s.ProvisioningStatus = val
}

// Ref: #/components/schemas/NotifierSummary
type NotifierSummary struct {
	// The name of the notifier.
	Name string       `json:"name"`
	Type NotifierType `json:"type"`
	// The endpoint that is used to send notifications. For SNS, this is the topic arn. For MQTT, this is
	// the broker address and topic.
	Endpoint string `json:"endpoint"`
}

// GetName returns the value of Name.
func (s *NotifierSummary) GetName() string {
	return s.Name
}

// GetType returns the value of Type.
func (s *NotifierSummary) GetType() NotifierType {
	return s.Type
}

// GetEndpoint returns the value of Endpoint.
func (s *NotifierSummary) GetEndpoint() string {
	return s.Endpoint
}

// SetName sets the value of Name.
func (s *NotifierSummary) SetName(val string) {
	s.Name = val
}

// SetType sets the value of Type.
func (s *NotifierSummary) SetType(val NotifierType) {
	s.Type = val
}

// SetEndpoint sets the value of Endpoint.
func (s *NotifierSummary) SetEndpoint(val string) {
	s.Endpoint = val
}

// The type of notifier.
// Ref: #/components/schemas/NotifierType
type NotifierType string

const (
	NotifierTypeAWSSns  NotifierType = "aws-sns"
	NotifierTypeMqtt    NotifierType = "mqtt"
	NotifierTypeUnknown NotifierType = "unknown"
)

// AllValues returns all NotifierType values.
func (NotifierType) AllValues() []NotifierType {
	return []NotifierType{
		NotifierTypeAWSSns,
		NotifierTypeMqtt,
		NotifierTypeUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s NotifierType) MarshalText() ([]byte, error) {
	switch s {
	case NotifierTypeAWSSns:
		return []byte(s), nil
	case NotifierTypeMqtt:
		return []byte(s), nil
	case NotifierTypeUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *NotifierType) UnmarshalText(data []byte) error {
	switch NotifierType(data) {
	case NotifierTypeAWSSns:
		*s = NotifierTypeAWSSns
		return nil
	case NotifierTypeMqtt:
		*s = NotifierTypeMqtt
		return nil
	case NotifierTypeUnknown:
		*s = NotifierTypeUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// NewOptBool returns new OptBool with value set to v.
func NewOptBool(v bool) OptBool {
	return OptBool{
		Value: v,
		Set:   true,
	}
}

// OptBool is optional bool.
type OptBool struct {
	Value bool
	Set   bool
}

// IsSet returns true if OptBool was set.
func (o OptBool) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptBool) Reset() {
	var v bool
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptBool) SetTo(v bool) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptBool) Get() (v bool, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptBool) Or(d bool) bool {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptFloat64 returns new OptFloat64 with value set to v.
func NewOptFloat64(v float64) OptFloat64 {
	return OptFloat64{
		Value: v,
		Set:   true,
	}
}

// OptFloat64 is optional float64.
type OptFloat64 struct {
	Value float64
	Set   bool
}

// IsSet returns true if OptFloat64 was set.
func (o OptFloat64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptFloat64) Reset() {
	var v float64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptFloat64) SetTo(v float64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptFloat64) Get() (v float64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptFloat64) Or(d float64) float64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptNodeSize returns new OptNodeSize with value set to v.
func NewOptNodeSize(v NodeSize) OptNodeSize {
	return OptNodeSize{
		Value: v,
		Set:   true,
	}
}

// OptNodeSize is optional NodeSize.
type OptNodeSize struct {
	Value NodeSize
	Set   bool
}

// IsSet returns true if OptNodeSize was set.
func (o OptNodeSize) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptNodeSize) Reset() {
	var v NodeSize
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptNodeSize) SetTo(v NodeSize) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptNodeSize) Get() (v NodeSize, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptNodeSize) Or(d NodeSize) NodeSize {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// Ref: #/components/schemas/ProviderSummary
type ProviderSummary struct {
	// The name of the provider.
	Name string       `json:"name"`
	Type ProviderType `json:"type"`
}

// GetName returns the value of Name.
func (s *ProviderSummary) GetName() string {
	return s.Name
}

// GetType returns the value of Type.
func (s *ProviderSummary) GetType() ProviderType {
	return s.Type
}

// SetName sets the value of Name.
func (s *ProviderSummary) SetName(val string) {
	s.Name = val
}

// SetType sets the value of Type.
func (s *ProviderSummary) SetType(val ProviderType) {
	s.Type = val
}

// The type of provider.
// Ref: #/components/schemas/ProviderType
type ProviderType string

const (
	ProviderTypeAWSEcs  ProviderType = "aws-ecs"
	ProviderTypeAWSEc2  ProviderType = "aws-ec2"
	ProviderTypeLinode  ProviderType = "linode"
	ProviderTypeHetzner ProviderType = "hetzner"
	ProviderTypeUnknown ProviderType = "unknown"
)

// AllValues returns all ProviderType values.
func (ProviderType) AllValues() []ProviderType {
	return []ProviderType{
		ProviderTypeAWSEcs,
		ProviderTypeAWSEc2,
		ProviderTypeLinode,
		ProviderTypeHetzner,
		ProviderTypeUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ProviderType) MarshalText() ([]byte, error) {
	switch s {
	case ProviderTypeAWSEcs:
		return []byte(s), nil
	case ProviderTypeAWSEc2:
		return []byte(s), nil
	case ProviderTypeLinode:
		return []byte(s), nil
	case ProviderTypeHetzner:
		return []byte(s), nil
	case ProviderTypeUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ProviderType) UnmarshalText(data []byte) error {
	switch ProviderType(data) {
	case ProviderTypeAWSEcs:
		*s = ProviderTypeAWSEcs
		return nil
	case ProviderTypeAWSEc2:
		*s = ProviderTypeAWSEc2
		return nil
	case ProviderTypeLinode:
		*s = ProviderTypeLinode
		return nil
	case ProviderTypeHetzner:
		*s = ProviderTypeHetzner
		return nil
	case ProviderTypeUnknown:
		*s = ProviderTypeUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/ProvisionNodeRequestContent
type ProvisionNodeRequestContent struct {
	// The name of the provider.
	Provider string `json:"provider"`
	// A location provided by a provider.
	Location string `json:"location"`
	// .
	Tailnet   string      `json:"tailnet"`
	Ephemeral OptBool     `json:"ephemeral"`
	Size      OptNodeSize `json:"size"`
}

// GetProvider returns the value of Provider.
func (s *ProvisionNodeRequestContent) GetProvider() string {
	return s.Provider
}

// GetLocation returns the value of Location.
func (s *ProvisionNodeRequestContent) GetLocation() string {
	return s.Location
}

// GetTailnet returns the value of Tailnet.
func (s *ProvisionNodeRequestContent) GetTailnet() string {
	return s.Tailnet
}

// GetEphemeral returns the value of Ephemeral.
func (s *ProvisionNodeRequestContent) GetEphemeral() OptBool {
	return s.Ephemeral
}

// GetSize returns the value of Size.
func (s *ProvisionNodeRequestContent) GetSize() OptNodeSize {
	return s.Size
}

// SetProvider sets the value of Provider.
func (s *ProvisionNodeRequestContent) SetProvider(val string) {
	s.Provider = val
}

// SetLocation sets the value of Location.
func (s *ProvisionNodeRequestContent) SetLocation(val string) {
	s.Location = val
}

// SetTailnet sets the value of Tailnet.
func (s *ProvisionNodeRequestContent) SetTailnet(val string) {
	s.Tailnet = val
}

// SetEphemeral sets the value of Ephemeral.
func (s *ProvisionNodeRequestContent) SetEphemeral(val OptBool) {
	s.Ephemeral = val
}

// SetSize sets the value of Size.
func (s *ProvisionNodeRequestContent) SetSize(val OptNodeSize) {
	s.Size = val
}

// Ref: #/components/schemas/ProvisionNodeResponseContent
type ProvisionNodeResponseContent struct {
	// A node's identifier.
	Execution string `json:"execution"`
}

// GetExecution returns the value of Execution.
func (s *ProvisionNodeResponseContent) GetExecution() string {
	return s.Execution
}

// SetExecution sets the value of Execution.
func (s *ProvisionNodeResponseContent) SetExecution(val string) {
	s.Execution = val
}

func (*ProvisionNodeResponseContent) provisionNodeRes() {}

// The provisioning status of a node.
// Ref: #/components/schemas/ProvisioningStatus
type ProvisioningStatus string

const (
	ProvisioningStatusCreated  ProvisioningStatus = "created"
	ProvisioningStatusCreating ProvisioningStatus = "creating"
	ProvisioningStatusFailed   ProvisioningStatus = "failed"
	ProvisioningStatusUnknown  ProvisioningStatus = "unknown"
)

// AllValues returns all ProvisioningStatus values.
func (ProvisioningStatus) AllValues() []ProvisioningStatus {
	return []ProvisioningStatus{
		ProvisioningStatusCreated,
		ProvisioningStatusCreating,
		ProvisioningStatusFailed,
		ProvisioningStatusUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s ProvisioningStatus) MarshalText() ([]byte, error) {
	switch s {
	case ProvisioningStatusCreated:
		return []byte(s), nil
	case ProvisioningStatusCreating:
		return []byte(s), nil
	case ProvisioningStatusFailed:
		return []byte(s), nil
	case ProvisioningStatusUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ProvisioningStatus) UnmarshalText(data []byte) error {
	switch ProvisioningStatus(data) {
	case ProvisioningStatusCreated:
		*s = ProvisioningStatusCreated
		return nil
	case ProvisioningStatusCreating:
		*s = ProvisioningStatusCreating
		return nil
	case ProvisioningStatusFailed:
		*s = ProvisioningStatusFailed
		return nil
	case ProvisioningStatusUnknown:
		*s = ProvisioningStatusUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/ResourceNotFoundErrorResponseContent
type ResourceNotFoundErrorResponseContent struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *ResourceNotFoundErrorResponseContent) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *ResourceNotFoundErrorResponseContent) SetMessage(val string) {
	s.Message = val
}

func (*ResourceNotFoundErrorResponseContent) deprovisionNodeRes()  {}
func (*ResourceNotFoundErrorResponseContent) describeNodeRes()     {}
func (*ResourceNotFoundErrorResponseContent) describeProviderRes() {}
func (*ResourceNotFoundErrorResponseContent) describeTailnetRes()  {}
func (*ResourceNotFoundErrorResponseContent) getExecutionRes()     {}
func (*ResourceNotFoundErrorResponseContent) getNodeStatusRes()    {}
func (*ResourceNotFoundErrorResponseContent) provisionNodeRes()    {}
func (*ResourceNotFoundErrorResponseContent) startNodeRes()        {}
func (*ResourceNotFoundErrorResponseContent) stopNodeRes()         {}

type SmithyAPIHttpApiKeyAuth struct {
	APIKey string
}

// GetAPIKey returns the value of APIKey.
func (s *SmithyAPIHttpApiKeyAuth) GetAPIKey() string {
	return s.APIKey
}

// SetAPIKey sets the value of APIKey.
func (s *SmithyAPIHttpApiKeyAuth) SetAPIKey(val string) {
	s.APIKey = val
}

// Ref: #/components/schemas/StartNodeResponseContent
type StartNodeResponseContent struct {
	Success bool `json:"success"`
}

// GetSuccess returns the value of Success.
func (s *StartNodeResponseContent) GetSuccess() bool {
	return s.Success
}

// SetSuccess sets the value of Success.
func (s *StartNodeResponseContent) SetSuccess(val bool) {
	s.Success = val
}

func (*StartNodeResponseContent) startNodeRes() {}

// Ref: #/components/schemas/StopNodeResponseContent
type StopNodeResponseContent struct {
	Success bool `json:"success"`
}

// GetSuccess returns the value of Success.
func (s *StopNodeResponseContent) GetSuccess() bool {
	return s.Success
}

// SetSuccess sets the value of Success.
func (s *StopNodeResponseContent) SetSuccess(val bool) {
	s.Success = val
}

func (*StopNodeResponseContent) stopNodeRes() {}

// Summary of a tailnet.
// Ref: #/components/schemas/TailnetSummary
type TailnetSummary struct {
	// .
	Name string      `json:"name"`
	Type TailnetType `json:"type"`
	// The server address of the tailnet. This must be set for headscale tailnets.
	ControlServer string `json:"controlServer"`
}

// GetName returns the value of Name.
func (s *TailnetSummary) GetName() string {
	return s.Name
}

// GetType returns the value of Type.
func (s *TailnetSummary) GetType() TailnetType {
	return s.Type
}

// GetControlServer returns the value of ControlServer.
func (s *TailnetSummary) GetControlServer() string {
	return s.ControlServer
}

// SetName sets the value of Name.
func (s *TailnetSummary) SetName(val string) {
	s.Name = val
}

// SetType sets the value of Type.
func (s *TailnetSummary) SetType(val TailnetType) {
	s.Type = val
}

// SetControlServer sets the value of ControlServer.
func (s *TailnetSummary) SetControlServer(val string) {
	s.ControlServer = val
}

// Ref: #/components/schemas/TailnetType
type TailnetType string

const (
	TailnetTypeTailscale TailnetType = "tailscale"
	TailnetTypeHeadscale TailnetType = "headscale"
	TailnetTypeUnknown   TailnetType = "unknown"
)

// AllValues returns all TailnetType values.
func (TailnetType) AllValues() []TailnetType {
	return []TailnetType{
		TailnetTypeTailscale,
		TailnetTypeHeadscale,
		TailnetTypeUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s TailnetType) MarshalText() ([]byte, error) {
	switch s {
	case TailnetTypeTailscale:
		return []byte(s), nil
	case TailnetTypeHeadscale:
		return []byte(s), nil
	case TailnetTypeUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *TailnetType) UnmarshalText(data []byte) error {
	switch TailnetType(data) {
	case TailnetTypeTailscale:
		*s = TailnetTypeTailscale
		return nil
	case TailnetTypeHeadscale:
		*s = TailnetTypeHeadscale
		return nil
	case TailnetTypeUnknown:
		*s = TailnetTypeUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// The name of a workflow.
// Ref: #/components/schemas/WorkflowName
type WorkflowName string

const (
	WorkflowNameProvisionNode   WorkflowName = "provision-node"
	WorkflowNameDeprovisionNode WorkflowName = "deprovision-node"
	WorkflowNameUnknown         WorkflowName = "unknown"
)

// AllValues returns all WorkflowName values.
func (WorkflowName) AllValues() []WorkflowName {
	return []WorkflowName{
		WorkflowNameProvisionNode,
		WorkflowNameDeprovisionNode,
		WorkflowNameUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s WorkflowName) MarshalText() ([]byte, error) {
	switch s {
	case WorkflowNameProvisionNode:
		return []byte(s), nil
	case WorkflowNameDeprovisionNode:
		return []byte(s), nil
	case WorkflowNameUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *WorkflowName) UnmarshalText(data []byte) error {
	switch WorkflowName(data) {
	case WorkflowNameProvisionNode:
		*s = WorkflowNameProvisionNode
		return nil
	case WorkflowNameDeprovisionNode:
		*s = WorkflowNameDeprovisionNode
		return nil
	case WorkflowNameUnknown:
		*s = WorkflowNameUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}
