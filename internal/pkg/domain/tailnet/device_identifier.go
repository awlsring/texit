package tailnet

import "github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/values"

type DeviceIdentifier string

func (i DeviceIdentifier) String() string {
	return string(i)
}

func DeviceIdentifierFromString(id string) (DeviceIdentifier, error) {
	identifier, err := values.NonNullString[DeviceIdentifier](id)
	if err != nil {
		return "", err
	}
	return DeviceIdentifier(identifier), nil
}

func FormDeviceIdentifier(location string, id string) DeviceIdentifier {
	return DeviceIdentifier(location + "-" + id)
}
