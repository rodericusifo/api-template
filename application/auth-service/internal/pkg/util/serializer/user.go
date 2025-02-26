// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package serializer

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"auth-service/internal/domain/core/user/service/dto/output"
	"auth-service/internal/domain/model/database/sql"

	pb_user "auth-service/internal/proto/user"
)

func SerializeUserToUserDTO(model *sql.User) *output.UserDTO {
	dto := &output.UserDTO{
		ID:        model.ID,
		XID:       model.XID,
		Name:      model.Name,
		Email:     model.Email,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}

	if model.Role.ID != 0 {
		dto.Role = output.RoleDTO{
			ID:     model.Role.ID,
			Name:   model.Role.Name,
			Slug:   model.Role.Slug,
			Status: model.Role.Status,
		}
	}

	return dto
}

func SerializeUserDTOToUserProto(dto *output.UserDTO) *pb_user.User {
	proto := &pb_user.User{
		ID:        dto.ID,
		XID:       dto.XID,
		Name:      dto.Name,
		Email:     dto.Email,
		CreatedAt: timestamppb.New(dto.CreatedAt),
		UpdatedAt: timestamppb.New(dto.UpdatedAt),
	}

	if dto.Role.ID != 0 {
		roleStatus := pb_user.RoleStatus_value[string(dto.Role.Status)]
		proto.Role = &pb_user.Role{
			ID:     dto.Role.ID,
			Name:   dto.Role.Name,
			Slug:   dto.Role.Slug,
			Status: pb_user.RoleStatus(roleStatus),
		}
	}

	return proto
}
