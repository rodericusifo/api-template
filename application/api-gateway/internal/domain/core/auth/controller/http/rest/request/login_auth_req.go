// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type LoginAuthRequestBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

func (r *LoginAuthRequestBody) CustomValidateRequestBody() error {
	return nil
}
