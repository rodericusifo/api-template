// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Author struct {
	// Migrated Fields
	ID        uint32 `gorm:"primaryKey"`
	XID       string `gorm:"column:xid"`
	Name      string `gorm:"uniqueIndex"`
	Bio       *string
	CreatedAt time.Time
	UpdatedAt time.Time

	// Relations
}

func (a *Author) BeforeCreate(tx *gorm.DB) error {
	if a.XID == "" {
		a.XID = uuid.NewString()
	}
	if a.Bio != nil {
		if *a.Bio == "" {
			a.Bio = nil
		}
	}
	return nil
}

func (Author) TableName() string {
	return "authors"
}
