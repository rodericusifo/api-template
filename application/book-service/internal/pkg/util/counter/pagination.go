// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package counter

import (
	"math"
)

func CountPaginationOffset(page, limit int) int {
	if page == 0 || limit == 0 {
		return 0
	}
	return (page - 1) * limit
}

func CountPaginationTotalPage(countDataPerPage, totalData int) int {
	if countDataPerPage == 0 || totalData == 0 {
		return 0
	}
	return int(math.Ceil(float64(totalData) / float64(countDataPerPage)))
}
