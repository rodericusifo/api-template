// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package service

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"author-service/internal/domain/core/author/service/dto/input"
	"author-service/internal/domain/core/author/service/dto/output"
	"author-service/internal/pkg/types"
	"author-service/internal/pkg/util/counter"
	"author-service/internal/pkg/util/definer"
	"author-service/internal/pkg/util/serializer"
)

func (s *AuthorService) GetAuthors(payload *input.GetAuthorsDTO) (output.GetAuthorsDTO, *types.Meta, error) {
	page, limit := definer.DefinePaginationPageLimit(payload.Page, payload.Limit)

	authorListModelRes, err := s.AuthorDatabaseSQLRepository.FindAuthors(&types.QuerySQL{
		Offset: counter.CountPaginationOffset(page, limit),
		Limit:  limit,
	})
	if err != nil {
		return nil, nil, err
	}
	countAuthorListModelRes := len(authorListModelRes)

	if len(authorListModelRes) < 1 {
		return nil, nil, status.Error(codes.NotFound, "authors not found")
	}

	countAuthorAllModelRes, err := s.AuthorDatabaseSQLRepository.CountAuthors(nil)
	if err != nil {
		return nil, nil, err
	}

	authorListDto := serializer.SerializeAuthorsToAuthorDTOs(authorListModelRes)

	meta := &types.Meta{
		CurrentPage:      int32(page),
		TotalDataPerPage: int32(countAuthorListModelRes),
		TotalData:        int32(countAuthorAllModelRes),
	}

	meta.TotalPage = int32(counter.CountPaginationTotalPage(int(meta.TotalDataPerPage), int(meta.TotalData)))

	return authorListDto, meta, nil
}
