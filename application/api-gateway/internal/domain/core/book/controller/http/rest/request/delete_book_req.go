// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type DeleteBookRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *DeleteBookRequestParams) CustomValidateRequestParams() error {
	return nil
}
