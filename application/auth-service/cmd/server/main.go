// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package main

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"

	"auth-service/internal/pkg/config"
	"auth-service/internal/pkg/constant"
	"auth-service/internal/pkg/interceptor"
	"auth-service/internal/pkg/util/loader"
	"auth-service/internal/pkg/util/seeder"

	internal_domain_core_auth_controller_http_grpc_handler "auth-service/internal/domain/core/auth/controller/http/grpc/handler"
	internal_domain_core_user_controller_http_grpc_handler "auth-service/internal/domain/core/user/controller/http/grpc/handler"
	pb_auth "auth-service/internal/proto/auth"
	pb_user "auth-service/internal/proto/user"
)

func init() {
	config.ConfigureLog()
	config.ConfigureVars()
	config.ConfigureDatabaseCache(constant.REDIS, config.DatabaseCacheDialects{
		Redis: config.DatabaseCacheConfig[config.DatabaseCacheRedisConfigOptions]{
			Enabled: config.GetVarsConfig().DatabaseCacheRedisEnabled,
			Primary: config.DatabaseCacheRedisConfigOptions{
				URL: config.GetVarsConfig().DatabaseCacheRedisPrimaryURL,
			},
			Replica: config.DatabaseCacheRedisConfigOptions{
				URL: config.GetVarsConfig().DatabaseCacheRedisReplicaURL,
			},
		},
	})
	config.ConfigureDatabaseSQL(constant.POSTGRES, config.DatabaseSQLDialects{
		Postgres: config.DatabaseSQLConfig[config.DatabaseSQLPostgresConfigOptions]{
			Enabled: config.GetVarsConfig().DatabaseSQLPostgresEnabled,
			Primary: config.DatabaseSQLPostgresConfigOptions{
				DSN:               config.GetVarsConfig().DatabaseSQLPostgresPrimaryDSN,
				ConnectionTimeout: config.GetVarsConfig().DatabaseSQLPostgresPrimaryConnectionTimeout,
				MaxIdleConnection: config.GetVarsConfig().DatabaseSQLPostgresPrimaryMaxIdleConnection,
				MaxOpenConnection: config.GetVarsConfig().DatabaseSQLPostgresPrimaryMaxOpenConnection,
				DebugMode:         config.GetVarsConfig().DatabaseSQLPostgresPrimaryDebugMode,
			},
			Replica: config.DatabaseSQLPostgresConfigOptions{
				DSN:               config.GetVarsConfig().DatabaseSQLPostgresReplicaDSN,
				ConnectionTimeout: config.GetVarsConfig().DatabaseSQLPostgresReplicaConnectionTimeout,
				MaxIdleConnection: config.GetVarsConfig().DatabaseSQLPostgresReplicaMaxIdleConnection,
				MaxOpenConnection: config.GetVarsConfig().DatabaseSQLPostgresReplicaMaxOpenConnection,
				DebugMode:         config.GetVarsConfig().DatabaseSQLPostgresReplicaDebugMode,
			},
		},
	})

	seeder.SeedRolesPermissions(constant.POSTGRES)
}

func main() {
	certificate, err := loader.LoadCert("/certs/auth-service.crt", "/certs/auth-service.key")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to load server certificate and key",
			"detail":  err,
		}).Fatal("[MAIN]")
	}

	rootCAs, err := loader.LoadCA("/certs/ca.crt")
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to read CA certificate",
			"detail":  err,
		}).Fatal("[MAIN]")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    rootCAs,
		ClientAuth:   tls.RequireAndVerifyClientCert,
	})

	server := grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			interceptor.UnaryGRPCLogger(),
			interceptor.UnaryGRPCRecover,
			interceptor.UnaryGRPCError,
		),
	)

	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(server, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	authHandler := internal_domain_core_auth_controller_http_grpc_handler.InitAuthHandler()
	pb_auth.RegisterAuthHandlerServer(server, authHandler)

	userHandler := internal_domain_core_user_controller_http_grpc_handler.InitUserHandler()
	pb_user.RegisterUserHandlerServer(server, userHandler)

	port := fmt.Sprintf(":%d", func() int {
		serverPort := config.GetVarsConfig().ServerPort
		if serverPort != 0 {
			return serverPort
		} else {
			return constant.DEFAULT_ENV_SERVER_PORT.(int)
		}
	}())
	listener, err := net.Listen("tcp", port)
	if err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to listen",
			"detail":  err,
		}).Fatal("[MAIN]")
	}

	config.GetLogConfig().WithFields(logrus.Fields{
		"message": fmt.Sprintf("🚀  gRPC server started successfully on port %s", port),
	}).Infoln("[MAIN]")
	if err := server.Serve(listener); err != nil {
		config.GetLogConfig().WithFields(logrus.Fields{
			"message": "failed to serve",
			"detail":  err,
		}).Fatal("[MAIN]")
	}
}
