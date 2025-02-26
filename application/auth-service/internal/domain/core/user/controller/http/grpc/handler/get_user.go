// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"auth-service/internal/domain/core/user/service/dto/input"
	"auth-service/internal/pkg/util/serializer"

	pb_user "auth-service/internal/proto/user"
)

func (h *UserHandler) GetUser(ctx context.Context, req *pb_user.GetUserRequest) (*pb_user.GetUserResponse, error) {
	userDtoRes, err := h.UserService.GetUser(&input.GetUserDTO{
		XID: req.GetXID(),
	})
	if err != nil {
		return nil, err
	}
	return &pb_user.GetUserResponse{
		Data: serializer.SerializeUserDTOToUserProto(userDtoRes),
	}, nil
}
