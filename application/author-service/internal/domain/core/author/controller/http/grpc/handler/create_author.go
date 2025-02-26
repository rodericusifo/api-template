// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"author-service/internal/domain/core/author/service/dto/input"

	pb_author "author-service/internal/proto/author"
)

func (h *AuthorHandler) CreateAuthor(ctx context.Context, req *pb_author.CreateAuthorRequest) (*emptypb.Empty, error) {
	bio := req.GetBio()
	err := h.AuthorService.CreateAuthor(&input.CreateAuthorDTO{
		Name: req.GetName(),
		Bio:  &bio,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
