// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package book

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

	pb_book "api-gateway/external/client/book-service/http/grpc/proto/book"
)

type IBookClientHandler interface {
	CreateBook(ctx context.Context, in *pb_book.CreateBookRequest) error
	GetBook(ctx context.Context, in *pb_book.GetBookRequest) (*pb_book.GetBookResponse, error)
	GetBooks(ctx context.Context, in *pb_book.GetBooksRequest) (*pb_book.GetBooksResponse, error)
	UpdateBook(ctx context.Context, in *pb_book.UpdateBookRequest) error
	DeleteBook(ctx context.Context, in *pb_book.DeleteBookRequest) error

	BorrowBook(ctx context.Context, in *pb_book.BorrowBookRequest) error
	ReturnBook(ctx context.Context, in *pb_book.ReturnBookRequest) error

	CloseConnection()
}

type BookClientHandler struct {
	Conn   *grpc.ClientConn
	Client pb_book.BookHandlerClient
}

func InitbookClientHandler() IBookClientHandler {
	certificate, err := loader.LoadCert("/certs/api-gateway.crt", "/certs/api-gateway.key")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to load server certificate and key",
			"detail":  err,
		}).Fatal("[INIT BOOK CLIENT HANDLER]")
	}

	rootCAs, err := loader.LoadCA("/certs/ca.crt")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to read CA certificate",
			"detail":  err,
		}).Fatal("[INIT BOOK CLIENT HANDLER]")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      rootCAs,
	})

	target := fmt.Sprintf("%s:%d", config.GetVarsConfig().BookServiceServerHost, config.GetVarsConfig().BookServiceServerPort)
	conn, err := grpc.NewClient(
		target,
		grpc.WithTransportCredentials(creds),
		grpc.WithUnaryInterceptor(interceptor.GRPCError),
	)
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "Failed to connect",
			"detail":  err,
		}).Fatal("[INIT BOOK CLIENT HANDLER]")
	}

	client := pb_book.NewBookHandlerClient(conn)

	return &BookClientHandler{
		Conn:   conn,
		Client: client,
	}
}

func (h *BookClientHandler) CloseConnection() {
	h.Conn.Close()
}
