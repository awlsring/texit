package provider

type Identifier string

func (i Identifier) String() string {
	return string(i)
}

func ProviderFromString(s string) (Identifier, error) {
	// TODO: validate
	return Identifier(s), nil
}
