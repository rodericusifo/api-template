// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"auth-service/internal/domain/core/user/service/dto/input"
	"auth-service/internal/domain/core/user/service/dto/output"
	"auth-service/internal/domain/repository/database/sql/role"
	"auth-service/internal/domain/repository/database/sql/user"

	rolepermission "auth-service/internal/domain/repository/database/sql/role-permission"
)

type IUserService interface {
	GetUser(payload *input.GetUserDTO) (output.GetUserDTO, error)

	ValidateUserRolePermissions(payload *input.ValidateUserRolePermissionsDTO) error
}

type UserService struct {
	UserDatabaseSQLRepository           user.IUserDatabaseSQLRepository
	RoleDatabaseSQLRepository           role.IRoleDatabaseSQLRepository
	RolePermissionDatabaseSQLRepository rolepermission.IRolePermissionDatabaseSQLRepository
}

func InitUserService(
	userDatabaseSQLRepository user.IUserDatabaseSQLRepository,
	roleDatabaseSQLRepository role.IRoleDatabaseSQLRepository,
	rolePermissionDatabaseSQLRepository rolepermission.IRolePermissionDatabaseSQLRepository,
) IUserService {
	return &UserService{
		UserDatabaseSQLRepository:           userDatabaseSQLRepository,
		RoleDatabaseSQLRepository:           roleDatabaseSQLRepository,
		RolePermissionDatabaseSQLRepository: rolePermissionDatabaseSQLRepository,
	}
}
