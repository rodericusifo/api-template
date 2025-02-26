// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"context"
	"fmt"
	"time"

	"github.com/eko/gocache/lib/v4/store"
)

func (r *UserDatabaseCacheRepository) SetUser(identifier string, payload string, TTL time.Duration) error {
	q := r.Writer()

	if err := q.Set(context.Background(), fmt.Sprintf("%s:%s", r.keyName, identifier), payload, store.WithExpiration(TTL)); err != nil {
		return err
	}

	return nil
}
