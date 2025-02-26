// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type UpdateAuthorRequestBody struct {
	Name *string `json:"name" validate:"omitempty"`
	Bio  *string `json:"bio" validate:"omitempty"`
}

func (r *UpdateAuthorRequestBody) CustomValidateRequestBody() error {
	return nil
}

type UpdateAuthorRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *UpdateAuthorRequestParams) CustomValidateRequestParams() error {
	return nil
}
