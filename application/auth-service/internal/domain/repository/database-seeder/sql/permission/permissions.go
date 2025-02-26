// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package permission

import (
	"fmt"

	"github.com/sirupsen/logrus"

	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/util/validator"
)

type PermissionPayload struct {
	XID  string `validate:"required,uuid4"`
	Name string `validate:"required"`
	Slug string `validate:"required"`
	Path string `validate:"required"`
}

func (r *PermissionPayload) CustomValidatePayload() error {
	return nil
}

type PermissionsPayload []*PermissionPayload

func (p PermissionsPayload) Validate() {
	valid := make([]*PermissionPayload, 0, len(p))
	for _, permission := range p {
		err := validator.ValidatePayload(permission)
		if err != nil {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": fmt.Sprintf("validation failed: permission with xid %s", permission.XID),
				"detail":  err,
			}).Errorln("[VALIDATE]")
			continue
		}
		valid = append(valid, permission)
	}
	p = valid
}

var Permissions = PermissionsPayload{
	// AUTHOR
	{
		XID:  "aa80c2b3-e8f7-402d-85b9-4f1d965bbcdc",
		Name: "Author Create",
		Slug: "author.create",
		Path: "/v1/authors/create",
	},
	{
		XID:  "6cf39bc8-c2a7-48dc-8ffb-89a9f60228d1",
		Name: "Author List",
		Slug: "author.list",
		Path: "/v1/authors/list",
	},
	{
		XID:  "431f2415-af89-4d08-8fad-850c4a9251f1",
		Name: "Author Detail",
		Slug: "author.detail",
		Path: "/v1/authors/:xid/detail",
	},
	{
		XID:  "cb35665c-6f34-443c-88ce-9c71a67aef5c",
		Name: "Author Update",
		Slug: "author.update",
		Path: "/v1/authors/:xid/update",
	},
	{
		XID:  "e0cfd5c7-4e67-4b70-b19a-c9790892f2c3",
		Name: "Author Delete",
		Slug: "author.delete",
		Path: "/v1/authors/:xid/delete",
	},
	// CATEGORY
	{
		XID:  "6fa5ae15-64fc-4c1e-8477-063ef6964dd1",
		Name: "Category Create",
		Slug: "category.create",
		Path: "/v1/categories/create",
	},
	{
		XID:  "59f8594e-2427-4638-b68a-e1a4ffc4057a",
		Name: "Category List",
		Slug: "category.list",
		Path: "/v1/categories/list",
	},
	{
		XID:  "da186374-db94-4df8-94c9-196a104a0269",
		Name: "Category Detail",
		Slug: "category.detail",
		Path: "/v1/categories/:xid/detail",
	},
	{
		XID:  "a5124700-4d42-4b7e-94e2-a7a1f1dac154",
		Name: "Category Update",
		Slug: "category.update",
		Path: "/v1/categories/:xid/update",
	},
	{
		XID:  "2ec4cfc6-eee0-4364-97ee-fd18eb42b5f5",
		Name: "Category Delete",
		Slug: "category.delete",
		Path: "/v1/categories/:xid/delete",
	},
	// BOOK
	{
		XID:  "a6d49827-4f42-49bb-9219-68134cef456e",
		Name: "Book Create",
		Slug: "book.create",
		Path: "/v1/books/create",
	},
	{
		XID:  "eb45511c-cc94-419d-9a74-f53742087b2a",
		Name: "Book List",
		Slug: "book.list",
		Path: "/v1/books/list",
	},
	{
		XID:  "1918f2b3-b6a1-48f8-bebe-eb2f977854c7",
		Name: "Book Detail",
		Slug: "book.detail",
		Path: "/v1/books/:xid/detail",
	},
	{
		XID:  "d3e83ce0-4321-4f59-b39e-a5012525a7cd",
		Name: "Book Update",
		Slug: "book.update",
		Path: "/v1/books/:xid/update",
	},
	{
		XID:  "8930a353-b66e-4b0e-a69d-82421e0cff91",
		Name: "Book Delete",
		Slug: "book.delete",
		Path: "/v1/books/:xid/delete",
	},
	{
		XID:  "32c99cfb-7111-49ea-8ee2-28cfe2b73efb",
		Name: "Book Borrow",
		Slug: "book.borrow",
		Path: "/v1/books/:xid/borrow",
	},
	{
		XID:  "1c0bee4c-4207-4712-a8c5-e51eb67db807",
		Name: "Book Return",
		Slug: "book.return",
		Path: "/v1/books/:xid/return",
	},
}
