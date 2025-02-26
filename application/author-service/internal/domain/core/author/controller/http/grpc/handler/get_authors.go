// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"author-service/internal/domain/core/author/service/dto/input"
	"author-service/internal/pkg/util/serializer"

	pb_author "author-service/internal/proto/author"
)

func (h *AuthorHandler) GetAuthors(ctx context.Context, req *pb_author.GetAuthorsRequest) (*pb_author.GetAuthorsResponse, error) {
	page, limit := int(req.GetPage()), int(req.GetLimit())

	authorDtosRes, meta, err := h.AuthorService.GetAuthors(&input.GetAuthorsDTO{
		Page:  &page,
		Limit: &limit,
	})
	if err != nil {
		return nil, err
	}
	return &pb_author.GetAuthorsResponse{
		Meta: &pb_author.Meta{
			CurrentPage:      int32(meta.CurrentPage),
			TotalDataPerPage: int32(meta.TotalDataPerPage),
			TotalData:        int32(meta.TotalData),
			TotalPage:        int32(meta.TotalPage),
		},
		Data: serializer.SerializeAuthorDTOsToAuthorProtos(authorDtosRes),
	}, nil
}
