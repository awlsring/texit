package config

// Configuration for the server
type ServerConfig struct {
	Address string `yaml:"address"`
}

func (c *ServerConfig) Validate() error {
	if c.Address == "" {
		c.Address = ":7032"
	}

	return nil
}
