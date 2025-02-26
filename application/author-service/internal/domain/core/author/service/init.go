// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"author-service/internal/domain/core/author/service/dto/input"
	"author-service/internal/domain/core/author/service/dto/output"
	"author-service/internal/domain/repository/database/sql/author"
	"author-service/internal/pkg/types"
)

type IAuthorService interface {
	CreateAuthor(payload *input.CreateAuthorDTO) error
	UpdateAuthor(payload *input.UpdateAuthorDTO) error
	DeleteAuthor(payload *input.DeleteAuthorDTO) error
	GetAuthors(payload *input.GetAuthorsDTO) (output.GetAuthorsDTO, *types.Meta, error)
	GetAuthor(payload *input.GetAuthorDTO) (output.GetAuthorDTO, error)
}

type AuthorService struct {
	AuthorDatabaseSQLRepository author.IAuthorDatabaseSQLRepository
}

func InitAuthorService(authorDatabaseSQLRepository author.IAuthorDatabaseSQLRepository) IAuthorService {
	return &AuthorService{
		AuthorDatabaseSQLRepository: authorDatabaseSQLRepository,
	}
}
