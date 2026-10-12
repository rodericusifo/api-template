// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"time"

	"github.com/eko/gocache/lib/v4/cache"

	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/interfaces"
)

type IUserCacheRepository interface {
	interfaces.ICacheRepository[config.CacheValueType]
	SetUser(identifier string, payload string, TTL time.Duration) error
	GetUser(identifier string) (string, error)
}

type UserCacheRepository struct {
	db      config.Cache[config.CacheValueType]
	keyName string
}

func InitUserCacheRepository(db config.Cache[config.CacheValueType]) IUserCacheRepository {
	return &UserCacheRepository{
		db:      db,
		keyName: "users",
	}
}

func (r *UserCacheRepository) Writer() *cache.Cache[config.CacheValueType] {
	return r.db.Writer
}

func (r *UserCacheRepository) Reader() *cache.Cache[config.CacheValueType] {
	return r.db.Reader
}
