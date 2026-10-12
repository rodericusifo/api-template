// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"auth-service/internal/domain/core/auth/service/dto/input"
	"auth-service/internal/domain/core/auth/service/dto/output"
	"auth-service/internal/domain/repository/database/sql/role"
	"auth-service/internal/domain/repository/database/sql/user"

	cache_user "auth-service/internal/domain/repository/cache/user"
)

type IAuthService interface {
	RegisterAuth(payload *input.RegisterAuthDTO) error
	LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error)
}

type AuthService struct {
	UserDatabaseSQLRepository   user.IUserDatabaseSQLRepository
	UserCacheRepository cache_user.IUserCacheRepository
	RoleDatabaseSQLRepository   role.IRoleDatabaseSQLRepository
}

func InitAuthService(
	userDatabaseSQLRepository user.IUserDatabaseSQLRepository,
	userCacheRepository cache_user.IUserCacheRepository,
	roleDatabaseSQLRepository role.IRoleDatabaseSQLRepository,
) IAuthService {
	return &AuthService{
		UserDatabaseSQLRepository:   userDatabaseSQLRepository,
		UserCacheRepository: userCacheRepository,
		RoleDatabaseSQLRepository:   roleDatabaseSQLRepository,
	}
}
