package database

import (
	"context"
	"time"

	"github.com/glebarez/sqlite"
	_ "github.com/lib/pq" //need to running query: postgres driver
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

type DbConfig struct {
	WriteDsn string
	ReadDsn  string
}

type dbManager struct {
	dbConfig DbConfig
}

type DbInstance struct {
	Db *gorm.DB
}

func (d *dbManager) CreateDbPostgresClient(ctx context.Context) (*DbInstance, error) {
	db, err := gorm.Open(postgres.Open(d.dbConfig.WriteDsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.Use(
		dbresolver.Register(
			dbresolver.Config{
				Sources: []gorm.Dialector{
					postgres.Open(d.dbConfig.ReadDsn),
				},
				Policy:            dbresolver.RandomPolicy{},
				TraceResolverMode: true,
			},
		).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)

	return &DbInstance{
		db,
	}, err
}

func (d *dbManager) CreateDbMysqlClient(ctx context.Context) (*DbInstance, error) {
	db, err := gorm.Open(mysql.Open(d.dbConfig.WriteDsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	db.Use(
		dbresolver.Register(
			dbresolver.Config{
				Sources: []gorm.Dialector{
					mysql.Open(d.dbConfig.ReadDsn),
				},
				Policy:            dbresolver.RandomPolicy{},
				TraceResolverMode: true,
			},
		).
			SetConnMaxIdleTime(time.Hour).
			SetConnMaxLifetime(24 * time.Hour).
			SetMaxIdleConns(100).
			SetMaxOpenConns(200),
	)

	return &DbInstance{
		db,
	}, err
}

func (d *dbManager) CreateDbSqliteClient(ctx context.Context) (*DbInstance, error) {
	db, err := gorm.Open(sqlite.Open(d.dbConfig.WriteDsn), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	return &DbInstance{
		db,
	}, err
}

func NewDbManager(dbConfig DbConfig) *dbManager {
	return &dbManager{dbConfig}
}
