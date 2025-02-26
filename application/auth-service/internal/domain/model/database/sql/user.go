// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	// Migrated Fields
	ID        uint32 `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
	RoleID uint32
	Role   Role
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.XID == "" {
		u.XID = uuid.NewString()
	}
	return nil
}

func (User) TableName() string {
	return "users"
}
