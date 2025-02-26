// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package types

import (
	"api-gateway/internal/pkg/constant"
)

type RequestRole struct {
	ID     uint32              `validate:"required"`
	Name   string              `validate:"required"`
	Slug   string              `validate:"required"`
	Status constant.RoleStatus `validate:"required"`
}

type RequestUser struct {
	ID    uint32      `validate:"required"`
	XID   string      `validate:"required,uuid4"`
	Name  string      `validate:"required"`
	Email string      `validate:"required,email"`
	Role  RequestRole `validate:"required,dive"`
}

func (r *RequestUser) CustomValidateRequestUser() error {
	return nil
}
