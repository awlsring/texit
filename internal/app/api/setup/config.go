package setup

import (
	"context"
	"io"
	"os"

	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/pkg/runtime"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	configEnvVar          = "CONFIG_PATH"
	defaultConfigLocation = "/etc/texit/config.yaml"
)

func LoadConfig() (*config.Config, error) {
	if runtime.IsLambda() {
		return loadConfigFromS3()
	}
	return config.LoadFromFile(getConfigPath())
}

func getConfigPath() string {
	path := os.Getenv(configEnvVar)
	if path == "" {
		return defaultConfigLocation
	}
	return path
}

func loadConfigFromS3() (*config.Config, error) {
	acfg, err := awsconfig.LoadDefaultConfig(context.Background())
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(acfg)
	bucketName := os.Getenv("CONFIG_BUCKET")
	if bucketName == "" {
		panic("CONFIG_BUCKET environment variable not set")
	}
	resp, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    aws.String("config.yaml"),
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadFromData(bytes)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
