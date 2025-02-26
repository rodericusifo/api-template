// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package interfaces

import (
	"gorm.io/gorm"
)

type IDatabaseSQLRepository interface {
	BeginTransaction(tx *gorm.DB)
	EndTransaction()
	Writer() *gorm.DB
	Reader() *gorm.DB
}
