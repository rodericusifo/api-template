// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package response

type AuthorResponse struct {
	Name string  `json:"name"`
	Bio  *string `json:"bio"`
}

func (a *AuthorResponse) Sanitize() {
	if a.Bio != nil {
		if *a.Bio == "" {
			a.Bio = nil
		}
	}
}

type CategoryResponse struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

func (c *CategoryResponse) Sanitize() {
	if c.Description != nil {
		if *c.Description == "" {
			c.Description = nil
		}
	}
}

type BookResponse struct {
	XID             string            `json:"xid"`
	Title           string            `json:"title"`
	ISBN            *string           `json:"isbn"`
	PublicationYear *int32            `json:"publication_year"`
	Stock           int32             `json:"stock"`
	BorrowStock     int32             `json:"borrow_stock"`
	CreatedAt       string            `json:"created_at"`
	UpdatedAt       string            `json:"updated_at"`
	Author          *AuthorResponse   `json:"author,omitempty"`
	Category        *CategoryResponse `json:"category,omitempty"`
}

func (b *BookResponse) Sanitize() {
	if b.ISBN != nil {
		if *b.ISBN == "" {
			b.ISBN = nil
		}
	}
	if b.PublicationYear != nil {
		if *b.PublicationYear == 0 {
			b.PublicationYear = nil
		}
	}
}
