// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package generator

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/types"
)

func GenerateJWTTokenFromClaims(claims *types.JwtCustomClaims) (string, error) {
	claims.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(config.GetVarsConfig().JWTExpiredDuration)),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(config.GetVarsConfig().JWTSecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}
