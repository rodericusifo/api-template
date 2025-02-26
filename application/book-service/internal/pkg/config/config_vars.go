// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"book-service/internal/pkg/constant"
)

type varsConfig *struct {
	// SERVER
	ServerPort                int    `mapstructure:"SERVER_PORT"`
	AuthorServiceServerHost   string `mapstructure:"AUTHOR_SERVICE_SERVER_HOST"`
	AuthorServiceServerPort   int    `mapstructure:"AUTHOR_SERVICE_SERVER_PORT"`
	CategoryServiceServerHost string `mapstructure:"CATEGORY_SERVICE_SERVER_HOST"`
	CategoryServiceServerPort int    `mapstructure:"CATEGORY_SERVICE_SERVER_PORT"`
	BookServiceServerHost     string `mapstructure:"BOOK_SERVICE_SERVER_HOST"`
	BookServiceServerPort     int    `mapstructure:"BOOK_SERVICE_SERVER_PORT"`

	// DATABASE SQL
	DatabaseSQLMysqlEnabled                  bool          `mapstructure:"DATABASE_SQL_MYSQL_ENABLED"`
	DatabaseSQLMysqlPrimaryDSN               string        `mapstructure:"DATABASE_SQL_MYSQL_PRIMARY_DSN"`
	DatabaseSQLMysqlPrimaryConnectionTimeout time.Duration `mapstructure:"DATABASE_SQL_MYSQL_PRIMARY_CONNECTION_TIMEOUT"`
	DatabaseSQLMysqlPrimaryMaxIdleConnection int           `mapstructure:"DATABASE_SQL_MYSQL_PRIMARY_MAX_IDLE_CONNECTION"`
	DatabaseSQLMysqlPrimaryMaxOpenConnection int           `mapstructure:"DATABASE_SQL_MYSQL_PRIMARY_MAX_OPEN_CONNECTION"`
	DatabaseSQLMysqlPrimaryDebugMode         bool          `mapstructure:"DATABASE_SQL_MYSQL_PRIMARY_DEBUG_MODE"`
	DatabaseSQLMysqlReplicaDSN               string        `mapstructure:"DATABASE_SQL_MYSQL_REPLICA_DSN"`
	DatabaseSQLMysqlReplicaConnectionTimeout time.Duration `mapstructure:"DATABASE_SQL_MYSQL_REPLICA_CONNECTION_TIMEOUT"`
	DatabaseSQLMysqlReplicaMaxIdleConnection int           `mapstructure:"DATABASE_SQL_MYSQL_REPLICA_MAX_IDLE_CONNECTION"`
	DatabaseSQLMysqlReplicaMaxOpenConnection int           `mapstructure:"DATABASE_SQL_MYSQL_REPLICA_MAX_OPEN_CONNECTION"`
	DatabaseSQLMysqlReplicaDebugMode         bool          `mapstructure:"DATABASE_SQL_MYSQL_REPLICA_DEBUG_MODE"`

	DatabaseSQLPostgresEnabled                  bool          `mapstructure:"DATABASE_SQL_POSTGRES_ENABLED"`
	DatabaseSQLPostgresPrimaryDSN               string        `mapstructure:"DATABASE_SQL_POSTGRES_PRIMARY_DSN"`
	DatabaseSQLPostgresPrimaryConnectionTimeout time.Duration `mapstructure:"DATABASE_SQL_POSTGRES_PRIMARY_CONNECTION_TIMEOUT"`
	DatabaseSQLPostgresPrimaryMaxIdleConnection int           `mapstructure:"DATABASE_SQL_POSTGRES_PRIMARY_MAX_IDLE_CONNECTION"`
	DatabaseSQLPostgresPrimaryMaxOpenConnection int           `mapstructure:"DATABASE_SQL_POSTGRES_PRIMARY_MAX_OPEN_CONNECTION"`
	DatabaseSQLPostgresPrimaryDebugMode         bool          `mapstructure:"DATABASE_SQL_POSTGRES_PRIMARY_DEBUG_MODE"`
	DatabaseSQLPostgresReplicaDSN               string        `mapstructure:"DATABASE_SQL_POSTGRES_REPLICA_DSN"`
	DatabaseSQLPostgresReplicaConnectionTimeout time.Duration `mapstructure:"DATABASE_SQL_POSTGRES_REPLICA_CONNECTION_TIMEOUT"`
	DatabaseSQLPostgresReplicaMaxIdleConnection int           `mapstructure:"DATABASE_SQL_POSTGRES_REPLICA_MAX_IDLE_CONNECTION"`
	DatabaseSQLPostgresReplicaMaxOpenConnection int           `mapstructure:"DATABASE_SQL_POSTGRES_REPLICA_MAX_OPEN_CONNECTION"`
	DatabaseSQLPostgresReplicaDebugMode         bool          `mapstructure:"DATABASE_SQL_POSTGRES_REPLICA_DEBUG_MODE"`

	// DATABASE CACHE
	DatabaseCacheRedisEnabled    bool   `mapstructure:"DATABASE_CACHE_REDIS_ENABLED"`
	DatabaseCacheRedisPrimaryURL string `mapstructure:"DATABASE_CACHE_REDIS_PRIMARY_URL"`
	DatabaseCacheRedisReplicaURL string `mapstructure:"DATABASE_CACHE_REDIS_REPLICA_URL"`

	// PASSWORD HASHING
	PasswordHashingHashSalt int `mapstructure:"PASSWORD_HASHING_HASH_SALT"`

	// JWT
	JWTSecretKey       string        `mapstructure:"JWT_SECRET_KEY"`
	JWTExpiredDuration time.Duration `mapstructure:"JWT_EXPIRED_DURATION"`
}

var (
	varsCfg varsConfig
)

func ConfigureVars() {
	switch mode := os.Getenv("MODE"); {
	case strings.Contains(string(constant.DEVELOPMENT), strings.ToLower(mode)):
		setEnv()
	default:
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("MODE %s is not valid", mode),
		}).Panic("[MAIN]")
	}
}

func GetVarsConfig() varsConfig {
	return varsCfg
}

func setEnv() {
	env := os.Getenv("ENV")

	if env == "" {
		GetLogConfig().WithFields(logrus.Fields{
			"message": "ENV is not set",
		}).Panic("[CONFIGURE VARS]")
	}

	path := fmt.Sprintf("/env/%s.env", env)
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		GetLogConfig().WithFields(logrus.Fields{
			"message": "read env fail",
			"detail":  err,
		}).Panic("[CONFIGURE VARS]")
	}

	if err := viper.Unmarshal(&varsCfg); err != nil {
		GetLogConfig().WithFields(logrus.Fields{
			"message": "load env fail",
			"detail":  err,
		}).Panic("[CONFIGURE VARS]")
	}
	GetLogConfig().WithFields(logrus.Fields{
		"message": "load env success",
	}).Infoln("[CONFIGURE VARS]")
}
