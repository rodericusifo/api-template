// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package role

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/util/validator"
)

type RolePayload struct {
	XID  string `validate:"required,uuid4"`
	Name string `validate:"required"`
	Slug string `validate:"required"`
}

func (r *RolePayload) CustomValidatePayload() error {
	return nil
}

type RolesPayload []*RolePayload

func (r RolesPayload) Validate() {
	valid := make([]*RolePayload, 0, len(r))
	for _, role := range r {
		err := validator.ValidatePayload(role)
		if err != nil {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("validation failed: role with xid %s", role.XID),
				"detail":  err,
			}).Errorln("[VALIDATE]")
			continue
		}
		valid = append(valid, role)
	}
	r = valid
}

var Roles = RolesPayload{
	{
		XID:  "8d5ccbad-b0d4-4d2c-9460-3eac3ce84112",
		Name: "Super Admin",
		Slug: "super_admin",
	},
	{
		XID:  "fbb6853e-4902-445a-b706-64f98453c59a",
		Name: "Librarian",
		Slug: "librarian",
	},
	{
		XID:  "f936a059-0b60-469d-826b-d50a4a21661b",
		Name: "Member",
		Slug: "member",
	},
}
