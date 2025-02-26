// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package sql

import (
	"time"
)

type BorrowRecord struct {
	// Migrated Fields
	ID         uint32 `gorm:"primaryKey"`
	BorrowDate time.Time
	ReturnDate *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// Relations
	UserID uint32 `gorm:"index"`
	BookID uint32
	Book   Book `gorm:"constraint:OnDelete:CASCADE;"`
}

func (BorrowRecord) TableName() string {
	return "borrow_records"
}
