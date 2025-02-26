// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"author-service/internal/pkg/config"
	"author-service/internal/pkg/constant"

	internal_domain_repository_database_sql_author "author-service/internal/domain/repository/database/sql/author"
)

func AuthorDatabaseSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_database_sql_author.IAuthorDatabaseSQLRepository {
	databaseSQL := config.GetDatabaseSQL(dialect)
	iAuthorDatabaseSQLRepository := internal_domain_repository_database_sql_author.InitAuthorDatabaseSQLRepository(databaseSQL)
	return iAuthorDatabaseSQLRepository
}
