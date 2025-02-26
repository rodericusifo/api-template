// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package generator

import (
	"golang.org/x/crypto/bcrypt"

	"auth-service/internal/pkg/config"
)

func GenerateHashFromPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), config.GetVarsConfig().PasswordHashingHashSalt)
	return string(bytes), err
}
