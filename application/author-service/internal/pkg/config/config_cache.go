// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package config

import (
	"context"
	"fmt"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"

	"author-service/internal/pkg/constant"

	redis_store "github.com/eko/gocache/store/redis/v4"
)

type (
	CacheValueType = string
	Cache[T any]   struct {
		Reader *cache.Cache[T]
		Writer *cache.Cache[T]
	}
)

type CacheRedisConfigOptions struct {
	URL string
}

type CacheConfig[T any] struct {
	Enabled bool
	Primary T
	Replica T
}

type CacheDialects struct {
	Redis CacheConfig[CacheRedisConfigOptions]
}

var (
	redisCache Cache[CacheValueType]
)

func ConfigureCache(dialect constant.DialectCache, CacheDialects CacheDialects) {
	switch dialect {
	case constant.REDIS:
		// Check if Redis is enabled
		if !CacheDialects.Redis.Enabled {
			GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("database cache %s is not enabled", dialect),
			}).Infoln("[CONFIGURE DATABASE CACHE]")
			return
		}

		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("configure database cache %s", dialect),
		}).Infoln("[CONFIGURE DATABASE CACHE]")

		// Database Cache Primary Setup
		cfgCachePrimary := CacheDialects.Redis.Primary

		urlPrimary := cfgCachePrimary.URL
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
		cachePrimary := cache.New[CacheValueType](redisStorePrimary)

		// Database Cache Replica Setup
		cfgCacheReplica := CacheDialects.Redis.Replica

		urlReplica := cfgCacheReplica.URL
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
		cacheReplica := cache.New[CacheValueType](redisStoreReplica)

		redisCache = Cache[CacheValueType]{
			Reader: cacheReplica,
			Writer: cachePrimary,
		}
	}
}

func GetCache(dialect constant.DialectCache) Cache[CacheValueType] {
	var Cache Cache[CacheValueType]

	switch dialect {
	case constant.REDIS:
		Cache = redisCache
	}

	return Cache
}
