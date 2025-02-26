// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

import (
	"book-service/internal/pkg/config"
	"book-service/internal/pkg/constant"

	internal_domain_repository_database_sql_book "book-service/internal/domain/repository/database/sql/book"
)

func BookDatabaseSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_database_sql_book.IBookDatabaseSQLRepository {
	databaseSQL := config.GetDatabaseSQL(dialect)
	iBookDatabaseSQLRepository := internal_domain_repository_database_sql_book.InitBookDatabaseSQLRepository(databaseSQL)
	return iBookDatabaseSQLRepository
}
