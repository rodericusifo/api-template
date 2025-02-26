// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package sql

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	// Migrated Fields
	ID              uint32  `gorm:"primaryKey"`
	XID             string  `gorm:"column:xid"`
	Title           string  `gorm:"uniqueIndex"`
	ISBN            *string `gorm:"column:isbn;uniqueIndex"`
	PublicationYear *int32
	Stock           int32
	BorrowStock     int32
	CreatedAt       time.Time
	UpdatedAt       time.Time

	// Relations
	AuthorID   uint32 `gorm:"index"`
	CategoryID uint32 `gorm:"index"`
}

func (b *Book) BeforeCreate(tx *gorm.DB) error {
	if b.XID == "" {
		b.XID = uuid.NewString()
	}
	if b.ISBN != nil {
		if *b.ISBN == "" {
			b.ISBN = nil
		}
	}
	if b.PublicationYear != nil {
		if *b.PublicationYear == 0 {
			b.PublicationYear = nil
		}
	}
	return nil
}

func (Book) TableName() string {
	return "books"
}
