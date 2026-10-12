// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"context"
	"fmt"
)

func (r *UserCacheRepository) GetUser(identifier string) (string, error) {
	q := r.Reader()
	if q == nil {
		// Cache is disabled/not configured: treat as a cache miss so callers
		// fall through to the source of truth instead of panicking.
		return "", nil
	}

	strUser, err := q.Get(context.Background(), fmt.Sprintf("%s:%s", r.keyName, identifier))
	if err != nil {
		return "", err
	}

	return string(strUser), nil
}
