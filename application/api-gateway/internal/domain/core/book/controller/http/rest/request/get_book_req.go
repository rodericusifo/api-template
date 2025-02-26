// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type GetBookRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *GetBookRequestParams) CustomValidateRequestParams() error {
	return nil
}
