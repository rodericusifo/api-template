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

	"author-service/internal/pkg/config"
	"author-service/internal/pkg/constant"
	"author-service/internal/pkg/interceptor"
	"author-service/internal/pkg/util/loader"

	internal_domain_core_author_controller_http_grpc_handler "author-service/internal/domain/core/author/controller/http/grpc/handler"
	pb_author "author-service/internal/proto/author"
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
}

func main() {
	certificate, err := loader.LoadCert("/certs/author-service.crt", "/certs/author-service.key")
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

	authorHandler := internal_domain_core_author_controller_http_grpc_handler.InitAuthorHandler()
	pb_author.RegisterAuthorHandlerServer(server, authorHandler)

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
