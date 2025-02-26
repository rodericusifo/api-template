// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package constant

import (
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/codes"
)

type ErrorGRPC struct {
	Code    codes.Code
	Message string
}

type ErrorAPI struct {
	Code    int
	Message string
}

var (
	ErrorGRPCToAPI = map[ErrorGRPC]ErrorAPI{
		{
			Code:    codes.NotFound,
			Message: "user not found",
		}: {
			Code:    fiber.StatusUnauthorized,
			Message: "you are not registered. please register first",
		},
		{
			Code:    codes.AlreadyExists,
			Message: "user already exist",
		}: {
			Code:    fiber.StatusConflict,
			Message: "you have already registered. you can log in now",
		},
		{
			Code:    codes.NotFound,
			Message: "borrow record not found",
		}: {
			Code:    fiber.StatusNotFound,
			Message: "you have not borrowed this book",
		},
		{
			Code:    codes.NotFound,
			Message: "role permissions not found",
		}: {
			Code:    fiber.StatusUnauthorized,
			Message: "you do not have permission to access this resource",
		},
		{
			Code:    codes.NotFound,
			Message: "role permission not found",
		}: {
			Code:    fiber.StatusUnauthorized,
			Message: "you do not have permission to access this resource",
		},
		{
			Code:    codes.PermissionDenied,
			Message: "role is inactive",
		}: {
			Code:    fiber.StatusUnauthorized,
			Message: "You cannot perform this action because the required role is inactive",
		},
		{
			Code:    codes.PermissionDenied,
			Message: "permission is inactive",
		}: {
			Code:    fiber.StatusUnauthorized,
			Message: "You cannot perform this action because the required permission is inactive",
		},
	}
)
