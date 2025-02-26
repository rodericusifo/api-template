// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"category-service/internal/domain/core/category/service/dto/input"
	"category-service/internal/pkg/config"
	"category-service/internal/pkg/constant"
	"category-service/internal/pkg/types"
)

func (s *CategoryService) UpdateCategory(payload *input.UpdateCategoryDTO) error {
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
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
			},
		},
		Clauses: []types.ClauseExpression{
			clause.Locking{Strength: "UPDATE"},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Error(codes.NotFound, "category not found")
		}
		return err
	}

	categoryModel := categoryModelRes

	if payload.Name != nil {
		categoryModelRes, err := s.CategoryDatabaseSQLRepository.FirstCategory(&types.QuerySQL{
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "name", Operator: "ILIKE", Value: payload.Name},
					{Field: "xid", Operator: "!=", Value: payload.XID},
				},
			},
		})
		if err != nil && err != gorm.ErrRecordNotFound {
			return err
		}
		if categoryModelRes != nil {
			return status.Error(codes.AlreadyExists, "category name already exist")
		}
		categoryModel.Name = *payload.Name
	}

	categoryModel.Description = payload.Description

	err = s.CategoryDatabaseSQLRepository.SaveCategory(categoryModel)
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
