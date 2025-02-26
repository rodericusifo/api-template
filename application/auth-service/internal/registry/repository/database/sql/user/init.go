// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/constant"

	internal_domain_repository_database_sql_user "auth-service/internal/domain/repository/database/sql/user"
)

func UserDatabaseSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_database_sql_user.IUserDatabaseSQLRepository {
	databaseSQL := config.GetDatabaseSQL(dialect)
	iUserDatabaseSQLRepository := internal_domain_repository_database_sql_user.InitUserDatabaseSQLRepository(databaseSQL)
	return iUserDatabaseSQLRepository
}
