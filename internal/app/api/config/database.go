package config

import (
	"errors"

	"github.com/awlsring/texit/internal/pkg/config"
)

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

const (
	DdbAccessKeyEnvVar = "DDB_AWS_ACCESS_KEY_ID"
	DdbSecretKeyEnvVar = "DDB_AWS_SECRET_ACCESS_KEY"
	DdbRegionEnvVar    = "DDB_AWS_REGION"
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
	// The region to use for a ddb table
	Region string `yaml:"region"`
	// The access key to use for a ddb table
	AccessKey string `yaml:"accessKey"`
	// The secret key to use for a ddb table
	SecretKey string `yaml:"secretKey"`
}

func NewDefaultDatabaseConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Engine:   DatabaseEngineSqlite,
		Location: defaultSqliteDbLocation,
	}
}

func (c *DatabaseConfig) Validate() error {
	if c.Engine == DatabaseEngineDynamoDb {
		if c.Region == "" {
			val, err := config.RegionFromEnv(DdbRegionEnvVar)
			if err != nil {
				return err
			}
			c.Region = val
		}
		if c.AccessKey == "" {
			val, err := config.AwsAccessKeyFromEnv(DdbAccessKeyEnvVar)
			if err == nil {
				c.AccessKey = val
			}
		}
		if c.SecretKey == "" {
			val, err := config.AwsAccessKeyFromEnv(DdbSecretKeyEnvVar)
			if err == nil {
				c.AccessKey = val
			}
		}
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
