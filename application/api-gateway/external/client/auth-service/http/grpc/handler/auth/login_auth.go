// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package auth

import (
	"context"

	pb_auth "api-gateway/external/client/auth-service/http/grpc/proto/auth"
)

func (h *AuthClientHandler) LoginAuth(ctx context.Context, in *pb_auth.LoginAuthRequest) (*pb_auth.LoginAuthResponse, error) {
	resp, err := h.Client.LoginAuth(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
