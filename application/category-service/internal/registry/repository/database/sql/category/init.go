// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

import (
	"category-service/internal/pkg/config"
	"category-service/internal/pkg/constant"

	internal_domain_repository_database_sql_category "category-service/internal/domain/repository/database/sql/category"
)

func CategoryDatabaseSQLRepository(dialect constant.DialectDatabaseSQL) internal_domain_repository_database_sql_category.ICategoryDatabaseSQLRepository {
	databaseSQL := config.GetDatabaseSQL(dialect)
	iCategoryDatabaseSQLRepository := internal_domain_repository_database_sql_category.InitCategoryDatabaseSQLRepository(databaseSQL)
	return iCategoryDatabaseSQLRepository
}
