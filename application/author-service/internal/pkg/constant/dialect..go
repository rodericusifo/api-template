// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package constant

type DialectDatabaseSQL string
type DialectDatabaseCache string

var (
	POSTGRES = DialectDatabaseSQL("postgres")
	MYSQL    = DialectDatabaseSQL("mysql")
)
var (
	REDIS = DialectDatabaseCache("redis")
)
