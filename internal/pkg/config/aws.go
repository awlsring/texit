package config

import (
	"errors"
	"os"
)

const (
	DefaultAccessKeyEnvVar = "AWS_ACCESS_KEY_ID"
	DefaultSecretKeyEnvVar = "AWS_SECRET_ACCESS_KEY"
	DefaultRegionEnvVar    = "AWS_REGION"
)

var (
	ErrMissingAccessKey = errors.New("missing access key")
	ErrMissingSecretKey = errors.New("missing secret key")
	ErrMissingRegionKey = errors.New("missing region")
)

func getCustomOrDefaultFromEnv(c, d string, e error) (string, error) {
	val := os.Getenv(c)
	if val == "" {
		v := os.Getenv(d)
		if v == "" {
			return "", e
		}
		val = v
	}
	return val, nil
}

func AwsAccessKeyFromEnv(k string) (string, error) {
	return getCustomOrDefaultFromEnv(k, DefaultAccessKeyEnvVar, ErrMissingAccessKey)
}

func SecretKeyFromEnv(k string) (string, error) {
	return getCustomOrDefaultFromEnv(k, DefaultSecretKeyEnvVar, ErrMissingSecretKey)
}

func RegionFromEnv(k string) (string, error) {
	return getCustomOrDefaultFromEnv(k, DefaultRegionEnvVar, ErrMissingRegionKey)
}
