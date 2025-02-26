// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type GetAuthorRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *GetAuthorRequestParams) CustomValidateRequestParams() error {
	return nil
}
