// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"author-service/internal/domain/core/author/service/dto/input"
	"author-service/internal/pkg/util/serializer"

	pb_author "author-service/internal/proto/author"
)

func (h *AuthorHandler) GetAuthor(ctx context.Context, req *pb_author.GetAuthorRequest) (*pb_author.GetAuthorResponse, error) {
	id := req.GetID()
	authorDtoRes, err := h.AuthorService.GetAuthor(&input.GetAuthorDTO{
		XID: req.GetXID(),
		ID:  &id,
	})
	if err != nil {
		return nil, err
	}
	return &pb_author.GetAuthorResponse{
		Data: serializer.SerializeAuthorDTOToAuthorProto(authorDtoRes),
	}, nil
}
