// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package serializer

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"category-service/internal/domain/core/category/service/dto/output"
	"category-service/internal/domain/model/database/sql"

	pb_category "category-service/internal/proto/category"
)

func SerializeCategoryToCategoryDTO(model *sql.Category) *output.CategoryDTO {
	return &output.CategoryDTO{
		ID:          model.ID,
		XID:         model.XID,
		Name:        model.Name,
		Description: model.Description,
		CreatedAt:   model.CreatedAt,
		UpdatedAt:   model.UpdatedAt,
	}
}

func SerializeCategoriesToCategoryDTOs(models []*sql.Category) []*output.CategoryDTO {
	result := make([]*output.CategoryDTO, 0)

	for _, model := range models {
		result = append(result, &output.CategoryDTO{
			XID:         model.XID,
			Name:        model.Name,
			Description: model.Description,
			CreatedAt:   model.CreatedAt,
			UpdatedAt:   model.UpdatedAt,
		})
	}

	return result
}

func SerializeCategoryDTOToCategoryProto(dto *output.CategoryDTO) *pb_category.Category {
	return &pb_category.Category{
		ID:          dto.ID,
		XID:         dto.XID,
		Name:        dto.Name,
		Description: dto.Description,
		CreatedAt:   timestamppb.New(dto.CreatedAt),
		UpdatedAt:   timestamppb.New(dto.UpdatedAt),
	}
}

func SerializeCategoryDTOsToCategoryProtos(dtos []*output.CategoryDTO) []*pb_category.Category {
	result := make([]*pb_category.Category, 0)

	for _, dto := range dtos {
		result = append(result, &pb_category.Category{
			XID:         dto.XID,
			Name:        dto.Name,
			Description: dto.Description,
			CreatedAt:   timestamppb.New(dto.CreatedAt),
			UpdatedAt:   timestamppb.New(dto.UpdatedAt),
		})
	}

	return result
}
