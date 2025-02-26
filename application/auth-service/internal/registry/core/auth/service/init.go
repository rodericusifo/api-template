// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"auth-service/internal/pkg/constant"

	internal_domain_core_auth_service "auth-service/internal/domain/core/auth/service"
	internal_registry_repository_database_cache_user "auth-service/internal/registry/repository/database/cache/user"
	internal_registry_repository_database_sql_role "auth-service/internal/registry/repository/database/sql/role"
	internal_registry_repository_database_sql_user "auth-service/internal/registry/repository/database/sql/user"
)

func AuthService() internal_domain_core_auth_service.IAuthService {
	iUserDatabaseSQLRepository := internal_registry_repository_database_sql_user.UserDatabaseSQLRepository(constant.POSTGRES)
	iUserDatabaseCacheRepository := internal_registry_repository_database_cache_user.UserDatabaseCacheRepository(constant.REDIS)
	iRoleDatabaseSQLRepository := internal_registry_repository_database_sql_role.RoleDatabaseSQLRepository(constant.POSTGRES)
	iAuthService := internal_domain_core_auth_service.InitAuthService(iUserDatabaseSQLRepository, iUserDatabaseCacheRepository, iRoleDatabaseSQLRepository)
	return iAuthService
}
