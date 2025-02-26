// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package constant

type RoleStatus string
type PermissionStatus string

const (
	ROLE_ACTIVE   = RoleStatus("ACTIVE")
	ROLE_INACTIVE = RoleStatus("INACTIVE")
)
const (
	PERMISSION_ACTIVE   = PermissionStatus("ACTIVE")
	PERMISSION_INACTIVE = PermissionStatus("INACTIVE")
)
