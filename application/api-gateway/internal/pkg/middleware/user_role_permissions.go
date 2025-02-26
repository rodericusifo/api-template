// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package middleware

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"api-gateway/internal/pkg/util/getter"

	external_client_authservice_http_grpc_handler_user "api-gateway/external/client/auth-service/http/grpc/handler/user"
	pb_user "api-gateway/external/client/auth-service/http/grpc/proto/user"
)

func APIUserRolePermissions(userClientHandler external_client_authservice_http_grpc_handler_user.IUserClientHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		reqUser := getter.GetRequestUser(c)

		if reqUser.Role.Slug == "super_admin" {
			return c.Next()
		}

		path := c.Path()

		err := userClientHandler.ValidateUserRolePermissions(context.Background(), &pb_user.ValidateUserRolePermissionsRequest{
			RoleID: reqUser.Role.ID,
			Path:   path,
		})
		if err != nil {
			return err
		}

		return c.Next()
	}
}
