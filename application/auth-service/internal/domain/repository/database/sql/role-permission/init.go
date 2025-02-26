// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rolepermission

import (
	"gorm.io/gorm"

	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/interfaces"
	"auth-service/internal/pkg/types"
)

type IRolePermissionDatabaseSQLRepository interface {
	interfaces.IDatabaseSQLRepository
	SaveRolePermission(payload *sql.RolePermission) error
	FirstRolePermission(query *types.QuerySQL) (*sql.RolePermission, error)
	FindRolePermissions(query *types.QuerySQL) ([]*sql.RolePermission, error)
}

type RolePermissionDatabaseSQLRepository struct {
	db    config.DatabaseSQL
	model sql.RolePermission
	tx    *gorm.DB
}

func InitRolePermissionDatabaseSQLRepository(db config.DatabaseSQL) IRolePermissionDatabaseSQLRepository {
	return &RolePermissionDatabaseSQLRepository{
		db:    db,
		model: sql.RolePermission{},
	}
}

func (r *RolePermissionDatabaseSQLRepository) BeginTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *RolePermissionDatabaseSQLRepository) EndTransaction() {
	r.tx = nil
}

func (r *RolePermissionDatabaseSQLRepository) Writer() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Writer
}

func (r *RolePermissionDatabaseSQLRepository) Reader() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Reader
}
