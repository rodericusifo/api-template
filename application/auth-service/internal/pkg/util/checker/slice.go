// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package checker

func CheckSliceContain[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
