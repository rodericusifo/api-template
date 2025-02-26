// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package handler

import (
	"book-service/internal/domain/core/book/service"

	pb_book "book-service/internal/proto/book"
	internal_registry_core_book_service "book-service/internal/registry/core/book/service"
)

type BookHandler struct {
	pb_book.UnimplementedBookHandlerServer
	BookService service.IBookService
}

func InitBookHandler() *BookHandler {
	return &BookHandler{
		UnimplementedBookHandlerServer: pb_book.UnimplementedBookHandlerServer{},
		BookService:                    internal_registry_core_book_service.BookService(),
	}
}
