// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package response

import (
	"api-gateway/internal/pkg/types"
)

func ResponseFail(message string, err any) types.Response[any] {
	return types.Response[any]{
		Success: false,
		Message: message,
		Error:   err,
	}
}
