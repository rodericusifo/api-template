// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package constant

import (
	"net/http"

	"google.golang.org/grpc/codes"
)

var (
	ErrorCodeGRPCToAPI = map[codes.Code]int{
		codes.OK:                 http.StatusOK,                  // 200
		codes.Canceled:           499,                            // Client Closed Request (Non-standard)
		codes.Unknown:            http.StatusInternalServerError, // 500
		codes.InvalidArgument:    http.StatusBadRequest,          // 400
		codes.DeadlineExceeded:   http.StatusGatewayTimeout,      // 504
		codes.NotFound:           http.StatusNotFound,            // 404
		codes.AlreadyExists:      http.StatusConflict,            // 409
		codes.PermissionDenied:   http.StatusForbidden,           // 403
		codes.Unauthenticated:    http.StatusUnauthorized,        // 401
		codes.ResourceExhausted:  http.StatusTooManyRequests,     // 429
		codes.FailedPrecondition: http.StatusBadRequest,          // 400
		codes.Aborted:            http.StatusConflict,            // 409
		codes.OutOfRange:         http.StatusBadRequest,          // 400
		codes.Unimplemented:      http.StatusNotImplemented,      // 501
		codes.Internal:           http.StatusInternalServerError, // 500
		codes.Unavailable:        http.StatusServiceUnavailable,  // 503
		codes.DataLoss:           http.StatusInternalServerError, // 500
	}
)
