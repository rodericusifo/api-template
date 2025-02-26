// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package sql

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"auth-service/internal/pkg/constant"
)

type Role struct {
	// Migrated Fields
	ID        uint32 `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string
	Slug      string `gorm:"uniqueIndex"`
	Status    constant.RoleStatus
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	if r.XID == "" {
		r.XID = uuid.NewString()
	}
	if r.Slug == "" {
		split_name := strings.Split(r.Name, " ")
		join_split_name := strings.Join(split_name, "_")
		r.Slug = strings.ToLower(join_split_name)
	}
	if r.Status == "" {
		r.Status = constant.ROLE_ACTIVE
	}
	return nil
}

func (Role) TableName() string {
	return "roles"
}
