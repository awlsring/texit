package platform_aws_ecs

import (
	"context"
	"errors"
	"fmt"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	taskRoleNamePrefix          = "tailscale-cloud-exit-nodes-task-role"
	executionRoleName           = "tailscale-cloud-exit-nodes-execution-role"
	manageExecutionTaskPolicy   = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
	ecsAssumeRolePolicyDocument = `{
		"Version": "2012-10-17",
		"Statement": [
			{
				"Effect": "Allow",
				"Principal": {
					"Service": "ecs-tasks.amazonaws.com"
				},
				"Action": "sts:AssumeRole"
			}
		]
	}`
)

// Makes an execution role for the task definition.
// This is idempotent, and will not error if the role already exists.
func makeExecutionRole(ctx context.Context, client interfaces.IamClient) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Making execution role")

	log.Debug().Msgf("Checking if role %s exists", executionRoleName)
	roleResp, err := client.GetRole(ctx, &iam.GetRoleInput{
		RoleName: aws.String(executionRoleName),
	})
	if err != nil {
		var notFound *types.NoSuchEntityException
		if !errors.As(err, &notFound) {
			log.Error().Err(err).Msg("Failed to get role")
			return "", err
		}
	} else {
		log.Debug().Msg("Role already exists")
		return *roleResp.Role.Arn, nil
	}

	log.Debug().Msgf("Creating role %s", executionRoleName)
	resp, err := client.CreateRole(ctx, &iam.CreateRoleInput{
		RoleName:                 aws.String(executionRoleName),
		AssumeRolePolicyDocument: aws.String(ecsAssumeRolePolicyDocument),
		Tags: []types.Tag{
			{
				Key:   aws.String("created-by"),
				Value: aws.String("tailscale-cloud-exit-nodes"),
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to make cluster role")
		return "", err
	}
	log.Debug().Msg("Execution role made")

	log.Debug().Msgf("Attaching managed execution policy to role %s", executionRoleName)
	_, err = client.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{
		RoleName:  aws.String(executionRoleName),
		PolicyArn: aws.String(manageExecutionTaskPolicy),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to attach managed execution policy to role")
		return "", err
	}

	log.Debug().Msg("Execution role made")
	return *resp.Role.Arn, nil
}

func makeTaskRoleName(tid tailnet.DeviceName) string {
	return taskRoleNamePrefix + "-" + tid.String()
}

func makeTaskPolicy(parameter string) string {
	return fmt.Sprintf(`{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "ssm:GetParameter",
        "ssm:PutParameter"
      ],
      "Resource": [
        "%s"
      ]
    }
  ]
}`, parameter)
}

func deleteTaskRole(ctx context.Context, client interfaces.IamClient, tid tailnet.DeviceName) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting task role")

	log.Debug().Msg("Getting attached role policies")
	resp, err := client.ListAttachedRolePolicies(ctx, &iam.ListAttachedRolePoliciesInput{
		RoleName: aws.String(makeTaskRoleName(tid)),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to list attached role policies")
		return err
	}

	log.Debug().Msg("Detaching role policies")
	for _, policy := range resp.AttachedPolicies {
		log.Debug().Msgf("Detaching policy %s", *policy.PolicyArn)
		_, err := client.DetachRolePolicy(ctx, &iam.DetachRolePolicyInput{
			RoleName:  aws.String(makeTaskRoleName(tid)),
			PolicyArn: policy.PolicyArn,
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to detach role policy")
			return err
		}
		log.Debug().Msg("Deleting the policy")
		_, err = client.DeletePolicy(ctx, &iam.DeletePolicyInput{
			PolicyArn: policy.PolicyArn,
		})
		if err != nil {
			log.Error().Err(err).Msg("Failed to delete policy")
			return err
		}
		log.Debug().Msg("Policy deleted")
	}

	roleName := makeTaskRoleName(tid)

	log.Debug().Msgf("Deleting role %s", roleName)
	_, err = client.DeleteRole(ctx, &iam.DeleteRoleInput{
		RoleName: aws.String(roleName),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete role")
		return err
	}

	log.Debug().Msg("Task role deleted")
	return nil
}

func makeTaskRole(ctx context.Context, client interfaces.IamClient, tid tailnet.DeviceName, parameter string) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Making task role")

	roleName := makeTaskRoleName(tid)

	log.Debug().Msgf("Creating role %s", roleName)
	resp, err := client.CreateRole(ctx, &iam.CreateRoleInput{
		RoleName:                 aws.String(roleName),
		AssumeRolePolicyDocument: aws.String(ecsAssumeRolePolicyDocument),
		Tags: []types.Tag{
			{
				Key:   aws.String("created-by"),
				Value: aws.String("tailscale-cloud-exit-nodes"),
			},
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to make task role")
		return "", err
	}
	log.Debug().Msg("Task role made")

	log.Debug().Msgf("Creating task policy")
	presp, err := client.CreatePolicy(ctx, &iam.CreatePolicyInput{
		PolicyName:     aws.String(roleName),
		PolicyDocument: aws.String(makeTaskPolicy(parameter)),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to make cluster role")
		return "", err
	}
	log.Debug().Msg("Task policy made")

	log.Debug().Msgf("Attaching task policy to role %s", roleName)
	_, err = client.AttachRolePolicy(ctx, &iam.AttachRolePolicyInput{
		RoleName:  aws.String(roleName),
		PolicyArn: presp.Policy.Arn,
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to attach policy to role")
		return "", err
	}

	log.Debug().Msg("Task role made")
	return *resp.Role.Arn, nil
}
