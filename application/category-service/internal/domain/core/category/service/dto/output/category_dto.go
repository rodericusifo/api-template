// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package output

import (
	"time"
)

type CategoryDTO struct {
	ID          uint32
	XID         string
	Name        string
	Description *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
