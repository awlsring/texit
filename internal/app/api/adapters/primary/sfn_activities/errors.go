package sfn_activities

import "fmt"

type DeviceNotFoundError struct {
	Message string
}

func (e *DeviceNotFoundError) Error() string {
	return fmt.Sprintf("DeviceNotFound: %s", e.Message)
}
