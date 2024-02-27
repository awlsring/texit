package setup

import (
	"context"

	dynamo_execution_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/execution/dynamo"
	sql_execution_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/execution/sql"
	dynamo_node_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/node/dynamo"
	sql_node_repository "github.com/awlsring/texit/internal/app/api/adapters/secondary/repository/node/sql"
	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/db"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

func LoadRepositories(cfg *config.DatabaseConfig) (repository.Node, repository.Execution) {
	switch cfg.Engine {
	case config.DatabaseEngineSqlite:
		return LoadSqliteRepositories(cfg)
	case config.DatabaseEnginePostgres:
		return LoadPostgresRepositories(cfg)
	case config.DatabaseEngineDynamoDb:
		return LoadDynamoRepositories(cfg)
	default:
		panic("unknown database engine")
	}
}

func LoadSqliteRepositories(cfg *config.DatabaseConfig) (repository.Node, repository.Execution) {
	db, err := sqlx.Connect("sqlite", cfg.Location)
	appinit.PanicOnErr(err)
	nodeRepo := sql_node_repository.New(db)
	err = nodeRepo.Init(context.Background())
	appinit.PanicOnErr(err)
	excRepo := sql_execution_repository.New(db)
	err = excRepo.Init(context.Background())
	appinit.PanicOnErr(err)
	return nodeRepo, excRepo
}

func LoadPostgresRepositories(cfg *config.DatabaseConfig) (repository.Node, repository.Execution) {
	db, err := sqlx.Connect("postgres", db.CreatePostgresConnectionURI(cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database, cfg.Ssl))
	appinit.PanicOnErr(err)
	nodeRepo := sql_node_repository.New(db)
	err = nodeRepo.Init(context.Background())
	appinit.PanicOnErr(err)
	excRepo := sql_execution_repository.New(db)
	err = excRepo.Init(context.Background())
	appinit.PanicOnErr(err)
	return nodeRepo, excRepo
}

func LoadDynamoRepositories(cfg *config.DatabaseConfig) (repository.Node, repository.Execution) {
	aCfg, err := loadAwsConfig(cfg.AccessKey, cfg.SecretKey, cfg.Region)
	appinit.PanicOnErr(err)

	ddb := dynamodb.NewFromConfig(aCfg)
	nodeRepo := dynamo_node_repository.New("TexitNodes", ddb)
	executionRepo := dynamo_execution_repository.New("TexitExecutions", ddb)
	return nodeRepo, executionRepo
}
