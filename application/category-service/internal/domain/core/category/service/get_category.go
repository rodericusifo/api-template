// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"category-service/internal/domain/core/category/service/dto/input"
	"category-service/internal/domain/core/category/service/dto/output"
	"category-service/internal/pkg/types"
	"category-service/internal/pkg/util/serializer"
)

func (s *CategoryService) GetCategory(payload *input.GetCategoryDTO) (output.GetCategoryDTO, error) {
	query := &types.QuerySQL{
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
			},
		},
	}
	if payload.ID != nil {
		query.Searches = append(query.Searches, []types.SearchQuerySQLOperation{
			{Field: "id", Operator: "=", Value: payload.ID},
		})
	}

	categoryModelRes, err := s.CategoryDatabaseSQLRepository.FirstCategory(query)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "category not found")
		}
		return nil, err
	}

	categoryDto := serializer.SerializeCategoryToCategoryDTO(categoryModelRes)

	return categoryDto, nil
}
