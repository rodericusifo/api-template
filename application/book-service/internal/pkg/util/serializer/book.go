// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package serializer

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"book-service/internal/domain/core/book/service/dto/output"
	"book-service/internal/domain/model/database/sql"

	pb_book "book-service/internal/proto/book"
)

func SerializeBookToBookDTO(model *sql.Book) *output.BookDTO {
	return &output.BookDTO{
		ID:              model.ID,
		XID:             model.XID,
		Title:           model.Title,
		ISBN:            model.ISBN,
		PublicationYear: model.PublicationYear,
		Stock:           model.Stock,
		BorrowStock:     model.BorrowStock,
		CreatedAt:       model.CreatedAt,
		UpdatedAt:       model.UpdatedAt,
		AuthorID:        model.AuthorID,
		CategoryID:      model.CategoryID,
	}
}

func SerializeBooksToBookDTOs(models []*sql.Book) []*output.BookDTO {
	result := make([]*output.BookDTO, 0)

	for _, model := range models {
		result = append(result, &output.BookDTO{
			XID:             model.XID,
			Title:           model.Title,
			ISBN:            model.ISBN,
			PublicationYear: model.PublicationYear,
			Stock:           model.Stock,
			BorrowStock:     model.BorrowStock,
			CreatedAt:       model.CreatedAt,
			UpdatedAt:       model.UpdatedAt,
		})
	}

	return result
}

func SerializeBookDTOToBookProto(dto *output.BookDTO) *pb_book.Book {
	proto := &pb_book.Book{
		ID:              dto.ID,
		XID:             dto.XID,
		Title:           dto.Title,
		ISBN:            dto.ISBN,
		PublicationYear: dto.PublicationYear,
		Stock:           dto.Stock,
		BorrowStock:     dto.BorrowStock,
		CreatedAt:       timestamppb.New(dto.CreatedAt),
		UpdatedAt:       timestamppb.New(dto.UpdatedAt),
		AuthorID:        dto.AuthorID,
		CategoryID:      dto.CategoryID,
	}

	return proto
}

func SerializeBookDTOsToBookProtos(dtos []*output.BookDTO) []*pb_book.Book {
	result := make([]*pb_book.Book, 0)

	for _, dto := range dtos {

		result = append(result, &pb_book.Book{
			XID:             dto.XID,
			Title:           dto.Title,
			ISBN:            dto.ISBN,
			Stock:           dto.Stock,
			BorrowStock:     dto.BorrowStock,
			PublicationYear: dto.PublicationYear,
			CreatedAt:       timestamppb.New(dto.CreatedAt),
			UpdatedAt:       timestamppb.New(dto.UpdatedAt),
		})
	}

	return result
}
