// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package input

type UpdateBookDTO struct {
	XID             string
	Title           *string
	ISBN            *string
	Stock           *int32
	PublicationYear *int32
	AuthorID        *uint32
	CategoryID      *uint32
}
