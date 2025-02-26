// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package user

import (
	"context"
	"crypto/tls"
	"fmt"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"api-gateway/external/pkg/interceptor"
	"api-gateway/external/pkg/util/loader"
	"api-gateway/internal/pkg/config"

	pb_user "api-gateway/external/client/auth-service/http/grpc/proto/user"
)

type IUserClientHandler interface {
	GetUser(ctx context.Context, in *pb_user.GetUserRequest) (*pb_user.GetUserResponse, error)

	ValidateUserRolePermissions(ctx context.Context, in *pb_user.ValidateUserRolePermissionsRequest) error

	CloseConnection()
}

type UserClientHandler struct {
	Conn   *grpc.ClientConn
	Client pb_user.UserHandlerClient
}

func InitUserClientHandler() IUserClientHandler {
	certificate, err := loader.LoadCert("/certs/api-gateway.crt", "/certs/api-gateway.key")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to load server certificate and key",
			"detail":  err,
		}).Fatal("[INIT USER CLIENT HANDLER]")
	}

	rootCAs, err := loader.LoadCA("/certs/ca.crt")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to read CA certificate",
			"detail":  err,
		}).Fatal("[INIT USER CLIENT HANDLER]")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      rootCAs,
	})

	target := fmt.Sprintf("%s:%d", config.GetVarsConfig().AuthServiceServerHost, config.GetVarsConfig().AuthServiceServerPort)
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(interceptor.GRPCError),
	)
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "Failed to connect",
			"detail":  err,
		}).Fatal("[INIT USER CLIENT HANDLER]")
	}

	client := pb_user.NewUserHandlerClient(conn)

	return &UserClientHandler{
		Conn:   conn,
		Client: client,
	}
}

func (h *UserClientHandler) CloseConnection() {
	h.Conn.Close()
}
