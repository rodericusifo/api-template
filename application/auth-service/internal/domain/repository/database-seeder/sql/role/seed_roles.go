// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package role

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"auth-service/internal/domain/model/database/sql"
	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/types"
)

func (r *RoleDatabaseSeederSQLRepository) SeedRoles() {
	Roles.Validate()

	for _, Role := range Roles {
		firstRoleModelRes, err := r.RoleDatabaseSQLRepository.FirstRole(&types.QuerySQL{
			Selects: []types.SelectQuerySQLOperation{
				{Field: "id"},
			},
			Searches: [][]types.SearchQuerySQLOperation{
				{
					{Field: "xid", Operator: "=", Value: Role.XID},
					{Field: "slug", Operator: "=", Value: Role.Slug},
				},
			},
		})
		if err != nil && err != gorm.ErrRecordNotFound {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": "get role fail",
				"detail":  err,
			}).Errorln("[SEED ROLES]")
			continue
		}
		if firstRoleModelRes != nil {
			logrus.WithFields(logrus.Fields{
				"message": fmt.Sprintf("role with xid %s and slug %s already registered", Role.XID, Role.Slug),
			}).Errorln("[SEED ROLES]")
			continue
		}

		roleSave := &sql.Role{
			XID:  Role.XID,
			Name: Role.Name,
			Slug: Role.Slug,
		}

		if err := r.RoleDatabaseSQLRepository.SaveRole(roleSave); err != nil {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": "save role fail",
				"detail":  err,
			}).Errorln("[SEED ROLES]")
			continue
		}
	}
}
