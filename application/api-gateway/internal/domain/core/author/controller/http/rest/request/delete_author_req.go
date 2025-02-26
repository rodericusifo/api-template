// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type DeleteAuthorRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *DeleteAuthorRequestParams) CustomValidateRequestParams() error {
	return nil
}
