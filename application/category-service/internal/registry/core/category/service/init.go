// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"category-service/internal/pkg/constant"

	internal_domain_core_category_service "category-service/internal/domain/core/category/service"
	internal_registry_repository_database_sql_category "category-service/internal/registry/repository/database/sql/category"
)

func CategoryService() internal_domain_core_category_service.ICategoryService {
	iCategoryDatabaseSQLRepository := internal_registry_repository_database_sql_category.CategoryDatabaseSQLRepository(constant.POSTGRES)
	iCategoryService := internal_domain_core_category_service.InitCategoryService(iCategoryDatabaseSQLRepository)
	return iCategoryService
}
