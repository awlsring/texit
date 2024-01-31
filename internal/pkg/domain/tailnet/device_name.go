package tailnet

import "github.com/awlsring/texit/internal/pkg/values"

type DeviceName string

func (i DeviceName) String() string {
	return string(i)
}

func DeviceNameFromString(id string) (DeviceName, error) {
	identifier, err := values.NonNullString[DeviceName](id)
	if err != nil {
		return "", err
	}
	return DeviceName(identifier), nil
}

func FormDeviceName(location string, id string) DeviceName {
	return DeviceName(location + "-" + id)
}
