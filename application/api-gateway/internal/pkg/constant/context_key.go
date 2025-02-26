// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package constant

type ContextKey any

var (
	CONTEXT_KEY_USER         = ContextKey("user")
	CONTEXT_KEY_REQUEST_USER = ContextKey("request-user")
)
