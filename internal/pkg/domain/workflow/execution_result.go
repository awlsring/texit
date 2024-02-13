package workflow

import "encoding/json"

type ExecutionResult interface {
	Serialize() (SerializedExecutionResult, error)
}

type SerializedExecutionResult string

func (r SerializedExecutionResult) IsEmpty() bool {
	return r.String() == ""
}

func (r SerializedExecutionResult) String() string {
	return string(r)
}

func DeserializeExecutionResult[T ExecutionResult](s SerializedExecutionResult) (*T, error) {
	var result T
	if s.IsEmpty() {
		return &result, nil
	}
	err := json.Unmarshal([]byte(s), &result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func serializeExecutionResult[T ExecutionResult](r T) (SerializedExecutionResult, error) {
	raw, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return SerializedExecutionResult(raw), nil
}

type DeprovisionNodeExecutionResult struct {
	ResourcesFailedToDelete []string `json:"resourcesFailedToDelete,omitempty"`
	Error                   *string  `json:"error,omitempty"`
}

func NewDeprovisionNodeExecutionResult() DeprovisionNodeExecutionResult {
	return DeprovisionNodeExecutionResult{}
}

func (r DeprovisionNodeExecutionResult) Serialize() (SerializedExecutionResult, error) {
	return serializeExecutionResult(r)
}

func (r *DeprovisionNodeExecutionResult) SetError(err string) {
	r.Error = &err
}

func (r *DeprovisionNodeExecutionResult) GetError() string {
	if r.Error == nil {
		return "unknown"
	}
	return *r.Error
}

type ProvisionNodeExecutionResult struct {
	Node  *string `json:"nodeId,omitempty"`
	Error *string `json:"error,omitempty"`
}

func NewProvisionNodeExecutionResult() ProvisionNodeExecutionResult {
	return ProvisionNodeExecutionResult{}
}

func (r ProvisionNodeExecutionResult) Serialize() (SerializedExecutionResult, error) {
	return serializeExecutionResult(r)
}

func (r *ProvisionNodeExecutionResult) SetError(err string) {
	r.Error = &err
}

func (r *ProvisionNodeExecutionResult) GetError() string {
	if r.Error == nil {
		return "unknown"
	}
	return *r.Error
}

func (r ProvisionNodeExecutionResult) GetNode() string {
	if r.Node == nil {
		return "unknown"
	}
	return *r.Node
}
