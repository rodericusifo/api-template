// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package role

import (
	"gorm.io/gorm"

	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/interfaces"
	"auth-service/internal/pkg/types"
)

type IRoleDatabaseSQLRepository interface {
	interfaces.IDatabaseSQLRepository
	SaveRole(payload *sql.Role) error
	FirstRole(query *types.QuerySQL) (*sql.Role, error)
}

type RoleDatabaseSQLRepository struct {
	db    config.DatabaseSQL
	model sql.Role
	tx    *gorm.DB
}

func InitRoleDatabaseSQLRepository(db config.DatabaseSQL) IRoleDatabaseSQLRepository {
	return &RoleDatabaseSQLRepository{
		db:    db,
		model: sql.Role{},
	}
}

func (r *RoleDatabaseSQLRepository) BeginTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *RoleDatabaseSQLRepository) EndTransaction() {
	r.tx = nil
}

func (r *RoleDatabaseSQLRepository) Writer() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Writer
}

func (r *RoleDatabaseSQLRepository) Reader() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Reader
}
