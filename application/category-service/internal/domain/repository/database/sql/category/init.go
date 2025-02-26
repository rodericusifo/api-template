// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

import (
	"gorm.io/gorm"

	"category-service/internal/domain/model/database/sql"
	"category-service/internal/pkg/config"
	"category-service/internal/pkg/interfaces"
	"category-service/internal/pkg/types"
)

type ICategoryDatabaseSQLRepository interface {
	interfaces.IDatabaseSQLRepository
	SaveCategory(payload *sql.Category) error
	DeleteCategory(payload *sql.Category) error
	FindCategories(query *types.QuerySQL) ([]*sql.Category, error)
	FirstCategory(query *types.QuerySQL) (*sql.Category, error)
	CountCategories(query *types.QuerySQL) (int64, error)
}

type CategoryDatabaseSQLRepository struct {
	db    config.DatabaseSQL
	model sql.Category
	tx    *gorm.DB
}

func InitCategoryDatabaseSQLRepository(db config.DatabaseSQL) ICategoryDatabaseSQLRepository {
	return &CategoryDatabaseSQLRepository{
		db:    db,
		model: sql.Category{},
	}
}

func (r *CategoryDatabaseSQLRepository) BeginTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *CategoryDatabaseSQLRepository) EndTransaction() {
	r.tx = nil
}

func (r *CategoryDatabaseSQLRepository) Writer() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Writer
}

func (r *CategoryDatabaseSQLRepository) Reader() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Reader
}
