// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/constant"

	internal_domain_repository_cache_user "auth-service/internal/domain/repository/cache/user"
)

func UserCacheRepository(dialect constant.DialectCache) internal_domain_repository_cache_user.IUserCacheRepository {
	Cache := config.GetCache(dialect)
	iUserCacheRepository := internal_domain_repository_cache_user.InitUserCacheRepository(Cache)
	return iUserCacheRepository
}
