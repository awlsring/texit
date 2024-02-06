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

func getFailureStep(failedStep *string) string {
	if failedStep == nil {
		return "unknown"
	}
	return *failedStep
}

type DeprovisionNodeExecutionResult struct {
	ResourcesFailedToDelete []string `json:"resourcesFailedToDelete,omitempty"`
	FailedStep              *string  `json:"failedStep,omitempty"`
	Errors                  []string `json:"errors,omitempty"`
}

func NewDeprovisionNodeExecutionResult(step string) DeprovisionNodeExecutionResult {
	return DeprovisionNodeExecutionResult{
		FailedStep: &step,
		Errors:     []string{},
	}
}

func (r DeprovisionNodeExecutionResult) Serialize() (SerializedExecutionResult, error) {
	return serializeExecutionResult(r)
}

func (r DeprovisionNodeExecutionResult) GetFailedStep() string {
	return getFailureStep(r.FailedStep)
}

type ProvisionNodeExecutionResult struct {
	Node       *string  `json:"nodeId,omitempty"`
	FailedStep *string  `json:"failedStep,omitempty"`
	Errors     []string `json:"errors,omitempty"`
}

func NewProvisionNodeExecutionResult(step string) ProvisionNodeExecutionResult {
	return ProvisionNodeExecutionResult{
		FailedStep: &step,
		Errors:     []string{},
	}
}

func (r ProvisionNodeExecutionResult) Serialize() (SerializedExecutionResult, error) {
	return serializeExecutionResult(r)
}

func (r ProvisionNodeExecutionResult) GetFailedStep() string {
	return getFailureStep(r.FailedStep)
}

func (r ProvisionNodeExecutionResult) GetNode() string {
	if r.Node == nil {
		return "unknown"
	}
	return *r.Node
}
