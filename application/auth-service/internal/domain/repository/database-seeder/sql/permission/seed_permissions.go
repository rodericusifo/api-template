// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package permission

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/types"
)

func (r *PermissionDatabaseSeederSQLRepository) SeedPermissions() {
	Permissions.Validate()

	for _, Permission := range Permissions {
		firstPermissionModelRes, err := r.PermissionDatabaseSQLRepository.FirstPermission(&types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "xid", Operator: "=", Value: Permission.XID},
					{Field: "slug", Operator: "=", Value: Permission.Slug},
				},
			},
		})
		if err != nil && err != gorm.ErrRecordNotFound {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": "get permission fail",
				"detail":  err,
			}).Errorln("[SEED PERMISSIONS]")
			continue
		}
		if firstPermissionModelRes != nil {
			logrus.WithFields(logrus.Fields{
				"message": fmt.Sprintf("permission with xid %s and slug %s already registered", Permission.XID, Permission.Slug),
			}).Errorln("[SEED PERMISSIONS]")
			continue
		}

		permissionSave := &sql.Permission{
			XID:  Permission.XID,
			Name: Permission.Name,
			Slug: Permission.Slug,
			Path: Permission.Path,
		}

		if err := r.PermissionDatabaseSQLRepository.SavePermission(permissionSave); err != nil {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": "save permission fail",
				"detail":  err,
			}).Errorln("[SEED PERMISSIONS]")
			continue
		}
	}
}
