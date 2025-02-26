// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"auth-service/internal/domain/core/user/service/dto/input"
	"auth-service/internal/domain/core/user/service/dto/output"
	"auth-service/internal/pkg/types"
	"auth-service/internal/pkg/util/serializer"
)

func (s *UserService) GetUser(payload *input.GetUserDTO) (output.GetUserDTO, error) {
	userModelRes, err := s.UserDatabaseSQLRepository.FirstUser(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "id"},
			{Field: "name"},
			{Field: "email"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "xid", Operator: "=", Value: payload.XID},
			},
		},
		InnerJoins: []types.InnerJoinQuerySQLOperation{
			{
				Relation: "Role",
				Selects: []types.SelectJoinQuerySQLOperation{
					{Field: "id"},
					{Field: "name"},
					{Field: "slug"},
					{Field: "status"},
				},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, err
	}

	userDto := serializer.SerializeUserToUserDTO(userModelRes)

	return userDto, nil
}
