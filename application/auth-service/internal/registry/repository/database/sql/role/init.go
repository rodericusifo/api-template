// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package role

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/constant"

	internal_domain_repository_database_sql_role "auth-service/internal/domain/repository/database/sql/role"
)

func RoleDatabaseSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_database_sql_role.IRoleDatabaseSQLRepository {
	databaseSQL := config.GetDatabaseSQL(dialect)
	iRoleDatabaseSQLRepository := internal_domain_repository_database_sql_role.InitRoleDatabaseSQLRepository(databaseSQL)
	return iRoleDatabaseSQLRepository
}
