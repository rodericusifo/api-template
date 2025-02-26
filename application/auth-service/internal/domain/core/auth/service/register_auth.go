// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"auth-service/internal/domain/core/auth/service/dto/input"
	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/types"
	"auth-service/internal/pkg/util/generator"
)

func (s *AuthService) RegisterAuth(payload *input.RegisterAuthDTO) error {
	userModelRes, err := s.UserDatabaseSQLRepository.FirstUser(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "email", Operator: "=", Value: payload.Email},
			},
		},
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if userModelRes != nil {
		return status.Error(codes.AlreadyExists, "user already exist")
	}

	roleModelRes, err := s.RoleDatabaseSQLRepository.FirstRole(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "slug", Operator: "=", Value: payload.RoleSlug},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Error(codes.NotFound, "role not found")
		}
		return err
	}

	hashedPassword, err := generator.GenerateHashFromPassword(payload.Password)
	if err != nil {
		return err
	}

	userModel := &sql.User{
		Name:     payload.Name,
		Email:    payload.Email,
		Password: hashedPassword,
		RoleID:   roleModelRes.ID,
	}
	err = s.UserDatabaseSQLRepository.SaveUser(userModel)
	if err != nil {
		return err
	}

	return nil
}
