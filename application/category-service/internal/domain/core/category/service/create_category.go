// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"category-service/internal/domain/core/category/service/dto/input"
	"category-service/internal/domain/model/database/sql"
	"category-service/internal/pkg/types"
)

func (s *CategoryService) CreateCategory(payload *input.CreateCategoryDTO) error {
	categoryModelRes, err := s.CategoryDatabaseSQLRepository.FirstCategory(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "name", Operator: "ILIKE", Value: payload.Name},
			},
		},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if categoryModelRes != nil {
		return status.Error(codes.AlreadyExists, "category already exist")
	}

	categoryModel := &sql.Category{
		Name:        payload.Name,
		Description: payload.Description,
	}
	err = s.CategoryDatabaseSQLRepository.SaveCategory(categoryModel)
	if err != nil {
		return err
	}

	return nil
}
