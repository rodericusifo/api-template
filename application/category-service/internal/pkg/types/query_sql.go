// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package types

import (
	"gorm.io/gorm/clause"
)

type SelectQuerySQLOperation struct {
	Field    string
	Alias    string
	Function string
}

type SelectJoinQuerySQLOperation struct {
	Field string
}

type SearchQuerySQLOperation struct {
	Field    string
	Operator string
	Value    any
}

type JoinQuerySQLOperation struct {
	Relation string
	Selects  []SelectJoinQuerySQLOperation
	Searches [][]SearchQuerySQLOperation
}

type InnerJoinQuerySQLOperation struct {
	Relation string
	Selects  []SelectJoinQuerySQLOperation
	Searches [][]SearchQuerySQLOperation
}

type OrderQuerySQLOperation struct {
	Field      string
	Descending bool
}

type GroupQuerySQLOperation struct {
	Field    string
	Function string
}

type ClauseExpression = clause.Expression

type QuerySQL struct {
	Selects     []SelectQuerySQLOperation
	Searches    [][]SearchQuerySQLOperation
	Joins       []JoinQuerySQLOperation
	InnerJoins  []InnerJoinQuerySQLOperation
	Orders      []OrderQuerySQLOperation
	Groups      []GroupQuerySQLOperation
	Clauses     []ClauseExpression
	Distinct    bool
	WithDeleted bool
	Limit       int
	Offset      int
}
