// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package interceptor

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	colorReset  = "\u001b[0m"
	colorGreen  = "\u001b[92m"
	colorBlue   = "\u001b[94m"
	colorYellow = "\u001b[93m"
	colorRed    = "\u001b[91m"
)

func colorStatusCode(code codes.Code) string {
	switch {
	case code == codes.OK:
		return colorGreen
	case code == codes.Canceled || code == codes.InvalidArgument:
		return colorBlue
	case code == codes.NotFound || code == codes.AlreadyExists || code == codes.FailedPrecondition:
		return colorYellow
	case code == codes.Internal || code == codes.Unavailable || code == codes.Unknown:
		return colorRed
	default:
		return colorReset
	}
}

type customFormatter struct{}

func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := entry.Time.Format("2006-01-02 15:04:05 MST")
	pid := os.Getpid()
	requestID := entry.Data["requestid"]
	statusCode, _ := entry.Data["statuscode"].(codes.Code)
	latency := entry.Data["latency"]
	path := entry.Data["path"]

	colorStatusCode := colorStatusCode(statusCode)

	logLine := fmt.Sprintf(
		"[%s] %d | %s | %s | %v | %s\n",
		timestamp, pid, requestID, fmt.Sprintf("%s%d%s", colorStatusCode, statusCode, colorReset), latency, path,
	)
	return []byte(logLine), nil
}

func UnaryGRPCLogger() grpc.UnaryServerInterceptor {
	log := logrus.New()
	log.SetFormatter(&customFormatter{})

	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		start := time.Now()

		requestID := uuid.New().String()

		res, err := handler(ctx, req)

		grpcStatus := status.Convert(err)
		statusCode := grpcStatus.Code()
		latency := time.Since(start)

		log.WithFields(logrus.Fields{
			"requestid":  requestID,
			"statuscode": statusCode,
			"latency":    latency,
			"path":       info.FullMethod,
		}).Info("[UNARY GRPC LOGGER]")

		return res, err
	}
}
