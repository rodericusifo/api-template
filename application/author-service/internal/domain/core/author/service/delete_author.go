// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"author-service/internal/domain/core/author/service/dto/input"
	"author-service/internal/pkg/config"
	"author-service/internal/pkg/constant"
	"author-service/internal/pkg/types"
)

func (s *AuthorService) DeleteAuthor(payload *input.DeleteAuthorDTO) error {
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
			return status.Error(codes.NotFound, "author not found")
		}
		return err
	}

	authorModel := authorModelRes

	err = s.AuthorDatabaseSQLRepository.DeleteAuthor(authorModel)
	if err != nil {
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}
	committed = true

	return nil
}
