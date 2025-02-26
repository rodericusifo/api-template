// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package sql

import (
	"time"
)

type RolePermission struct {
	// Migrated Fields
	ID        uint32 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	RoleID       uint32
	Role         Role `gorm:"constraint:OnDelete:CASCADE;"`
	PermissionID uint32
	Permission   Permission `gorm:"constraint:OnDelete:CASCADE;"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
