// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"auth-service/internal/pkg/constant"
)

type Permission struct {
	// Migrated Fields
	ID        uint32 `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string
	Slug      string `gorm:"uniqueIndex"`
	Path      string
	Status    constant.PermissionStatus
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
}

func (p *Permission) BeforeCreate(tx *gorm.DB) error {
	if p.XID == "" {
		p.XID = uuid.NewString()
	}
	if p.Status == "" {
		p.Status = constant.PERMISSION_ACTIVE
	}
	return nil
}

func (Permission) TableName() string {
	return "permissions"
}
