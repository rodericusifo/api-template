// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"author-service/internal/domain/core/author/service/dto/input"

	pb_author "author-service/internal/proto/author"
)

func (h *AuthorHandler) DeleteAuthor(ctx context.Context, req *pb_author.DeleteAuthorRequest) (*emptypb.Empty, error) {
	err := h.AuthorService.DeleteAuthor(&input.DeleteAuthorDTO{
		XID: req.GetXID(),
	})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
