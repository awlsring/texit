package provider

type Provider struct {
	Name     Identifier
	Platform Type
	Default  bool
}
