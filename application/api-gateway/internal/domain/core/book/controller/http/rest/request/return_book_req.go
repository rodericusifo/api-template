// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type ReturnBookRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *ReturnBookRequestParams) CustomValidateRequestParams() error {
	return nil
}
