// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package definer

func DefinePaginationPageLimit(page, limit *int) (pagePagination int, limitPagination int) {
	pagePagination = 1
	limitPagination = 10

	if page != nil {
		pagePagination = *page
	}
	if limit != nil {
		limitPagination = *limit
	}

	return
}
