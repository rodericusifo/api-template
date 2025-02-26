// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type UpdateCategoryRequestBody struct {
	Name        *string `json:"name" validate:"omitempty"`
	Description *string `json:"description" validate:"omitempty"`
}

func (r *UpdateCategoryRequestBody) CustomValidateRequestBody() error {
	return nil
}

type UpdateCategoryRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *UpdateCategoryRequestParams) CustomValidateRequestParams() error {
	return nil
}
