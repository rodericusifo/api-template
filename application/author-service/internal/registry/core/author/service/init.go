// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"author-service/internal/pkg/constant"

	internal_domain_core_author_service "author-service/internal/domain/core/author/service"
	internal_registry_repository_database_sql_author "author-service/internal/registry/repository/database/sql/author"
)

func AuthorService() internal_domain_core_author_service.IAuthorService {
	iAuthorDatabaseSQLRepository := internal_registry_repository_database_sql_author.AuthorDatabaseSQLRepository(constant.POSTGRES)
	iAuthorService := internal_domain_core_author_service.InitAuthorService(iAuthorDatabaseSQLRepository)
	return iAuthorService
}
