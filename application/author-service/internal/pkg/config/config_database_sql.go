// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"author-service/internal/pkg/constant"
	"author-service/internal/pkg/migration"
)

type (
	DatabaseSQL struct {
		Reader *gorm.DB
		Writer *gorm.DB
	}
)

type DatabaseSQLMysqlConfigOptions struct {
	DSN               string
	ConnectionTimeout time.Duration
	MaxIdleConnection int
	MaxOpenConnection int
	DebugMode         bool
}

type DatabaseSQLPostgresConfigOptions struct {
	DSN               string
	ConnectionTimeout time.Duration
	MaxIdleConnection int
	MaxOpenConnection int
	DebugMode         bool
}

type DatabaseSQLConfig[T any] struct {
	Enabled bool
	Primary T
	Replica T
}

type DatabaseSQLDialects struct {
	Postgres DatabaseSQLConfig[DatabaseSQLPostgresConfigOptions]
	Mysql    DatabaseSQLConfig[DatabaseSQLMysqlConfigOptions]
}

var (
	mysqlDatabaseSQL    DatabaseSQL
	postgresDatabaseSQL DatabaseSQL
)

func ConfigureDatabaseSQL(dialect constant.DialectDatabaseSQL, databaseSQLDialects DatabaseSQLDialects) {
	switch dialect {
	case constant.MYSQL:
		// Check if MySQL is enabled
		if !databaseSQLDialects.Mysql.Enabled {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("database sql %s is not enabled", dialect),
			}).Infoln("[CONFIGURE DATABASE SQL]")
			return
		}

		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("configure database sql %s", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		// Database SQL Primary Setup
		cfgDatabaseSQLPrimary := databaseSQLDialects.Mysql.Primary
		cfgGormDatabaseSQLPrimary := &gorm.Config{
			Logger: logger.Default,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
		}

		if cfgDatabaseSQLPrimary.DebugMode {
			cfgGormDatabaseSQLPrimary.Logger = logger.Default.LogMode(logger.Info)
		}

		dsnPrimary := cfgDatabaseSQLPrimary.DSN
		dbPrimary, err := gorm.Open(mysql.Open(dsnPrimary), cfgGormDatabaseSQLPrimary)
		if err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("connect to primary database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("connect to primary database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		// Auto Migration Models
		dbPrimary.AutoMigrate(migration.AutoMigrateModelList...)

		sqlDbPrimary, err := dbPrimary.DB()
		if err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("set up primary database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("set up primary database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")
		sqlDbPrimary.SetConnMaxIdleTime(cfgDatabaseSQLPrimary.ConnectionTimeout)
		sqlDbPrimary.SetMaxIdleConns(cfgDatabaseSQLPrimary.MaxIdleConnection)
		sqlDbPrimary.SetMaxOpenConns(cfgDatabaseSQLPrimary.MaxOpenConnection)

		mysqlDatabaseSQL = DatabaseSQL{
			Reader: dbPrimary,
			Writer: dbPrimary,
		}
	case constant.POSTGRES:
		// Check if Postgres is enabled
		if !databaseSQLDialects.Postgres.Enabled {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("database sql %s is not enabled", dialect),
			}).Infoln("[CONFIGURE DATABASE SQL]")
			return
		}

		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("configure database sql %s", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		// Database SQL Primary Setup
		cfgDatabaseSQLPrimary := databaseSQLDialects.Postgres.Primary
		cfgGormDatabaseSQLPrimary := &gorm.Config{
			Logger: logger.Default,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
		}

		if cfgDatabaseSQLPrimary.DebugMode {
			cfgGormDatabaseSQLPrimary.Logger = logger.Default.LogMode(logger.Info)
		}

		dsnPrimary := cfgDatabaseSQLPrimary.DSN
		dbPrimary, err := gorm.Open(postgres.Open(dsnPrimary), cfgGormDatabaseSQLPrimary)
		if err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("connect to primary database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("connect to primary database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		// Auto Migration Models
		dbPrimary.AutoMigrate(migration.AutoMigrateModelList...)

		sqlDbPrimary, err := dbPrimary.DB()
		if err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("set up primary database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("set up primary database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")
		sqlDbPrimary.SetConnMaxIdleTime(cfgDatabaseSQLPrimary.ConnectionTimeout)
		sqlDbPrimary.SetMaxIdleConns(cfgDatabaseSQLPrimary.MaxIdleConnection)
		sqlDbPrimary.SetMaxOpenConns(cfgDatabaseSQLPrimary.MaxOpenConnection)

		// Database SQL Replica Setup
		cfgDatabaseSQLReplica := databaseSQLDialects.Postgres.Replica
		cfgGormDatabaseSQLReplica := &gorm.Config{
			Logger: logger.Default,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
		}

		if cfgDatabaseSQLReplica.DebugMode {
			cfgGormDatabaseSQLReplica.Logger = logger.Default.LogMode(logger.Info)
		}

		dsnReplica := cfgDatabaseSQLReplica.DSN
		dbReplica, err := gorm.Open(postgres.Open(dsnReplica), cfgGormDatabaseSQLReplica)
		if err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("connect to replica database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("connect to replica database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")

		sqlDbReplica, err := dbReplica.DB()
		if err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("set up replica database sql %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE SQL]")
		}
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("set up replica database sql %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE SQL]")
		sqlDbReplica.SetConnMaxIdleTime(cfgDatabaseSQLReplica.ConnectionTimeout)
		sqlDbReplica.SetMaxIdleConns(cfgDatabaseSQLReplica.MaxIdleConnection)
		sqlDbReplica.SetMaxOpenConns(cfgDatabaseSQLReplica.MaxOpenConnection)

		postgresDatabaseSQL = DatabaseSQL{
			Reader: dbReplica,
			Writer: dbPrimary,
		}
	}
}

func GetDatabaseSQL(dialect constant.DialectDatabaseSQL) DatabaseSQL {
	var databaseSQL DatabaseSQL

	switch dialect {
	case constant.MYSQL:
		databaseSQL = mysqlDatabaseSQL
	case constant.POSTGRES:
		databaseSQL = postgresDatabaseSQL
	}

	return databaseSQL
}
