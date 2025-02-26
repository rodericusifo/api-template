// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rolepermission

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/types"
)

func (r *RolePermissionDatabaseSeederSQLRepository) SeedRolePermissions() {
	RolePermissions.Validate()

	for _, RolePermission := range RolePermissions {
		firstRoleModelRes, err := r.RoleDatabaseSQLRepository.FirstRole(&types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "slug", Operator: "=", Value: RolePermission.RoleSlug},
				},
			},
		})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				config.GetLogConfig().WithFields(logrus.Fields{
					"message": fmt.Sprintf("role with slug %s not found", RolePermission.RoleSlug),
					"detail":  err,
				}).Errorln("[SEED ROLE PERMISSIONS]")
				continue
			}
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": "get role fail",
				"detail":  err,
			}).Errorln("[SEED ROLE PERMISSIONS]")
			continue
		}

		if len(RolePermission.PermissionsSlug) > 0 {
			for _, permissionSlug := range RolePermission.PermissionsSlug {
				firstPermissionModelRes, err := r.PermissionDatabaseSQLRepository.FirstPermission(&types.QuerySQL{
					Selects: []types.SelectQuerySQLOperation{
						{Field: "id"},
					},
					Searches: [][]types.SearchQuerySQLOperation{
						{
							{Field: "slug", Operator: "=", Value: permissionSlug},
						},
					},
				})
				if err != nil {
					if err == gorm.ErrRecordNotFound {
						config.GetLogConfig().WithFields(logrus.Fields{
							"message": fmt.Sprintf("permission with slug %s not found", permissionSlug),
							"detail":  err,
						}).Errorln("[SEED ROLE PERMISSIONS]")
						continue
					}
					config.GetLogConfig().WithFields(logrus.Fields{
						"message": "get permission fail",
						"detail":  err,
					}).Errorln("[SEED ROLE PERMISSIONS]")
					continue
				}

				firstRolePermissionModelRes, err := r.RolePermissionDatabaseSQLRepository.FirstRolePermission(&types.QuerySQL{
					Selects: []types.SelectQuerySQLOperation{
						{Field: "id"},
					},
					Searches: [][]types.SearchQuerySQLOperation{
						{
							{Field: "role_id", Operator: "=", Value: firstRoleModelRes.ID},
							{Field: "permission_id", Operator: "=", Value: firstPermissionModelRes.ID},
						},
					},
				})
				if err != nil && err != gorm.ErrRecordNotFound {
					config.GetLogConfig().WithFields(logrus.Fields{
						"message": "get role permission fail",
						"detail":  err,
					}).Errorln("[SEED ROLE PERMISSIONS]")
					continue
				}
				if firstRolePermissionModelRes != nil {
					logrus.WithFields(logrus.Fields{
						"message": fmt.Sprintf("role permission with role_id %d and permission_id %d already registered", firstRoleModelRes.ID, firstPermissionModelRes.ID),
					}).Errorln("[SEED ROLE PERMISSIONS]")
					continue
				}

				rolePermissionSave := &sql.RolePermission{
					RoleID:       firstRoleModelRes.ID,
					PermissionID: firstPermissionModelRes.ID,
				}

				if err := r.RolePermissionDatabaseSQLRepository.SaveRolePermission(rolePermissionSave); err != nil {
					config.GetLogConfig().WithFields(logrus.Fields{
						"message": "save role permission fail",
						"detail":  err,
					}).Errorln("[SEED ROLE PERMISSIONS]")
					continue
				}
			}
		}
	}
}
