// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"author-service/internal/domain/core/author/service/dto/input"
	"author-service/internal/domain/model/database/sql"
	"author-service/internal/pkg/types"
)

func (s *AuthorService) CreateAuthor(payload *input.CreateAuthorDTO) error {
	authorModelRes, err := s.AuthorDatabaseSQLRepository.FirstAuthor(&types.QuerySQL{
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
	if authorModelRes != nil {
		return status.Error(codes.AlreadyExists, "author already exist")
	}

	authorModel := &sql.Author{
		Name: payload.Name,
		Bio:  payload.Bio,
	}
	err = s.AuthorDatabaseSQLRepository.SaveAuthor(authorModel)
	if err != nil {
		return err
	}

	return nil
}
