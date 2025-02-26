// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package config

import (
	"api-gateway/internal/pkg/types"

	jwtware "github.com/gofiber/contrib/jwt"
)

var (
	jwtAuthConfig jwtware.Config
)

func ConfigureAuth() {
	jwtAuthConfig = jwtware.Config{
		Claims: &types.JwtCustomClaims{},
		SigningKey: jwtware.SigningKey{
			Key: []byte(GetVarsConfig().JWTSecretKey),
		},
	}
}

func GetAuthConfig() jwtware.Config {
	return jwtAuthConfig
}
