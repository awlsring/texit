package node

import (
	"math/rand"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/values"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type PlatformIdentifier string

func (i PlatformIdentifier) String() string {
	return string(i)
}

func PlatformIdentifierFromString(id string) (PlatformIdentifier, error) {
	identifier, err := values.NonNullString[PlatformIdentifier](id)
	if err != nil {
		return "", err
	}
	return PlatformIdentifier(identifier), nil
}

func NewPlatformIdentifier(id string) PlatformIdentifier {
	return PlatformIdentifier(id)
}

type Identifier string

func (i Identifier) String() string {
	return string(i)
}

func IdentifierFromString(id string) (Identifier, error) {
	identifier, err := values.NonNullString[Identifier](id)
	if err != nil {
		return "", err
	}
	return Identifier(identifier), nil
}

func FormNewNodeIdentifier() Identifier {
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return Identifier(string(b))
}
