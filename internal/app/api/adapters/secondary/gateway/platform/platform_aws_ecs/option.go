package platform_aws_ecs

import "github.com/aws/aws-sdk-go-v2/credentials"

type PlatformAwsEcsGatewayOptions func(*PlatformAwsEcsGateway)

func WithKeyPair(accessKey, secretKey string) PlatformAwsEcsGatewayOptions {
	return func(g *PlatformAwsEcsGateway) {
		credProvider := credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")
		g.creds = &credProvider
	}
}
