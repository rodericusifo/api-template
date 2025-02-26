// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package role

import (
	"auth-service/internal/pkg/constant"

	internal_domain_repository_databaseseeder_sql_role "auth-service/internal/domain/repository/database-seeder/sql/role"
	internal_registry_repository_database_sql_role "auth-service/internal/registry/repository/database/sql/role"
)

func RoleDatabaseSeederSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_databaseseeder_sql_role.IRoleDatabaseSeederSQLRepository {
	iRoleDatabaseSQLRepository := internal_registry_repository_database_sql_role.RoleDatabaseSQLRepository(dialect)
	iRoleDatabaseSeederSQLRepository := internal_domain_repository_databaseseeder_sql_role.InitRoleDatabaseSeederSQLRepository(iRoleDatabaseSQLRepository)
	return iRoleDatabaseSeederSQLRepository
}
