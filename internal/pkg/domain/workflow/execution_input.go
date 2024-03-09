package workflow

import "encoding/json"

type ExecutionInput interface {
	ToJson() (string, error)
	ExecutionIdentifier() string
}

type ProvisionNodeInput struct {
	ExecutionId          string `json:"executionId"`
	ProviderName         string `json:"providerName"`
	Location             string `json:"location"`
	TailnetName          string `json:"tailnetName"`
	TailnetControlServer string `json:"tailnetControlServer"`
	Size                 string `json:"size"`
	Ephemeral            bool   `json:"ephemeral"`
}

func (i *ProvisionNodeInput) ToJson() (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (i ProvisionNodeInput) ExecutionIdentifier() string {
	return i.ExecutionId
}

type DeprovisionNodeInput struct {
	ExecutionId     string `json:"executionId"`
	NodeId          string `json:"nodeId"`
	Tailnet         string `json:"tailnetName"`
	TailnetDeviceId string `json:"tailnetDeviceId"`
}

func (i *DeprovisionNodeInput) ToJson() (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func (i DeprovisionNodeInput) ExecutionIdentifier() string {
	return i.ExecutionId
}
