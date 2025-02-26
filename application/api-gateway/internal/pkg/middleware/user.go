// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package middleware

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"

	"api-gateway/internal/pkg/constant"
	"api-gateway/internal/pkg/types"
	"api-gateway/internal/pkg/util/validator"

	external_client_authservice_http_grpc_handler_user "api-gateway/external/client/auth-service/http/grpc/handler/user"
	pb_user "api-gateway/external/client/auth-service/http/grpc/proto/user"
)

func APIUser(userClientHandler external_client_authservice_http_grpc_handler_user.IUserClientHandler) fiber.Handler {
	return func(c *fiber.Ctx) error {
		claims := c.Locals(constant.CONTEXT_KEY_USER).(*jwt.Token).Claims
		user, ok := claims.(*types.JwtCustomClaims)
		if !ok {
			return fiber.NewError(fiber.StatusUnprocessableEntity, fmt.Sprintf("invalid claims type. correct type: %T", claims))
		}

		userRes, err := userClientHandler.GetUser(context.Background(), &pb_user.GetUserRequest{
			XID: user.XID,
		})
		if err != nil {
			return err
		}

		roleStatus := pb_user.RoleStatus_name[int32(userRes.GetData().GetRole().GetStatus())]
		reqUser := new(types.RequestUser)
		reqUser = &types.RequestUser{
			ID:    userRes.GetData().GetID(),
			XID:   user.XID,
			Name:  userRes.GetData().GetName(),
			Email: userRes.GetData().GetEmail(),
			Role: types.RequestRole{
				ID:     userRes.GetData().GetRole().GetID(),
				Name:   userRes.GetData().GetRole().GetName(),
				Slug:   userRes.GetData().GetRole().GetSlug(),
				Status: constant.RoleStatus(roleStatus),
			},
		}
		if err := validator.ValidateRequestUser(reqUser); err != nil {
			return err
		}

		c.Locals(constant.CONTEXT_KEY_REQUEST_USER, reqUser)
		return c.Next()
	}
}
