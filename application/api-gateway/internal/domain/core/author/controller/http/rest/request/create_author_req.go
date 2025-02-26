// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type CreateAuthorRequestBody struct {
	Name string  `json:"name" validate:"required"`
	Bio  *string `json:"bio" validate:"omitempty"`
}

func (r *CreateAuthorRequestBody) CustomValidateRequestBody() error {
	return nil
}
