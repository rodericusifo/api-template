// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"author-service/internal/domain/core/author/service/dto/input"
	"author-service/internal/pkg/config"
	"author-service/internal/pkg/constant"
	"author-service/internal/pkg/types"
)

func (s *AuthorService) UpdateAuthor(payload *input.UpdateAuthorDTO) error {
	databaseSQL := config.GetDatabaseSQL(constant.POSTGRES)
	tx := databaseSQL.Writer.Begin()

	s.AuthorDatabaseSQLRepository.BeginTransaction(tx)
	defer s.AuthorDatabaseSQLRepository.EndTransaction()
	committed := false
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	authorModelRes, err := s.AuthorDatabaseSQLRepository.FirstAuthor(&types.QuerySQL{
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
			return status.Error(codes.NotFound, "author not found")
		}
		return err
	}

	authorModel := authorModelRes

	if payload.Name != nil {
		authorModelRes, err := s.AuthorDatabaseSQLRepository.FirstAuthor(&types.QuerySQL{
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
		if authorModelRes != nil {
			return status.Error(codes.AlreadyExists, "author name already exist")
		}
		authorModel.Name = *payload.Name
	}

	authorModel.Bio = payload.Bio

	err = s.AuthorDatabaseSQLRepository.SaveAuthor(authorModel)
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
