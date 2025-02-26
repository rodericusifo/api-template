// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"book-service/internal/pkg/constant"

	internal_domain_core_book_service "book-service/internal/domain/core/book/service"
	internal_registry_repository_database_sql_book "book-service/internal/registry/repository/database/sql/book"
	internal_registry_repository_database_sql_borrowrecord "book-service/internal/registry/repository/database/sql/borrow-record"
)

func BookService() internal_domain_core_book_service.IBookService {
	iBookDatabaseSQLRepository := internal_registry_repository_database_sql_book.BookDatabaseSQLRepository(constant.POSTGRES)
	iBorrowRecordDatabaseSQLRepository := internal_registry_repository_database_sql_borrowrecord.BorrowRecordDatabaseSQLRepository(constant.POSTGRES)
	iBookService := internal_domain_core_book_service.InitBookService(iBookDatabaseSQLRepository, iBorrowRecordDatabaseSQLRepository)
	return iBookService
}
