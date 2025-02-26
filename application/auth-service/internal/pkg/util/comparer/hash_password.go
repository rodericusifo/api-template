// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package comparer

import (
	"golang.org/x/crypto/bcrypt"
)

func CompareHashAndPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
