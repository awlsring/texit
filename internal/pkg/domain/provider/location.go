package provider

import "github.com/pkg/errors"

type Location string

func (l Location) String() string {
	return string(l)
}

var (
	ErrUnknownLocation = errors.New("unknown location")
)

func LocationFromString(location string, provider Type) (Location, error) {
	switch provider {
	case TypeAwsEcs:
		return awsLocation(location)
	default:
		return "", errors.Wrap(ErrUnknownLocation, location)
	}
}

var (
	validAwsRegions = []string{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"ap-south-1",
		"ap-northeast-2",
		"ap-southeast-1",
		"ap-southeast-2",
		"ap-northeast-1",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"sa-east-1",
	}
)

func awsLocation(loc string) (Location, error) {
	for _, r := range validAwsRegions {
		if r == loc {
			return Location(loc), nil
		}
	}
	return "", ErrUnknownLocation
}
