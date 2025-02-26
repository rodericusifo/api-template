// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package input

type CreateBookDTO struct {
	Title           string
	ISBN            *string
	PublicationYear *int32
	Stock           int32
	AuthorID        uint32
	CategoryID      uint32
}
