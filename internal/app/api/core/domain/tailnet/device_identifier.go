package tailnet

type DeviceIdentifier string

func (i DeviceIdentifier) String() string {
	return string(i)
}

func DeviceIdentifierFromString(id string) DeviceIdentifier {
	return DeviceIdentifier(id)
}

func FormDeviceIdentifier(location string, id string) DeviceIdentifier {
	return DeviceIdentifier(location + "-" + id)
}
