// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package output

import (
	"time"
)

type AuthorDTO struct {
	ID        uint32
	XID       string
	Name      string
	Bio       *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
