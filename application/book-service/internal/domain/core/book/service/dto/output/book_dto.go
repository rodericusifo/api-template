// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package output

import (
	"time"
)

type BookDTO struct {
	ID              uint32
	XID             string
	Title           string
	ISBN            *string
	PublicationYear *int32
	Stock           int32
	BorrowStock     int32
	CreatedAt       time.Time
	UpdatedAt       time.Time
	AuthorID        uint32
	CategoryID      uint32
}
