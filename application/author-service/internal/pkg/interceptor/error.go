// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package interceptor

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"author-service/internal/pkg/config"
)

func UnaryGRPCError(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"type":   fmt.Sprintf("%T", err),
			"detail": err,
		}).Errorln("[UNARY GRPC ERROR]")
		st, ok := status.FromError(err)
		if ok {
			return nil, st.Err()
		}
		me, ok := err.(*json.MarshalerError)
		if ok {
			return nil, status.Error(codes.InvalidArgument, me.Error())
		}
		re, ok := err.(runtime.Error)
		if ok {
			return nil, status.Error(codes.InvalidArgument, re.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}
	return resp, nil
}
