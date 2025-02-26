// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"category-service/internal/domain/core/category/service/dto/input"
	"category-service/internal/domain/core/category/service/dto/output"
	"category-service/internal/domain/repository/database/sql/category"
	"category-service/internal/pkg/types"
)

type ICategoryService interface {
	CreateCategory(payload *input.CreateCategoryDTO) error
	UpdateCategory(payload *input.UpdateCategoryDTO) error
	DeleteCategory(payload *input.DeleteCategoryDTO) error
	GetCategories(payload *input.GetCategoriesDTO) (output.GetCategoriesDTO, *types.Meta, error)
	GetCategory(payload *input.GetCategoryDTO) (output.GetCategoryDTO, error)
}

type CategoryService struct {
	CategoryDatabaseSQLRepository category.ICategoryDatabaseSQLRepository
}

func InitCategoryService(categoryDatabaseSQLRepository category.ICategoryDatabaseSQLRepository) ICategoryService {
	return &CategoryService{
		CategoryDatabaseSQLRepository: categoryDatabaseSQLRepository,
	}
}
