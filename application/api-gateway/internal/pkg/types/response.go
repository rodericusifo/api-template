// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package types

type Meta struct {
	CurrentPage      int32 `json:"current_page"`
	TotalDataPerPage int32 `json:"total_data_per_page"`
	TotalData        int32 `json:"total_data"`
	TotalPage        int32 `json:"total_page"`
}

type Response[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Meta    *Meta  `json:"meta,omitempty"`
	Data    T      `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}
