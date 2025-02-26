// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package response

import (
	"api-gateway/internal/pkg/types"
)

func ResponseSuccess[T any](message string, data T, meta *types.Meta) types.Response[T] {
	return types.Response[T]{
		Success: true,
		Message: message,
		Meta:    meta,
		Data:    data,
	}
}
