// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"auth-service/internal/pkg/constant"

	internal_domain_core_user_service "auth-service/internal/domain/core/user/service"
	internal_registry_repository_database_sql_role "auth-service/internal/registry/repository/database/sql/role"
	internal_registry_repository_database_sql_rolepermission "auth-service/internal/registry/repository/database/sql/role-permission"
	internal_registry_repository_database_sql_user "auth-service/internal/registry/repository/database/sql/user"
)

func UserService() internal_domain_core_user_service.IUserService {
	iUserDatabaseSQLRepository := internal_registry_repository_database_sql_user.UserDatabaseSQLRepository(constant.POSTGRES)
	iRoleDatabaseSQLRepository := internal_registry_repository_database_sql_role.RoleDatabaseSQLRepository(constant.POSTGRES)
	iRolePermissionDatabaseSQLRepository := internal_registry_repository_database_sql_rolepermission.RolePermissionDatabaseSQLRepository(constant.POSTGRES)
	iUserService := internal_domain_core_user_service.InitUserService(iUserDatabaseSQLRepository, iRoleDatabaseSQLRepository, iRolePermissionDatabaseSQLRepository)
	return iUserService
}
