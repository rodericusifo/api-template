// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package serializer

import (
	"api-gateway/internal/domain/core/author/controller/http/rest/response"
	"api-gateway/internal/pkg/constant"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
)

func SerializeAuthorProtoToAuthorResponse(proto *pb_author.Author) *response.AuthorResponse {
	bio := proto.GetBio()
	response := &response.AuthorResponse{
		XID:       proto.GetXID(),
		Name:      proto.GetName(),
		Bio:       &bio,
		CreatedAt: proto.GetCreatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
		UpdatedAt: proto.GetUpdatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
	}
	response.Sanitize()
	return response
}

func SerializeAuthorProtosToAuthorResponses(protos []*pb_author.Author) []*response.AuthorResponse {
	result := make([]*response.AuthorResponse, 0)

	for _, proto := range protos {
		bio := proto.GetBio()
		response := &response.AuthorResponse{
			XID:       proto.GetXID(),
			Name:      proto.GetName(),
			Bio:       &bio,
			CreatedAt: proto.GetCreatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
			UpdatedAt: proto.GetUpdatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
		}
		response.Sanitize()
		result = append(result, response)
	}

	return result
}
