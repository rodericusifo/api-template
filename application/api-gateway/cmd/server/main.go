// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/sirupsen/logrus"

	"api-gateway/internal/pkg/config"
	"api-gateway/internal/pkg/constant"
	"api-gateway/internal/pkg/middleware"

	external_client_authservice_http_grpc_handler_auth "api-gateway/external/client/auth-service/http/grpc/handler/auth"
	external_client_authservice_http_grpc_handler_user "api-gateway/external/client/auth-service/http/grpc/handler/user"
	external_client_authorservice_http_grpc_handler_author "api-gateway/external/client/author-service/http/grpc/handler/author"
	external_client_bookservice_http_grpc_handler_book "api-gateway/external/client/book-service/http/grpc/handler/book"
	external_client_categoryservice_http_grpc_handler_category "api-gateway/external/client/category-service/http/grpc/handler/category"
	internal_domain_core_auth_controller_http_rest "api-gateway/internal/domain/core/auth/controller/http/rest"
	internal_domain_core_author_controller_http_rest "api-gateway/internal/domain/core/author/controller/http/rest"
	internal_domain_core_book_controller_http_rest "api-gateway/internal/domain/core/book/controller/http/rest"
	internal_domain_core_category_controller_http_rest "api-gateway/internal/domain/core/category/controller/http/rest"
)

func init() {
	config.ConfigureLog()
	config.ConfigureVars()
	config.ConfigureAuth()
}

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.RESTError,
		ServerHeader: "Fiber",
		AppName:      fmt.Sprintf("%s v%s", os.Getenv("NAME"), os.Getenv("VERSION")),
	})

	app.Use(
		requestid.New(),
		logger.New(logger.Config{
			Format: "[${time}] ${pid} | ${locals:requestid} | ${status} | ${latency} | ${method} | ${path}\n",
		}),
		recover.New(recover.Config{
			EnableStackTrace:  true,
			StackTraceHandler: middleware.APIRecover,
		}),
		cors.New(cors.Config{
			AllowMethods: "GET,POST,DELETE,PUT",
		}),
	)

	apiVersion := "/v" + strings.Split(os.Getenv("VERSION"), ".")[0]

	// AUTH SERVICE
	authClientHandler := external_client_authservice_http_grpc_handler_auth.InitAuthClientHandler()
	defer authClientHandler.CloseConnection()
	userClientHandler := external_client_authservice_http_grpc_handler_user.InitUserClientHandler()
	defer userClientHandler.CloseConnection()

	// AUTHOR SERVICE
	authorClientHandler := external_client_authorservice_http_grpc_handler_author.InitAuthorClientHandler()
	defer authorClientHandler.CloseConnection()

	// CATEGORY SERVICE
	categoryClientHandler := external_client_categoryservice_http_grpc_handler_category.InitCategoryClientHandler()
	defer categoryClientHandler.CloseConnection()

	// BOOK SERVICE
	bookClientHandler := external_client_bookservice_http_grpc_handler_book.InitbookClientHandler()
	defer bookClientHandler.CloseConnection()

	internal_domain_core_auth_controller_http_rest.InitREST(app.Group(apiVersion), authClientHandler)
	internal_domain_core_author_controller_http_rest.InitREST(app.Group(apiVersion), userClientHandler, authorClientHandler, bookClientHandler)
	internal_domain_core_category_controller_http_rest.InitREST(app.Group(apiVersion), userClientHandler, categoryClientHandler, bookClientHandler)
	internal_domain_core_book_controller_http_rest.InitREST(app.Group(apiVersion), userClientHandler, bookClientHandler, authorClientHandler, categoryClientHandler)

	go func() {
		err := app.Listen(fmt.Sprintf(":%d", func() int {
			serverPort := config.GetVarsConfig().ServerPort
			if serverPort != 0 {
				return serverPort
			} else {
				return constant.DEFAULT_ENV_SERVER_PORT.(int)
			}
		}()))
		if err != nil {
			config.GetLogConfig().WithFields(logrus.Fields{
				"message": "error starting server",
				"detail":  err,
			}).Fatal("[MAIN]")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	config.GetLogConfig().WithFields(logrus.Fields{
		"message": "shutting down server...",
	}).Infoln("[MAIN]")

	if err := app.Shutdown(); err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "server forced to shutdown",
			"detail":  err,
		}).Fatal("[MAIN]")
	}

	config.GetLogConfig().WithFields(logrus.Fields{
		"message": "server shutdown gracefully",
	}).Infoln("[MAIN]")
}
