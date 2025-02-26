// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"auth-service/internal/domain/core/user/service"

	pb_user "auth-service/internal/proto/user"
	internal_registry_core_user_service "auth-service/internal/registry/core/user/service"
)

type UserHandler struct {
	pb_user.UnimplementedUserHandlerServer
	UserService service.IUserService
}

func InitUserHandler() *UserHandler {
	return &UserHandler{
		UnimplementedUserHandlerServer: pb_user.UnimplementedUserHandlerServer{},
		UserService:                    internal_registry_core_user_service.UserService(),
	}
}
