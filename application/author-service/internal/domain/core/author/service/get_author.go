// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"author-service/internal/domain/core/author/service/dto/input"
	"author-service/internal/domain/core/author/service/dto/output"
	"author-service/internal/pkg/types"
	"author-service/internal/pkg/util/serializer"
)

func (s *AuthorService) GetAuthor(payload *input.GetAuthorDTO) (output.GetAuthorDTO, error) {
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

	authorModelRes, err := s.AuthorDatabaseSQLRepository.FirstAuthor(query)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "author not found")
		}
		return nil, err
	}

	authorDto := serializer.SerializeAuthorToAuthorDTO(authorModelRes)

	return authorDto, nil
}
