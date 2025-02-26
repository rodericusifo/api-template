// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package getter

import (
	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/pkg/constant"
	"api-gateway/internal/pkg/types"
)

func GetRequestUser(c *fiber.Ctx) *types.RequestUser {
	return c.Locals(constant.CONTEXT_KEY_REQUEST_USER).(*types.RequestUser)
}
