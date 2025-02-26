// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"context"

	pb_user "api-gateway/external/client/auth-service/http/grpc/proto/user"
)

func (h *UserClientHandler) ValidateUserRolePermissions(ctx context.Context, in *pb_user.ValidateUserRolePermissionsRequest) error {
	_, err := h.Client.ValidateUserRolePermissions(ctx, in)
	if err != nil {
		return err
	}
	return nil
}
