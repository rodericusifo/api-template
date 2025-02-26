// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package merger

import (
	"author-service/internal/pkg/util/checker"
)

func MergeSlices[T comparable](withUnique bool, slices ...[]T) []T {
	resultMerged := make([]T, 0)
	resultMergedAndUnique := make([]T, 0)

	for _, s := range slices {
		resultMerged = append(resultMerged, s...)
	}

	if !withUnique {
		return resultMerged
	}

	for _, v := range resultMerged {
		if !checker.CheckSliceContain(resultMergedAndUnique, v) {
			resultMergedAndUnique = append(resultMergedAndUnique, v)
		}
	}

	return resultMergedAndUnique
}
