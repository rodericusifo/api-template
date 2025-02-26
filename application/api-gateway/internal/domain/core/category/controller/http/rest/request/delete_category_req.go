// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type DeleteCategoryRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *DeleteCategoryRequestParams) CustomValidateRequestParams() error {
	return nil
}
