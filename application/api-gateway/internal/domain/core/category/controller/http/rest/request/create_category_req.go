// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type CreateCategoryRequestBody struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description" validate:"omitempty"`
}

func (r *CreateCategoryRequestBody) CustomValidateRequestBody() error {
	return nil
}
