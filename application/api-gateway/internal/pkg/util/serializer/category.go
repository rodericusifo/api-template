// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package serializer

import (
	"api-gateway/internal/domain/core/category/controller/http/rest/response"
	"api-gateway/internal/pkg/constant"

	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func SerializeCategoryProtoToCategoryResponse(proto *pb_category.Category) *response.CategoryResponse {
	description := proto.GetDescription()
	response := &response.CategoryResponse{
		XID:         proto.GetXID(),
		Name:        proto.GetName(),
		Description: &description,
		CreatedAt:   proto.GetCreatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
		UpdatedAt:   proto.GetUpdatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
	}
	response.Sanitize()
	return response
}

func SerializeCategoryProtosToCategoryResponses(protos []*pb_category.Category) []*response.CategoryResponse {
	result := make([]*response.CategoryResponse, 0)

	for _, proto := range protos {
		description := proto.GetDescription()
		response := &response.CategoryResponse{
			XID:         proto.GetXID(),
			Name:        proto.GetName(),
			Description: &description,
			CreatedAt:   proto.GetCreatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
			UpdatedAt:   proto.GetUpdatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
		}
		response.Sanitize()
		result = append(result, response)
	}

	return result
}
