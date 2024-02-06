package workflow

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecutionResult(t *testing.T) {

	t.Run("Provision Node can serialize and deserialize", func(t *testing.T) {
		failedStep := "step"
		errors := []string{"error"}
		node := "node"

		result := ProvisionNodeExecutionResult{
			Node:       &node,
			FailedStep: &failedStep,
			Errors:     errors,
		}

		serialized, err := result.Serialize()
		assert.NoError(t, err)
		fmt.Println(serialized)

		deserialized, err := DeserializeExecutionResult[ProvisionNodeExecutionResult](serialized)
		assert.NoError(t, err)

		if *deserialized.FailedStep != "step" {
			t.Errorf("unexpected failed step: %v", *deserialized.FailedStep)
		}

		if len(deserialized.Errors) != 1 || deserialized.Errors[0] != "error" {
			t.Errorf("unexpected errors: %v", deserialized.Errors)
		}

		if *deserialized.Node != "node" {
			t.Errorf("unexpected node: %v", *deserialized.Node)
		}
	})

	t.Run("Provision Node can serialize and deserialize partial", func(t *testing.T) {
		node := "node"

		result := ProvisionNodeExecutionResult{
			Node: &node,
		}

		serialized, err := result.Serialize()
		assert.NoError(t, err)

		deserialized, err := DeserializeExecutionResult[ProvisionNodeExecutionResult](serialized)
		assert.NoError(t, err)

		assert.Nil(t, deserialized.FailedStep)
		assert.Nil(t, deserialized.Errors)

		if *deserialized.Node != "node" {
			t.Errorf("unexpected node: %v", *deserialized.Node)
		}
	})

	t.Run("Provision Node can serialize and deserialize when empty", func(t *testing.T) {
		result := ProvisionNodeExecutionResult{}

		serialized, err := result.Serialize()
		assert.NoError(t, err)

		deserialized, err := DeserializeExecutionResult[ProvisionNodeExecutionResult](serialized)
		assert.NoError(t, err)

		assert.Nil(t, deserialized.FailedStep)
		assert.Nil(t, deserialized.Node)
		assert.Empty(t, deserialized.Errors)
	})

	t.Run("Deprovision Node can serialize and deserialize", func(t *testing.T) {
		failedStep := "step"
		errors := []string{"error"}
		resources := []string{"resource1", "resource2"}

		result := DeprovisionNodeExecutionResult{
			ResourcesFailedToDelete: resources,
			FailedStep:              &failedStep,
			Errors:                  errors,
		}

		serialized, err := result.Serialize()
		assert.NoError(t, err)

		deserialized, err := DeserializeExecutionResult[DeprovisionNodeExecutionResult](serialized)
		assert.NoError(t, err)

		if *deserialized.FailedStep != "step" {
			t.Errorf("unexpected failed step: %v", *deserialized.FailedStep)
		}

		if len(deserialized.Errors) != 1 || deserialized.Errors[0] != "error" {
			t.Errorf("unexpected errors: %v", deserialized.Errors)
		}

		if len(deserialized.ResourcesFailedToDelete) != 2 || deserialized.ResourcesFailedToDelete[0] != "resource1" || deserialized.ResourcesFailedToDelete[1] != "resource2" {
			t.Errorf("unexpected resources: %v", deserialized.ResourcesFailedToDelete)
		}

	})

	t.Run("Deprovision Node can serialize and deserialize partial", func(t *testing.T) {
		resources := []string{"resource1", "resource2"}

		result := DeprovisionNodeExecutionResult{
			ResourcesFailedToDelete: resources,
		}

		serialized, err := result.Serialize()
		assert.NoError(t, err)

		deserialized, err := DeserializeExecutionResult[DeprovisionNodeExecutionResult](serialized)
		assert.NoError(t, err)

		assert.Nil(t, deserialized.FailedStep)
		assert.Nil(t, deserialized.Errors)

		if len(deserialized.ResourcesFailedToDelete) != 2 || deserialized.ResourcesFailedToDelete[0] != "resource1" || deserialized.ResourcesFailedToDelete[1] != "resource2" {
			t.Errorf("unexpected resources: %v", deserialized.ResourcesFailedToDelete)
		}
	})

	t.Run("Deprovision Node can serialize and deserialize when empty", func(t *testing.T) {
		result := DeprovisionNodeExecutionResult{}

		serialized, err := result.Serialize()
		assert.NoError(t, err)

		deserialized, err := DeserializeExecutionResult[DeprovisionNodeExecutionResult](serialized)
		assert.NoError(t, err)

		assert.Nil(t, deserialized.FailedStep)
		assert.Empty(t, deserialized.Errors)
		assert.Empty(t, deserialized.ResourcesFailedToDelete)
	})
}
