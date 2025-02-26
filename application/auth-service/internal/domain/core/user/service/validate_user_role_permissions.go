// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"auth-service/internal/domain/core/user/service/dto/input"
	"auth-service/internal/pkg/constant"
	"auth-service/internal/pkg/types"
)

func (s *UserService) ValidateUserRolePermissions(payload *input.ValidateUserRolePermissionsDTO) error {
	roleModelRes, err := s.RoleDatabaseSQLRepository.FirstRole(&types.QuerySQL{
		Selects: []types.SelectQuerySQLOperation{
			{Field: "status"},
		},
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "id", Operator: "=", Value: payload.RoleID},
			},
		},
	})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return status.Error(codes.NotFound, "role not found")
		}
		return err
	}

	if roleModelRes.Status == constant.ROLE_INACTIVE {
		return status.Error(codes.PermissionDenied, "role is inactive")
	}

	findRolePermissionsModelRes, err := s.RolePermissionDatabaseSQLRepository.FindRolePermissions(&types.QuerySQL{
		Searches: [][]types.SearchQuerySQLOperation{
			{
				{Field: "role_id", Operator: "=", Value: payload.RoleID},
			},
		},
		InnerJoins: []types.InnerJoinQuerySQLOperation{
			{
				Relation: "Permission",
				Selects: []types.SelectJoinQuerySQLOperation{
					{Field: "path"},
					{Field: "status"},
				},
			},
		},
	})
	if err != nil {
		return err
	}

	if len(findRolePermissionsModelRes) < 1 {
		return status.Error(codes.NotFound, "role permissions not found")
	}

	found := false
	active := false
	for _, rolePermission := range findRolePermissionsModelRes {
		permissionPathSegments := strings.Split(rolePermission.Permission.Path, "/")
		pathSegments := strings.Split(payload.Path, "/")

		if len(pathSegments) != len(permissionPathSegments) {
			continue
		}

		matches := make([]bool, 0, len(pathSegments))
		for i := range pathSegments {
			if strings.HasPrefix(permissionPathSegments[i], ":") {
				matches = append(matches, true)
				continue
			}
			if permissionPathSegments[i] == pathSegments[i] {
				matches = append(matches, true)
				continue
			}
			matches = append(matches, false)
		}
		found = func() bool {
			for _, v := range matches {
				if !v {
					return false
				}
			}
			return true
		}()
		if found {
			active = rolePermission.Permission.Status == constant.PERMISSION_ACTIVE
			break
		}
	}

	if !found {
		return status.Error(codes.NotFound, "role permission not found")
	}
	if !active {
		return status.Error(codes.PermissionDenied, "permission is inactive")
	}

	return nil
}
