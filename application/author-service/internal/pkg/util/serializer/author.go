// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package serializer

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"author-service/internal/domain/core/author/service/dto/output"
	"author-service/internal/domain/model/database/sql"

	pb_author "author-service/internal/proto/author"
)

func SerializeAuthorToAuthorDTO(model *sql.Author) *output.AuthorDTO {
	return &output.AuthorDTO{
		ID:        model.ID,
		XID:       model.XID,
		Name:      model.Name,
		Bio:       model.Bio,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func SerializeAuthorsToAuthorDTOs(models []*sql.Author) []*output.AuthorDTO {
	result := make([]*output.AuthorDTO, 0)

	for _, model := range models {
		result = append(result, &output.AuthorDTO{
			XID:       model.XID,
			Name:      model.Name,
			Bio:       model.Bio,
			CreatedAt: model.CreatedAt,
			UpdatedAt: model.UpdatedAt,
		})
	}

	return result
}

func SerializeAuthorDTOToAuthorProto(dto *output.AuthorDTO) *pb_author.Author {
	return &pb_author.Author{
		ID:        dto.ID,
		XID:       dto.XID,
		Name:      dto.Name,
		Bio:       dto.Bio,
		CreatedAt: timestamppb.New(dto.CreatedAt),
		UpdatedAt: timestamppb.New(dto.UpdatedAt),
	}
}

func SerializeAuthorDTOsToAuthorProtos(dtos []*output.AuthorDTO) []*pb_author.Author {
	result := make([]*pb_author.Author, 0)

	for _, dto := range dtos {
		result = append(result, &pb_author.Author{
			XID:       dto.XID,
			Name:      dto.Name,
			Bio:       dto.Bio,
			CreatedAt: timestamppb.New(dto.CreatedAt),
			UpdatedAt: timestamppb.New(dto.UpdatedAt),
		})
	}

	return result
}
