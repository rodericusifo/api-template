// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package constant

import (
	"author-service/internal/pkg/types"
)

type Default any
type DefaultEnv any
type DefaultSelects []types.SelectQuerySQLOperation
type DefaultSelectsJoin []types.SelectJoinQuerySQLOperation

var (
	DEFAULT_TIME_LAYOUT = Default("2006-01-02 15:04:05 MST")
)
var (
	DEFAULT_ENV_SERVER_PORT = DefaultEnv(5050)
)
var (
	DEFAULT_SELECTS_COLUMNS      = DefaultSelects([]types.SelectQuerySQLOperation{})
	DEFAULT_SELECTS_JOIN_COLUMNS = DefaultSelectsJoin([]types.SelectJoinQuerySQLOperation{})
)
