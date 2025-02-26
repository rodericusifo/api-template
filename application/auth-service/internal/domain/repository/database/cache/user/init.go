// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"time"

	"github.com/eko/gocache/lib/v4/cache"

	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/interfaces"
)

type IUserDatabaseCacheRepository interface {
	interfaces.IDatabaseCacheRepository[config.DatabaseCacheValueType]
	SetUser(identifier string, payload string, TTL time.Duration) error
	GetUser(identifier string) (string, error)
}

type UserDatabaseCacheRepository struct {
	db      config.DatabaseCache[config.DatabaseCacheValueType]
	keyName string
}

func InitUserDatabaseCacheRepository(db config.DatabaseCache[config.DatabaseCacheValueType]) IUserDatabaseCacheRepository {
	return &UserDatabaseCacheRepository{
		db:      db,
		keyName: "users",
	}
}

func (r *UserDatabaseCacheRepository) Writer() *cache.Cache[config.DatabaseCacheValueType] {
	return r.db.Writer
}

func (r *UserDatabaseCacheRepository) Reader() *cache.Cache[config.DatabaseCacheValueType] {
	return r.db.Reader
}
