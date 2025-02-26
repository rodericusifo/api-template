// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"author-service/internal/domain/core/author/service/dto/input"

	pb_author "author-service/internal/proto/author"
)

func (h *AuthorHandler) UpdateAuthor(ctx context.Context, req *pb_author.UpdateAuthorRequest) (*emptypb.Empty, error) {
	name, bio := req.GetName(), req.GetBio()
	err := h.AuthorService.UpdateAuthor(&input.UpdateAuthorDTO{
		XID:  req.GetXID(),
		Name: &name,
		Bio:  &bio,
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
