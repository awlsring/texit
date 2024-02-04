package handler

import (
	"fmt"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/urfave/cli/v2"
)

func (h *Handler) ProviderInit(c *cli.Context) error {
	t, err := provider.TypeFromString(c.String(flag.ProviderType))
	if err != nil {
		return err
	}

	switch t {
	case provider.TypeAwsEcs:
		return h.initAwsEcsProvider(c)
	default:
		return fmt.Errorf("unknown provider type: %s", t.String())
	}
}

func (h *Handler) initAwsEcsProvider(ctx *cli.Context) error {
	fmt.Println("To use an AWS ECS provider, you must have an AWS account and have the AWS CLI installed and configured.")
	fmt.Println("For more on this information, see: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-quickstart.html")
	fmt.Println("")
	fmt.Println("Once you have the AWS CLI installed and configured, you can run the following commands to create the required resources:")
	fmt.Println("")
	fmt.Println("First, create an IAM user that the provider will use")
	fmt.Println("aws iam create-user --user-name texit-ecs-provider")
	fmt.Println("")
	fmt.Println("Then create a JSON file with a policy document with the following command.")
	fmt.Println("")
	fmt.Println(`echo '{
	"Version": "2012-10-17",
	"Statement": [
		{
			"Effect": "Allow",
			"Action": [
				"ec2:DescribeVpcs",
				"ec2:DescribeSubnets",
				"ec2:DescribeSecurityGroups",
				"ecs:CreateCluster",
				"ecs:DeleteCluster",
				"ecs:DescribeClusters",
				"ecs:ListClusters",
				"ecs:CreateService",
				"ecs:DeleteService",
				"ecs:DescribeServices",
				"ecs:ListServices",
				"ecs:UpdateService",
				"ecs:RegisterTaskDefinition",
				"ecs:DeregisterTaskDefinition",
				"ecs:DescribeTaskDefinition",
				"ecs:ListTaskDefinitions",
				"ecs:DeleteTaskDefinitions",
				"ecs:ListTasks",
				"ecs:DescribeTasks",
				"ecs:TagResource",
				"ssm:PutParameter",
				"ssm:GetParameter",
				"ssm:DeleteParameter",
				"ssm:AddTagsToResource",
				"iam:CreateRole",
				"iam:GetRole",
				"iam:AttachRolePolicy",
				"iam:DetachRolePolicy",
				"iam:DeleteRole",
				"iam:ListAttachedRolePolicies",
				"iam:ListRolePolicies",
				"iam:CreatePolicy",
				"iam:DeletePolicy",
				"iam:TagPolicy",
				"iam:UntagPolicy",
				"iam:TagRole",
				"iam:UntagRole"
			],
			"Resource": "*"
		},
		{
            "Action": "iam:PassRole",
            "Effect": "Allow",
            "Resource": [
                "*"
            ],
            "Condition": {
                "StringLike": {
                    "iam:PassedToService": "ecs-tasks.amazonaws.com"
                }
            }
        },
        {
            "Action": "iam:PassRole",
            "Effect": "Allow",
            "Resource": [
                "arn:aws:iam::*:role/ecsInstanceRole*"
            ],
            "Condition": {
                "StringLike": {
                    "iam:PassedToService": [
                        "ec2.amazonaws.com",
                        "ec2.amazonaws.com.cn"
                    ]
                }
            }
        },
        {
            "Action": "iam:PassRole",
            "Effect": "Allow",
            "Resource": [
                "arn:aws:iam::*:role/ecsAutoscaleRole*"
            ],
            "Condition": {
                "StringLike": {
                    "iam:PassedToService": [
                        "application-autoscaling.amazonaws.com",
                        "application-autoscaling.amazonaws.com.cn"
                    ]
                }
            }
        }
	]
}' > texit-ecs-provider-policy.json`)
	fmt.Println("")
	fmt.Println("Next, create the IAM policy from the JSON file. Note the ARN of the policy as it will be used next.")
	fmt.Println("aws iam create-policy --policy-name texit-ecs-provider-policy --policy-document file://texit-ecs-provider-policy.json")
	fmt.Println("")
	fmt.Println("Attach the policy to the user")
	fmt.Println("aws iam attach-user-policy --user-name texit-ecs-provider --policy-arn <policy-arn>")
	fmt.Println("")
	fmt.Println("Create an access key for the user")
	fmt.Println("aws iam create-access-key --user-name texit-ecs-provider")
	fmt.Println("")
	fmt.Println("Finally, configure the provider with the access key and secret key in the texit configuration file")
	return nil
}
