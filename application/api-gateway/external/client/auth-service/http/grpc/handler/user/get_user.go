// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"context"

	pb_user "api-gateway/external/client/auth-service/http/grpc/proto/user"
)

func (h *UserClientHandler) GetUser(ctx context.Context, in *pb_user.GetUserRequest) (*pb_user.GetUserResponse, error) {
	resp, err := h.Client.GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
