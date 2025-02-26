// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type GetCategoryRequestParams struct {
	XID string `param:"xid" validate:"required,uuid4"`
}

func (r *GetCategoryRequestParams) CustomValidateRequestParams() error {
	return nil
}
