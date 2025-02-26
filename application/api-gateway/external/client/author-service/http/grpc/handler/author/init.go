// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package author

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

	pb_author "api-gateway/external/client/author-service/http/grpc/proto/author"
)

type IAuthorClientHandler interface {
	CreateAuthor(ctx context.Context, in *pb_author.CreateAuthorRequest) error
	GetAuthor(ctx context.Context, in *pb_author.GetAuthorRequest) (*pb_author.GetAuthorResponse, error)
	GetAuthors(ctx context.Context, in *pb_author.GetAuthorsRequest) (*pb_author.GetAuthorsResponse, error)
	UpdateAuthor(ctx context.Context, in *pb_author.UpdateAuthorRequest) error
	DeleteAuthor(ctx context.Context, in *pb_author.DeleteAuthorRequest) error

	CloseConnection()
}

type AuthorClientHandler struct {
	Conn   *grpc.ClientConn
	Client pb_author.AuthorHandlerClient
}

func InitAuthorClientHandler() IAuthorClientHandler {
	certificate, err := loader.LoadCert("/certs/api-gateway.crt", "/certs/api-gateway.key")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to load server certificate and key",
			"detail":  err,
		}).Fatal("[INIT AUTHOR CLIENT HANDLER]")
	}

	rootCAs, err := loader.LoadCA("/certs/ca.crt")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to read CA certificate",
			"detail":  err,
		}).Fatal("[INIT AUTHOR CLIENT HANDLER]")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      rootCAs,
	})

	target := fmt.Sprintf("%s:%d", config.GetVarsConfig().AuthorServiceServerHost, config.GetVarsConfig().AuthorServiceServerPort)
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(interceptor.GRPCError),
	)
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "Failed to connect",
			"detail":  err,
		}).Fatal("[INIT AUTHOR CLIENT HANDLER]")
	}

	client := pb_author.NewAuthorHandlerClient(conn)

	return &AuthorClientHandler{
		Conn:   conn,
		Client: client,
	}
}

func (h *AuthorClientHandler) CloseConnection() {
	h.Conn.Close()
}
