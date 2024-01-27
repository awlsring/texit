package node

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type PlatformIdentifier string

func (i PlatformIdentifier) String() string {
	return string(i)
}

func NewPlatformIdentifier(id string) PlatformIdentifier {
	return PlatformIdentifier(id)
}

type Identifier string

func (i Identifier) String() string {
	return string(i)
}

func IdentifierFromString(id string) (Identifier, error) {
	// TODO: do validation here
	return Identifier(id), nil
}

func FormNewNodeIdentifier() Identifier {
	b := make([]byte, 8)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return Identifier(string(b))
}
