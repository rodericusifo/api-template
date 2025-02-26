// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"author-service/internal/domain/core/author/service"

	pb_author "author-service/internal/proto/author"
	internal_registry_core_author_service "author-service/internal/registry/core/author/service"
)

type AuthorHandler struct {
	pb_author.UnimplementedAuthorHandlerServer
	AuthorService service.IAuthorService
}

func InitAuthorHandler() *AuthorHandler {
	return &AuthorHandler{
		UnimplementedAuthorHandlerServer: pb_author.UnimplementedAuthorHandlerServer{},
		AuthorService:                    internal_registry_core_author_service.AuthorService(),
	}
}
