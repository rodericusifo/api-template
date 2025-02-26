// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package interceptor

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"api-gateway/internal/pkg/config"
	"api-gateway/internal/pkg/constant"
)

func GRPCError(
	ctx context.Context,
	method string,
	req interface{},
	reply interface{},
	cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker,
	opts ...grpc.CallOption,
) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"type":   fmt.Sprintf("%T", err),
			"detail": err,
		}).Errorln("[GRPC ERROR]")
		st, ok := status.FromError(err)
		if ok {
			v, exist := constant.ErrorGRPCToAPI[constant.ErrorGRPC{
				Code:    st.Code(),
				Message: st.Message(),
			}]
			if exist {
				return fiber.NewError(v.Code, v.Message)
			}
			return fiber.NewError(constant.ErrorCodeGRPCToAPI[st.Code()], st.Message())
		}
		return fiber.NewError(fiber.StatusInternalServerError, "Internal Server Error")
	}
	return nil
}
