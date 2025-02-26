// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	// Migrated Fields
	ID          uint32 `gorm:"primaryKey"`
	XID         string `gorm:"column:xid"`
	Name        string `gorm:"uniqueIndex"`
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time

	// Relations
}

func (c *Category) BeforeCreate(tx *gorm.DB) error {
	if c.XID == "" {
		c.XID = uuid.NewString()
	}
	if c.Description != nil {
		if *c.Description == "" {
			c.Description = nil
		}
	}
	return nil
}

func (Category) TableName() string {
	return "categories"
}
