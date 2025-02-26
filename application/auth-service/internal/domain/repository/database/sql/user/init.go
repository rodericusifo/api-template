// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"gorm.io/gorm"

	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/interfaces"
	"auth-service/internal/pkg/types"
)

type IUserDatabaseSQLRepository interface {
	interfaces.IDatabaseSQLRepository
	SaveUser(payload *sql.User) error
	FirstUser(query *types.QuerySQL) (*sql.User, error)
}

type UserDatabaseSQLRepository struct {
	db    config.DatabaseSQL
	model sql.User
	tx    *gorm.DB
}

func InitUserDatabaseSQLRepository(db config.DatabaseSQL) IUserDatabaseSQLRepository {
	return &UserDatabaseSQLRepository{
		db:    db,
		model: sql.User{},
	}
}

func (r *UserDatabaseSQLRepository) BeginTransaction(tx *gorm.DB) {
	r.tx = tx
}

func (r *UserDatabaseSQLRepository) EndTransaction() {
	r.tx = nil
}

func (r *UserDatabaseSQLRepository) Writer() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Writer
}

func (r *UserDatabaseSQLRepository) Reader() *gorm.DB {
	if r.tx != nil {
		return r.tx
	}
	return r.db.Reader
}
