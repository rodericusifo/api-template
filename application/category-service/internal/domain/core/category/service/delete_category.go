// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"category-service/internal/domain/core/category/service/dto/input"
	"category-service/internal/pkg/config"
	"category-service/internal/pkg/constant"
	"category-service/internal/pkg/types"
)

func (s *CategoryService) DeleteCategory(payload *input.DeleteCategoryDTO) error {
	databaseSQL := config.GetDatabaseSQL(constant.POSTGRES)
	tx := databaseSQL.Writer.Begin()

	s.CategoryDatabaseSQLRepository.BeginTransaction(tx)
	defer s.CategoryDatabaseSQLRepository.EndTransaction()
	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	categoryModelRes, err := s.CategoryDatabaseSQLRepository.FirstCategory(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Error(codes.NotFound, "category not found")
		}
		return err
	}

	categoryModel := categoryModelRes

	err = s.CategoryDatabaseSQLRepository.DeleteCategory(categoryModel)
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
