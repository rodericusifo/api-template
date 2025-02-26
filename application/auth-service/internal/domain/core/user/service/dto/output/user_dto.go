// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package output

import (
	"time"

	"auth-service/internal/pkg/constant"
)

type RoleDTO struct {
	ID     uint32
	Name   string
	Slug   string
	Status constant.RoleStatus
}

type UserDTO struct {
	ID        uint32
	XID       string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
	Role      RoleDTO
}
