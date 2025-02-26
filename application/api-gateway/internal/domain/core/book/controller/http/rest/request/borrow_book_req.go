// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type BorrowBookRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *BorrowBookRequestParams) CustomValidateRequestParams() error {
	return nil
}
