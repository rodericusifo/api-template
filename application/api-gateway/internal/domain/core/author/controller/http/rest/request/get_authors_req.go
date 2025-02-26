// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package request

type GetAuthorsRequestQuery struct {
	Page  *int32 `query:"page" validate:"omitempty,min=0"`
	Limit *int32 `query:"limit" validate:"omitempty,min=0"`
}

func (r *GetAuthorsRequestQuery) CustomValidateRequestQuery() error {
	return nil
}
