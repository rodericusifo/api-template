// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"auth-service/internal/domain/core/auth/service/dto/input"

	pb_auth "auth-service/internal/proto/auth"
)

func (h *AuthHandler) LoginAuth(ctx context.Context, req *pb_auth.LoginAuthRequest) (*pb_auth.LoginAuthResponse, error) {
	authLoginDtoRes, err := h.AuthService.LoginAuth(&input.LoginAuthDTO{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	})
	if err != nil {
		return nil, err
	}
	return &pb_auth.LoginAuthResponse{
		Token: authLoginDtoRes.Token,
	}, nil
}
