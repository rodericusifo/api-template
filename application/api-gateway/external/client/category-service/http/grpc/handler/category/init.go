// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package category

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

	pb_category "api-gateway/external/client/category-service/http/grpc/proto/category"
)

type ICategoryClientHandler interface {
	CreateCategory(ctx context.Context, in *pb_category.CreateCategoryRequest) error
	GetCategory(ctx context.Context, in *pb_category.GetCategoryRequest) (*pb_category.GetCategoryResponse, error)
	GetCategories(ctx context.Context, in *pb_category.GetCategoriesRequest) (*pb_category.GetCategoriesResponse, error)
	UpdateCategory(ctx context.Context, in *pb_category.UpdateCategoryRequest) error
	DeleteCategory(ctx context.Context, in *pb_category.DeleteCategoryRequest) error

	CloseConnection()
}

type CategoryClientHandler struct {
	Conn   *grpc.ClientConn
	Client pb_category.CategoryHandlerClient
}

func InitCategoryClientHandler() ICategoryClientHandler {
	certificate, err := loader.LoadCert("/certs/api-gateway.crt", "/certs/api-gateway.key")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to load server certificate and key",
			"detail":  err,
		}).Fatal("[INIT CATEGORY CLIENT HANDLER]")
	}

	rootCAs, err := loader.LoadCA("/certs/ca.crt")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to read CA certificate",
			"detail":  err,
		}).Fatal("[INIT CATEGORY CLIENT HANDLER]")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      rootCAs,
	})

	target := fmt.Sprintf("%s:%d", config.GetVarsConfig().CategoryServiceServerHost, config.GetVarsConfig().CategoryServiceServerPort)
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(interceptor.GRPCError),
	)
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "Failed to connect",
			"detail":  err,
		}).Fatal("[INIT CATEGORY CLIENT HANDLER]")
	}

	client := pb_category.NewCategoryHandlerClient(conn)

	return &CategoryClientHandler{
		Conn:   conn,
		Client: client,
	}
}

func (h *CategoryClientHandler) CloseConnection() {
	h.Conn.Close()
}
