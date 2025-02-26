// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package rolepermission

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/util/validator"
)

type RolePermissionPayload struct {
	RoleSlug        string   `validate:"required"`
	PermissionsSlug []string `validate:"omitempty,dive,min=1"`
}

func (r *RolePermissionPayload) CustomValidatePayload() error {
	return nil
}

type RolePermissionsPayload []*RolePermissionPayload

func (r RolePermissionsPayload) Validate() {
	valid := make([]*RolePermissionPayload, 0, len(r))
	for _, rolePermission := range r {
		err := validator.ValidatePayload(rolePermission)
		if err != nil {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("validation failed: role permission with role_slug %s", rolePermission.RoleSlug),
				"detail":  err,
			}).Errorln("[VALIDATE]")
			continue
		}
		valid = append(valid, rolePermission)
	}
	r = valid
}

var RolePermissions = RolePermissionsPayload{
	{
		RoleSlug:        "super_admin",
		PermissionsSlug: []string{},
	},
	{
		RoleSlug: "librarian",
		PermissionsSlug: []string{
			"author.create",
			"author.list",
			"author.detail",
			"author.update",
			"author.delete",
			"category.create",
			"category.list",
			"category.detail",
			"category.update",
			"category.delete",
			"book.create",
			"book.list",
			"book.detail",
			"book.update",
			"book.delete",
			"book.borrow",
			"book.return",
		},
	},
	{
		RoleSlug: "member",
		PermissionsSlug: []string{
			"author.list",
			"author.detail",
			"category.list",
			"category.detail",
			"book.list",
			"book.detail",
			"book.borrow",
			"book.return",
		},
	},
}
