// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package auth

import (
	"context"

	pb_auth "api-gateway/external/client/auth-service/http/grpc/proto/auth"
)

func (h *AuthClientHandler) RegisterAuth(ctx context.Context, in *pb_auth.RegisterAuthRequest) error {
	_, err := h.Client.RegisterAuth(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
