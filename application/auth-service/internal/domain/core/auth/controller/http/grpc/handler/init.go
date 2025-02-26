// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"auth-service/internal/domain/core/auth/service"

	pb_auth "auth-service/internal/proto/auth"
	internal_registry_core_auth_service "auth-service/internal/registry/core/auth/service"
)

type AuthHandler struct {
	pb_auth.UnimplementedAuthHandlerServer
	AuthService service.IAuthService
}

func InitAuthHandler() *AuthHandler {
	return &AuthHandler{
		UnimplementedAuthHandlerServer: pb_auth.UnimplementedAuthHandlerServer{},
		AuthService:                    internal_registry_core_auth_service.AuthService(),
	}
}
