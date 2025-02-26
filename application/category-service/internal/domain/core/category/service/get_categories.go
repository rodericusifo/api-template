// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"category-service/internal/domain/core/category/service/dto/input"
	"category-service/internal/domain/core/category/service/dto/output"
	"category-service/internal/pkg/types"
	"category-service/internal/pkg/util/counter"
	"category-service/internal/pkg/util/definer"
	"category-service/internal/pkg/util/serializer"
)

func (s *CategoryService) GetCategories(payload *input.GetCategoriesDTO) (output.GetCategoriesDTO, *types.Meta, error) {
	page, limit := definer.DefinePaginationPageLimit(payload.Page, payload.Limit)

	categoryListModelRes, err := s.CategoryDatabaseSQLRepository.FindCategories(&types.QuerySQL{
		Offset: counter.CountPaginationOffset(page, limit),
		Limit:  limit,
	})
	if err != nil {
		return nil, nil, err
	}
	countCategoryListModelRes := len(categoryListModelRes)

	if len(categoryListModelRes) < 1 {
		return nil, nil, status.Error(codes.NotFound, "categories not found")
	}

	countCategoryAllModelRes, err := s.CategoryDatabaseSQLRepository.CountCategories(nil)
	if err != nil {
		return nil, nil, err
	}

	categoryListDto := serializer.SerializeCategoriesToCategoryDTOs(categoryListModelRes)

	meta := &types.Meta{
		CurrentPage:      int32(page),
		TotalDataPerPage: int32(countCategoryListModelRes),
		TotalData:        int32(countCategoryAllModelRes),
	}

	meta.TotalPage = int32(counter.CountPaginationTotalPage(int(meta.TotalDataPerPage), int(meta.TotalData)))

	return categoryListDto, meta, nil
}
