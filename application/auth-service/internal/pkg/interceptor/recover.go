// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package interceptor

import (
	"context"
	"runtime/debug"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"auth-service/internal/pkg/config"
)

func UnaryGRPCRecover(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": "recovered from panic",
				"detail":  r,
			}).Errorln("[UNARY GRPC RECOVER]")
			config.GetLogConfig().Errorf("%s", debug.Stack())
			err = status.Errorf(codes.Internal, "internal server error")
		}
	}()
	return handler(ctx, req)
}
