package node

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
