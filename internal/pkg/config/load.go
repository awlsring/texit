package config

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gopkg.in/yaml.v2"
)

type Config interface {
	Validate() error
}

// Loads the application config from a file at the specified path.
func LoadFromFile[C any](path string) (*C, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return LoadFromData[C](data)
}

// Loads the application config from an S3 bucket.
func LoadFromS3[C any](client *s3.Client, bucket, key string) (*C, error) {
	resp, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return LoadFromData[C](bytes)
}

// Loads the application config from a byte slice.
func LoadFromData[C any](data []byte) (*C, error) {
	var cfg C
	err := yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
