// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"auth-service/internal/domain/core/auth/service/dto/input"
	"auth-service/internal/domain/core/auth/service/dto/output"
	"auth-service/internal/domain/repository/database/sql/role"
	"auth-service/internal/domain/repository/database/sql/user"

	cache_user "auth-service/internal/domain/repository/database/cache/user"
)

type IAuthService interface {
	RegisterAuth(payload *input.RegisterAuthDTO) error
	LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error)
}

type AuthService struct {
	UserDatabaseSQLRepository   user.IUserDatabaseSQLRepository
	UserDatabaseCacheRepository cache_user.IUserDatabaseCacheRepository
	RoleDatabaseSQLRepository   role.IRoleDatabaseSQLRepository
}

func InitAuthService(
	userDatabaseSQLRepository user.IUserDatabaseSQLRepository,
	userDatabaseCacheRepository cache_user.IUserDatabaseCacheRepository,
	roleDatabaseSQLRepository role.IRoleDatabaseSQLRepository,
) IAuthService {
	return &AuthService{
		UserDatabaseSQLRepository:   userDatabaseSQLRepository,
		UserDatabaseCacheRepository: userDatabaseCacheRepository,
		RoleDatabaseSQLRepository:   roleDatabaseSQLRepository,
	}
}
