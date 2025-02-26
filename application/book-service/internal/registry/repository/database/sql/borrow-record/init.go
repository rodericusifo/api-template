// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package borrowrecord

import (
	"book-service/internal/pkg/config"
	"book-service/internal/pkg/constant"

	internal_domain_repository_database_sql_borrowrecord "book-service/internal/domain/repository/database/sql/borrow-record"
)

func BorrowRecordDatabaseSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_database_sql_borrowrecord.IBorrowRecordDatabaseSQLRepository {
	databaseSQL := config.GetDatabaseSQL(dialect)
	iBorrowRecordDatabaseSQLRepository := internal_domain_repository_database_sql_borrowrecord.InitBorrowRecordDatabaseSQLRepository(databaseSQL)
	return iBorrowRecordDatabaseSQLRepository
}
