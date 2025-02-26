// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package interfaces

import (
	"github.com/eko/gocache/lib/v4/cache"
)

type IDatabaseCacheRepository[T any] interface {
	Writer() *cache.Cache[T]
	Reader() *cache.Cache[T]
}
