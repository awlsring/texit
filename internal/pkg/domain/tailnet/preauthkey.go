package tailnet

import "github.com/awlsring/texit/internal/pkg/values"

type PreauthKey string

func (k PreauthKey) String() string {
	return string(k)
}

func PreauthKeyFromString(s string) (PreauthKey, error) {
	key, err := values.NonNullString[PreauthKey](s)
	if err != nil {
		return "", err
	}
	return PreauthKey(key), nil
}
