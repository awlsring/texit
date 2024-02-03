package config

import "errors"

type DatabaseEngine string

const (
	DatabaseEngineSqlite3 DatabaseEngine = "sqlite3"
)

const (
	defaultSqliteDbLocation = "/var/lib/texit/texit.db"
)

var (
	ErrMissingDatabaseHost = errors.New("missing database host")
	ErrMissingDatabasePort = errors.New("missing database port")
	ErrMissingDatabaseUser = errors.New("missing database user")
	ErrMissingDatabasePass = errors.New("missing database pass")
	ErrMissingDatabaseName = errors.New("missing database name")
)

// Configuration for the database
type DatabaseConfig struct {
	// The database engine to use
	Engine DatabaseEngine `yaml:"engine"`
	// The host of the database
	Host string `yaml:"host"`
	// The port of the database
	Port int `yaml:"port"`
	// The username to connect to the database
	Username string `yaml:"username"`
	// The password to connect to the database
	Password string `yaml:"password"`
	// The name of the database
	Database string `yaml:"database"`
	// Location of the database file. For sqlite3 only
	Location string `yaml:"location"`
}

func (c *DatabaseConfig) Validate() error {
	if c.Engine == "" {
		c.Engine = DatabaseEngineSqlite3
	}

	if c.Engine == DatabaseEngineSqlite3 {
		if c.Location == "" {
			c.Location = defaultSqliteDbLocation
		}
		return nil
	}

	if c.Host == "" {
		return ErrMissingDatabaseHost
	}

	if c.Port == 0 {
		return ErrMissingDatabasePort
	}

	if c.Username == "" {
		return ErrMissingDatabaseUser
	}

	if c.Password == "" {
		return ErrMissingDatabasePass
	}

	if c.Database == "" {
		return ErrMissingDatabaseName
	}

	return nil
}
