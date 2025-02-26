// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package response

type AuthorResponse struct {
	XID       string  `json:"xid"`
	Name      string  `json:"name"`
	Bio       *string `json:"bio"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func (a *AuthorResponse) Sanitize() {
	if a.Bio != nil {
		if *a.Bio == "" {
			a.Bio = nil
		}
	}
}
