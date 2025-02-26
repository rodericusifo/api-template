// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package permission

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/constant"

	internal_domain_repository_database_sql_permission "auth-service/internal/domain/repository/database/sql/permission"
)

func PermissionDatabaseSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_database_sql_permission.IPermissionDatabaseSQLRepository {
	databaseSQL := config.GetDatabaseSQL(dialect)
	iPermissionDatabaseSQLRepository := internal_domain_repository_database_sql_permission.InitPermissionDatabaseSQLRepository(databaseSQL)
	return iPermissionDatabaseSQLRepository
}
