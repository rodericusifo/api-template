// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package migration

import (
	"gorm.io/gorm"

	"auth-service/internal/pkg/constant"
)

func BeforeAutoMigrate(tx *gorm.DB) {
	dialect := constant.DialectDatabaseSQL(tx.Dialector.Name())
	switch dialect {
	case constant.POSTGRES:
		queryCreateTypeRoleStatusEnum := `
			CREATE TYPE role_status AS ENUM (
				'ACTIVE',
				'INACTIVE'
			);
		`
		tx.Exec(queryCreateTypeRoleStatusEnum)
		queryCreateTypePermissionStatusEnum := `
			CREATE TYPE permission_status AS ENUM (
				'ACTIVE',
				'INACTIVE'
			);
		`
		tx.Exec(queryCreateTypePermissionStatusEnum)
	}
}

func AfterAutoMigrate(tx *gorm.DB) {
	dialect := constant.DialectDatabaseSQL(tx.Dialector.Name())
	switch dialect {
	case constant.POSTGRES:
		queryAlterTableRoleColumnStatus := `
			ALTER TABLE roles
			ALTER COLUMN status TYPE role_status
			USING status::role_status;
		`
		tx.Exec(queryAlterTableRoleColumnStatus)
		queryAlterTableRoleColumnStatusDefault := `
			ALTER TABLE roles
			ALTER COLUMN status SET DEFAULT 'ACTIVE'::role_status;
		`
		tx.Exec(queryAlterTableRoleColumnStatusDefault)
		queryAlterTablePermissionColumnStatus := `
			ALTER TABLE permissions
			ALTER COLUMN status TYPE permission_status
			USING status::permission_status;
		`
		tx.Exec(queryAlterTablePermissionColumnStatus)
		queryAlterTablePermissionColumnStatusDefault := `
			ALTER TABLE permissions
			ALTER COLUMN status SET DEFAULT 'ACTIVE'::permission_status;
		`
		tx.Exec(queryAlterTablePermissionColumnStatusDefault)
	}
}
