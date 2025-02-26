// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package auth

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

	pb_auth "api-gateway/external/client/auth-service/http/grpc/proto/auth"
)

type IAuthClientHandler interface {
	RegisterAuth(ctx context.Context, in *pb_auth.RegisterAuthRequest) error
	LoginAuth(ctx context.Context, in *pb_auth.LoginAuthRequest) (*pb_auth.LoginAuthResponse, error)

	CloseConnection()
}

type AuthClientHandler struct {
	Conn   *grpc.ClientConn
	Client pb_auth.AuthHandlerClient
}

func InitAuthClientHandler() IAuthClientHandler {
	certificate, err := loader.LoadCert("/certs/api-gateway.crt", "/certs/api-gateway.key")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to load server certificate and key",
			"detail":  err,
		}).Fatal("[INIT AUTH CLIENT HANDLER]")
	}

	rootCAs, err := loader.LoadCA("/certs/ca.crt")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to read CA certificate",
			"detail":  err,
		}).Fatal("[INIT AUTH CLIENT HANDLER]")
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
		}).Fatal("[INIT AUTH CLIENT HANDLER]")
	}

	client := pb_auth.NewAuthHandlerClient(conn)

	return &AuthClientHandler{
		Conn:   conn,
		Client: client,
	}
}

func (h *AuthClientHandler) CloseConnection() {
	h.Conn.Close()
}
