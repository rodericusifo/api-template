// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package seeder

import (
	"github.com/sirupsen/logrus"

	"auth-service/internal/pkg/constant"
	"auth-service/internal/registry/repository/database-seeder/sql/permission"
	"auth-service/internal/registry/repository/database-seeder/sql/role"

	rolepermission "auth-service/internal/registry/repository/database-seeder/sql/role-permission"
)

func SeedRolesPermissions(dialect constant.DialectDatabaseSQL) {
	RoleDatabaseSeederSQLRepository := role.RoleDatabaseSeederSQLRepository(dialect)
	PermissionDatabaseSeederSQLRepository := permission.PermissionDatabaseSeederSQLRepository(dialect)
	RolePermissionDatabaseSeederSQLRepository := rolepermission.RolePermissionDatabaseSeederSQLRepository(dialect)

	RoleDatabaseSeederSQLRepository.SeedRoles()
	PermissionDatabaseSeederSQLRepository.SeedPermissions()
	RolePermissionDatabaseSeederSQLRepository.SeedRolePermissions()
	logrus.WithFields(logrus.Fields{
		"message": "seed permission roles success",
	}).Infoln("[SEED PERMISSIONS ROLES]")
}
