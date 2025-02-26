// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package permission

import (
	"gorm.io/gorm"

	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/interfaces"
	"auth-service/internal/pkg/types"
)

type IPermissionDatabaseSQLRepository interface {
	interfaces.IDatabaseSQLRepository
	SavePermission(payload *sql.Permission) error
	FirstPermission(query *types.QuerySQL) (*sql.Permission, error)
}

type PermissionDatabaseSQLRepository struct {
	db    config.DatabaseSQL
	model sql.Permission
	tx    *gorm.DB
}

func InitPermissionDatabaseSQLRepository(db config.DatabaseSQL) IPermissionDatabaseSQLRepository {
	return &PermissionDatabaseSQLRepository{
		db:    db,
		model: sql.Permission{},
	}
}

func (r *PermissionDatabaseSQLRepository) BeginTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *PermissionDatabaseSQLRepository) EndTransaction() {
	r.tx = nil
}

func (r *PermissionDatabaseSQLRepository) Writer() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Writer
}

func (r *PermissionDatabaseSQLRepository) Reader() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Reader
}
