// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"context"
	"fmt"
)

func (r *UserDatabaseCacheRepository) GetUser(identifier string) (string, error) {
	q := r.Reader()

	strUser, err := q.Get(context.Background(), fmt.Sprintf("%s:%s", r.keyName, identifier))
	if err != nil {
		return "", err
	}

	return string(strUser), nil
}
