// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"auth-service/internal/domain/core/user/service/dto/input"

	pb_user "auth-service/internal/proto/user"
)

func (h *UserHandler) ValidateUserRolePermissions(ctx context.Context, req *pb_user.ValidateUserRolePermissionsRequest) (*emptypb.Empty, error) {
	err := h.UserService.ValidateUserRolePermissions(&input.ValidateUserRolePermissionsDTO{
		RoleID: req.GetRoleID(),
		Path:   req.GetPath(),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
