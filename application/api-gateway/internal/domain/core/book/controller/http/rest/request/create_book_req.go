// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type CreateBookRequestBody struct {
	Title           string  `json:"title" validate:"required"`
	ISBN            *string `json:"isbn" validate:"omitempty,isbn"`
	PublicationYear *int32  `json:"publication_year" validate:"omitempty,min=1"`
	Stock           int32   `json:"stock" validate:"required,min=1"`
	AuthorXID       string  `json:"author_xid" validate:"required,uuid4"`
	CategoryXID     string  `json:"category_xid" validate:"required,uuid4"`
}

func (r *CreateBookRequestBody) CustomValidateRequestBody() error {
	return nil
}
