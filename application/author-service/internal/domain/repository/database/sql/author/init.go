// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

import (
	"gorm.io/gorm"

	"author-service/internal/domain/model/database/sql"
	"author-service/internal/pkg/config"
	"author-service/internal/pkg/interfaces"
	"author-service/internal/pkg/types"
)

type IAuthorDatabaseSQLRepository interface {
	interfaces.IDatabaseSQLRepository
	SaveAuthor(payload *sql.Author) error
	DeleteAuthor(payload *sql.Author) error
	FindAuthors(query *types.QuerySQL) ([]*sql.Author, error)
	FirstAuthor(query *types.QuerySQL) (*sql.Author, error)
	CountAuthors(query *types.QuerySQL) (int64, error)
}

type AuthorDatabaseSQLRepository struct {
	db    config.DatabaseSQL
	model sql.Author
	tx    *gorm.DB
}

func InitAuthorDatabaseSQLRepository(db config.DatabaseSQL) IAuthorDatabaseSQLRepository {
	return &AuthorDatabaseSQLRepository{
		db:    db,
		model: sql.Author{},
	}
}

func (r *AuthorDatabaseSQLRepository) BeginTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *AuthorDatabaseSQLRepository) EndTransaction() {
	r.tx = nil
}

func (r *AuthorDatabaseSQLRepository) Writer() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Writer
}

func (r *AuthorDatabaseSQLRepository) Reader() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Reader
}
