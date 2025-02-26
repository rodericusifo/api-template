// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package serializer

import (
	"api-gateway/internal/domain/core/book/controller/http/rest/response"
	"api-gateway/internal/pkg/constant"

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

func SerializeBookProtoToBookResponse(book *pb_book.Book, author *pb_author.Author, category *pb_category.Category) *response.BookResponse {
	publicationYear := book.GetPublicationYear()
	isbn := book.GetISBN()

	res := &response.BookResponse{
		XID:             book.GetXID(),
		Title:           book.GetTitle(),
		ISBN:            &isbn,
		PublicationYear: &publicationYear,
		Stock:           book.GetStock(),
		BorrowStock:     book.GetBorrowStock(),
		CreatedAt:       book.GetCreatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
		UpdatedAt:       book.GetUpdatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
	}

	if author != nil {
		bio := author.GetBio()
		res.Author = &response.AuthorResponse{
			Name: author.GetName(),
			Bio:  &bio,
		}
		res.Author.Sanitize()
	}

	if category != nil {
		description := category.GetDescription()
		res.Category = &response.CategoryResponse{
			Name:        category.GetName(),
			Description: &description,
		}
		res.Category.Sanitize()
	}

	res.Sanitize()

	return res
}

func SerializeBookProtosToBookResponses(protos []*pb_book.Book) []*response.BookResponse {
	result := make([]*response.BookResponse, 0)

	for _, proto := range protos {
		publicationYear := proto.GetPublicationYear()
		isbn := proto.GetISBN()

		res := &response.BookResponse{
			XID:             proto.GetXID(),
			Title:           proto.GetTitle(),
			ISBN:            &isbn,
			PublicationYear: &publicationYear,
			Stock:           proto.GetStock(),
			BorrowStock:     proto.GetBorrowStock(),
			CreatedAt:       proto.GetCreatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
			UpdatedAt:       proto.GetUpdatedAt().AsTime().Format(constant.DEFAULT_TIME_LAYOUT.(string)),
		}
		res.Sanitize()
		result = append(result, res)
	}

	return result
}
