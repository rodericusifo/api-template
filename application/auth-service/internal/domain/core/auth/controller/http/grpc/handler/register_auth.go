// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"auth-service/internal/domain/core/auth/service/dto/input"

	pb_auth "auth-service/internal/proto/auth"
)

func (h *AuthHandler) RegisterAuth(ctx context.Context, req *pb_auth.RegisterAuthRequest) (*emptypb.Empty, error) {
	err := h.AuthService.RegisterAuth(&input.RegisterAuthDTO{
		Name:     req.GetName(),
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		RoleSlug: req.GetRoleSlug(),
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
