// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package constant

type Default any
type DefaultEnv any

var (
	DEFAULT_TIME_LAYOUT = Default("2006-01-02 15:04:05 MST")
)
var (
	DEFAULT_ENV_SERVER_PORT = DefaultEnv(8080)
)
