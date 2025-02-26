// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package config

import (
	"context"
	"fmt"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	"category-service/internal/pkg/constant"

	redis_store "github.com/eko/gocache/store/redis/v4"
)

type (
	DatabaseCacheValueType = string
	DatabaseCache[T any]   struct {
		Reader *cache.Cache[T]
		Writer *cache.Cache[T]
	}
)

type DatabaseCacheRedisConfigOptions struct {
	URL string
}

type DatabaseCacheConfig[T any] struct {
	Enabled bool
	Primary T
	Replica T
}

type DatabaseCacheDialects struct {
	Redis DatabaseCacheConfig[DatabaseCacheRedisConfigOptions]
}

var (
	redisDatabaseCache DatabaseCache[DatabaseCacheValueType]
)

func ConfigureDatabaseCache(dialect constant.DialectDatabaseCache, databaseCacheDialects DatabaseCacheDialects) {
	switch dialect {
	case constant.REDIS:
		// Check if Redis is enabled
		if !databaseCacheDialects.Redis.Enabled {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("database cache %s is not enabled", dialect),
			}).Infoln("[CONFIGURE DATABASE CACHE]")
			return
		}

		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("configure database cache %s", dialect),
		}).Infoln("[CONFIGURE DATABASE CACHE]")

		// Database Cache Primary Setup
		cfgDatabaseCachePrimary := databaseCacheDialects.Redis.Primary

		urlPrimary := cfgDatabaseCachePrimary.URL
		optsPrimary, err := redis.ParseURL(urlPrimary)
		if err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("parse URL primary database cache %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE CACHE]")
		}

		clientPrimary := redis.NewClient(optsPrimary)
		ctxPrimary := context.Background()
		if err := clientPrimary.Ping(ctxPrimary).Err(); err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("connect to primary database cache %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE CACHE]")
		}
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("connect to primary database cache %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE CACHE]")

		redisStorePrimary := redis_store.NewRedis(clientPrimary)
		cachePrimary := cache.New[DatabaseCacheValueType](redisStorePrimary)

		// Database Cache Replica Setup
		cfgDatabaseCacheReplica := databaseCacheDialects.Redis.Replica

		urlReplica := cfgDatabaseCacheReplica.URL
		optsReplica, err := redis.ParseURL(urlReplica)
		if err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("parse URL replica database cache %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE CACHE]")
		}

		clientReplica := redis.NewClient(optsReplica)
		ctxReplica := context.Background()
		if err := clientReplica.Ping(ctxReplica).Err(); err != nil {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("connect to replica database cache %s failed", dialect),
				"detail":  err,
			}).Panic("[CONFIGURE DATABASE CACHE]")
		}
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("connect to replica database cache %s success", dialect),
		}).Infoln("[CONFIGURE DATABASE CACHE]")

		redisStoreReplica := redis_store.NewRedis(clientReplica)
		cacheReplica := cache.New[DatabaseCacheValueType](redisStoreReplica)

		redisDatabaseCache = DatabaseCache[DatabaseCacheValueType]{
			Reader: cacheReplica,
			Writer: cachePrimary,
		}
	}
}

func GetDatabaseCache(dialect constant.DialectDatabaseCache) DatabaseCache[DatabaseCacheValueType] {
	var databaseCache DatabaseCache[DatabaseCacheValueType]

	switch dialect {
	case constant.REDIS:
		databaseCache = redisDatabaseCache
	}

	return databaseCache
}
