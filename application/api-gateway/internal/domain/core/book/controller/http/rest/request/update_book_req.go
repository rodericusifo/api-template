// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type UpdateBookRequestBody struct {
	Title           *string `json:"title" validate:"omitempty"`
	ISBN            *string `json:"isbn" validate:"omitempty,isbn"`
	PublicationYear *int32  `json:"publication_year" validate:"omitempty,min=1"`
	Stock           *int32  `json:"stock" validate:"omitempty,min=1"`
	AuthorXID       *string `json:"author_xid" validate:"omitempty,uuid4"`
	CategoryXID     *string `json:"category_xid" validate:"omitempty,uuid4"`
}

func (r *UpdateBookRequestBody) CustomValidateRequestBody() error {
	return nil
}

type UpdateBookRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *UpdateBookRequestParams) CustomValidateRequestParams() error {
	return nil
}
