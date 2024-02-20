package config

import "errors"

type DatabaseEngine string

const (
	DatabaseEngineSqlite   DatabaseEngine = "sqlite"
	DatabaseEnginePostgres DatabaseEngine = "postgres"
	DatabaseEngineDynamoDb DatabaseEngine = "dynamodb"
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
	// Location of the database file. For sqlite only
	Location string `yaml:"location"`
	// Whether to use SSL for the connection
	Ssl bool `yaml:"ssl"`
}

func (c *DatabaseConfig) Validate() error {
	if c.Engine == DatabaseEngineDynamoDb {
		return nil
	}

	if c.Engine == "" {
		c.Engine = DatabaseEngineSqlite
	}

	if c.Engine == DatabaseEngineSqlite {
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
