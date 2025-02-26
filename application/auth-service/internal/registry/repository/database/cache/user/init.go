// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/constant"

	internal_domain_repository_database_cache_user "auth-service/internal/domain/repository/database/cache/user"
)

func UserDatabaseCacheRepository(dialect constant.DialectDatabaseCache) internal_domain_repository_database_cache_user.IUserDatabaseCacheRepository {
	databaseCache := config.GetDatabaseCache(dialect)
	iUserDatabaseCacheRepository := internal_domain_repository_database_cache_user.InitUserDatabaseCacheRepository(databaseCache)
	return iUserDatabaseCacheRepository
}
