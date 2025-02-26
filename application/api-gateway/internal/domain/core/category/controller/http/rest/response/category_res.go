// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package response

type CategoryResponse struct {
	XID         string  `json:"xid"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func (c *CategoryResponse) Sanitize() {
	if c.Description != nil {
		if *c.Description == "" {
			c.Description = nil
		}
	}
}
