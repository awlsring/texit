package config

import (
	"errors"
	"os"

	"github.com/awlsring/texit/internal/pkg/config"
)

type WorkflowType string

const (
	WorkflowTypeLocal WorkflowType = "local"
	WorkflowTypeSfn   WorkflowType = "sfn"
)

const (
	ProvisionWorkflowArnEnvVar   = "PROVISION_NODE_WORKFLOW_ARN"
	DeprovisionWorkflowArnEnvVar = "DEPROVISION_NODE_WORKFLOW_ARN"
	WorkflowAwsAccessKeyEnvVar   = "SFN_AWS_ACCESS_KEY_ID"
	WorkflowAwsSecretKeyEnvVar   = "SFN_AWS_SECRET_ACCESS_KEY"
	WorkflowAwsRegionEnvVar      = "SFN_AWS_REGION"
)

// Configuration for the notifier
type WorkflowConfig struct {
	// the workflow type
	Type WorkflowType `yaml:"type"`

	// SFN Specifics

	// sfn provision workflow arn
	ProvisionWorkflowArn string `yaml:"provisionWorkflowArn"`
	// sfn deprovision workflow arn
	DeprovisionWorkflowArn string `yaml:"deprovisionWorkflowArn"`
	// Access key to use when calling sfn
	AccessKey string `yaml:"accessKey"`
	// Secret key to use when calling sfn
	SecretKey string `yaml:"secretKey"`
	// The region to use when calling sfn
	Region string `yaml:"region"`
}

func NewDefaultWorkflowConfig() *WorkflowConfig {
	return &WorkflowConfig{
		Type: WorkflowTypeLocal,
	}
}

func (c *WorkflowConfig) Validate() error {
	switch c.Type {
	case WorkflowTypeLocal:
		return nil
	case WorkflowTypeSfn:
		return c.validateSfn()
	default:
		return errors.New("unknown workflow type")
	}
}

func (c *WorkflowConfig) validateSfn() error {
	if c.ProvisionWorkflowArn == "" {
		v := os.Getenv(ProvisionWorkflowArnEnvVar)
		if v == "" {
			return errors.New("missing provision workflow arn")
		}
		c.ProvisionWorkflowArn = v
	}

	if c.DeprovisionWorkflowArn == "" {
		v := os.Getenv(DeprovisionWorkflowArnEnvVar)
		if v == "" {
			return errors.New("missing deprovision workflow arn")
		}
		c.DeprovisionWorkflowArn = v
	}

	if c.AccessKey == "" {
		val, err := config.AwsAccessKeyFromEnv(WorkflowAwsAccessKeyEnvVar)
		if err == nil {
			c.AccessKey = val
		}
	}

	if c.SecretKey == "" {
		val, err := config.SecretKeyFromEnv(WorkflowAwsSecretKeyEnvVar)
		if err == nil {
			c.SecretKey = val
		}
	}

	if c.Region == "" {
		val, err := config.RegionFromEnv(WorkflowAwsRegionEnvVar)
		if err != nil {
			return err
		}
		c.Region = val
	}

	return nil
}
