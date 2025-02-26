// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package builder

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"auth-service/internal/pkg/constant"
	"auth-service/internal/pkg/types"
	"auth-service/internal/pkg/util/merger"
)

func BuildQuerySQL(tableName string, db *gorm.DB, query *types.QuerySQL, dialect constant.DialectDatabaseSQL) *gorm.DB {
	q := db

	if len(query.Joins) > 0 {
		for _, join := range query.Joins {
			if len(join.Selects) > 0 || len(join.Searches) > 0 {
				qj := db.Begin()
				if len(join.Selects) > 0 {
					querySlice := buildSelectJoinQuerySQLSlice(join.Relation, merger.MergeSlices(true, join.Selects, constant.DEFAULT_SELECTS_JOIN_COLUMNS), dialect)
					if query.Distinct {
						qj = qj.Distinct(querySlice)
					} else {
						qj = qj.Select(querySlice)
					}
				}
				if len(join.Searches) > 0 {
					queryString, bindValues := buildWhereQuerySQLStringAndBindValues(join.Relation, join.Searches, dialect)
					qj = qj.Where(queryString, bindValues...)
				}
				qj.Commit()
				q = q.Joins(join.Relation, qj)
			} else {
				q = q.Joins(join.Relation)
			}
		}
	}
	if len(query.InnerJoins) > 0 {
		for _, innerJoin := range query.InnerJoins {
			if len(innerJoin.Selects) > 0 || len(innerJoin.Searches) > 0 {
				qj := db.Begin()
				if len(innerJoin.Selects) > 0 {
					querySlice := buildSelectJoinQuerySQLSlice(innerJoin.Relation, merger.MergeSlices(true, innerJoin.Selects, constant.DEFAULT_SELECTS_JOIN_COLUMNS), dialect)
					if query.Distinct {
						qj = qj.Distinct(querySlice)
					} else {
						qj = qj.Select(querySlice)
					}
				}
				if len(innerJoin.Searches) > 0 {
					queryString, bindValues := buildWhereQuerySQLStringAndBindValues(innerJoin.Relation, innerJoin.Searches, dialect)
					qj = qj.Where(queryString, bindValues...)
				}
				qj.Commit()
				q = q.InnerJoins(innerJoin.Relation, qj)
			} else {
				q = q.InnerJoins(innerJoin.Relation)
			}
		}
	}
	if len(query.Selects) > 0 {
		querySlice := buildSelectQuerySQLSlice(tableName, merger.MergeSlices(true, query.Selects, constant.DEFAULT_SELECTS_COLUMNS), dialect)
		if query.Distinct {
			q = q.Distinct(querySlice)
		} else {
			q = q.Select(querySlice)
		}
	}
	if len(query.Searches) > 0 {
		queryString, bindValues := buildWhereQuerySQLStringAndBindValues(tableName, query.Searches, dialect)
		q = q.Where(queryString, bindValues...)
	}
	if len(query.Orders) > 0 {
		queryString := buildOrderQuerySQLString(tableName, query.Orders, dialect)
		q = q.Order(queryString)
	}
	if len(query.Groups) > 0 {
		queryString := buildGroupQuerySQLString(tableName, query.Groups, dialect)
		q = q.Group(queryString)
	}
	if len(query.Clauses) > 0 {
		q = q.Clauses(query.Clauses...)
	}
	if query.WithDeleted {
		q = q.Unscoped()
	}
	if query.Limit != 0 {
		q = q.Limit(query.Limit)
	}
	if query.Offset != 0 {
		q = q.Offset(query.Offset)
	}

	return q
}

func buildSelectQuerySQLSlice(tableAlias string, selects []types.SelectQuerySQLOperation, dialect constant.DialectDatabaseSQL) []string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case constant.MYSQL:
		for _, s := range selects {
			fieldSelectStr := fmt.Sprintf("`%s`.`%s`", tableAlias, s.Field)
			if s.Function != "" {
				function := s.Function
				fieldStr := fmt.Sprintf("`%s`.`%s`", tableAlias, s.Field)
				fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
			}

			if s.Alias != "" {
				querySlice = append(querySlice, fmt.Sprintf("%s AS `%s`", fieldSelectStr, s.Alias))
			} else {
				querySlice = append(querySlice, fieldSelectStr)
			}
		}
	case constant.POSTGRES:
		for _, s := range selects {
			fieldSelectStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, s.Field)
			if s.Function != "" {
				function := s.Function
				fieldStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, s.Field)
				fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
			}

			if s.Alias != "" {
				querySlice = append(querySlice, fmt.Sprintf(`%s AS "%s"`, fieldSelectStr, s.Alias))
			} else {
				querySlice = append(querySlice, fieldSelectStr)
			}
		}
	}

	return querySlice
}

func buildSelectJoinQuerySQLSlice(tableAlias string, selects []types.SelectJoinQuerySQLOperation, dialect constant.DialectDatabaseSQL) []string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case constant.MYSQL:
		for _, s := range selects {
			fieldSelectStr := fmt.Sprintf("`%s`.`%s`", tableAlias, s.Field)

			querySlice = append(querySlice, fieldSelectStr)
		}
	case constant.POSTGRES:
		for _, s := range selects {
			fieldSelectStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, s.Field)

			querySlice = append(querySlice, fieldSelectStr)
		}
	}

	return querySlice
}

func buildWhereQuerySQLStringAndBindValues(tableAlias string, searches [][]types.SearchQuerySQLOperation, dialect constant.DialectDatabaseSQL) (string, []any) {
	queryString := ""
	bindValues := make([]any, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case constant.MYSQL:
		for indexOuter, searchOuter := range searches {
			if indexOuter > 0 {
				queryString += " OR "
			}
			for indexInner, searchInner := range searchOuter {
				if indexInner > 0 {
					queryString += " AND "
				}
				if searchInner.Value != nil {
					queryString += fmt.Sprintf("`%s`.`%s` %s ?", tableAlias, searchInner.Field, searchInner.Operator)
					bindValues = append(bindValues, searchInner.Value)
				} else {
					queryString += fmt.Sprintf("`%s`.`%s` %s", tableAlias, searchInner.Field, searchInner.Operator)
				}
			}
		}
	case constant.POSTGRES:
		for indexOuter, searchOuter := range searches {
			if indexOuter > 0 {
				queryString += " OR "
			}
			for indexInner, searchInner := range searchOuter {
				if indexInner > 0 {
					queryString += " AND "
				}
				if searchInner.Value != nil {
					queryString += fmt.Sprintf(`"%s"."%s" %s ?`, tableAlias, searchInner.Field, searchInner.Operator)
					bindValues = append(bindValues, searchInner.Value)
				} else {
					queryString += fmt.Sprintf(`"%s"."%s" %s`, tableAlias, searchInner.Field, searchInner.Operator)
				}
			}
		}
	}

	return queryString, bindValues
}

func buildOrderQuerySQLString(tableAlias string, orders []types.OrderQuerySQLOperation, dialect constant.DialectDatabaseSQL) string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case constant.MYSQL:
		for _, order := range orders {
			if order.Descending {
				querySlice = append(querySlice, fmt.Sprintf("`%s`.`%s` DESC", tableAlias, order.Field))
			} else {
				querySlice = append(querySlice, fmt.Sprintf("`%s`.`%s`", tableAlias, order.Field))
			}
		}
	case constant.POSTGRES:
		for _, order := range orders {
			if order.Descending {
				querySlice = append(querySlice, fmt.Sprintf(`"%s"."%s" DESC`, tableAlias, order.Field))
			} else {
				querySlice = append(querySlice, fmt.Sprintf(`"%s"."%s"`, tableAlias, order.Field))
			}
		}
	}

	return strings.Join(querySlice, ",")
}

func buildGroupQuerySQLString(tableAlias string, groups []types.GroupQuerySQLOperation, dialect constant.DialectDatabaseSQL) string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case constant.MYSQL:
		for _, group := range groups {
			fieldSelectStr := ""
			if group.Function != "" {
				function := group.Function
				fieldStr := fmt.Sprintf("`%s`.`%s`", tableAlias, group.Field)
				fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
			} else {
				fieldSelectStr = fmt.Sprintf("`%s`.`%s`", tableAlias, group.Field)
			}
			querySlice = append(querySlice, fieldSelectStr)
		}
	case constant.POSTGRES:
		for _, group := range groups {
			fieldSelectStr := ""
			if group.Function != "" {
				function := group.Function
				fieldStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, group.Field)
				fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
			} else {
				fieldSelectStr = fmt.Sprintf(`"%s"."%s"`, tableAlias, group.Field)
			}
			querySlice = append(querySlice, fieldSelectStr)
		}
	}

	return strings.Join(querySlice, ",")
}
