// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"api-gateway/internal/pkg/constant"
)

type varsConfig *struct {
	// SERVER
	ServerPort                int    `mapstructure:"SERVER_PORT"`
	AuthServiceServerHost     string `mapstructure:"AUTH_SERVICE_SERVER_HOST"`
	AuthServiceServerPort     int    `mapstructure:"AUTH_SERVICE_SERVER_PORT"`
	AuthorServiceServerHost   string `mapstructure:"AUTHOR_SERVICE_SERVER_HOST"`
	AuthorServiceServerPort   int    `mapstructure:"AUTHOR_SERVICE_SERVER_PORT"`
	CategoryServiceServerHost string `mapstructure:"CATEGORY_SERVICE_SERVER_HOST"`
	CategoryServiceServerPort int    `mapstructure:"CATEGORY_SERVICE_SERVER_PORT"`
	BookServiceServerHost     string `mapstructure:"BOOK_SERVICE_SERVER_HOST"`
	BookServiceServerPort     int    `mapstructure:"BOOK_SERVICE_SERVER_PORT"`

	// JWT
	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

var (
	varsCfg varsConfig
)

func ConfigureVars() {
	switch mode := os.Getenv("MODE"); {
	case strings.Contains(string(constant.DEVELOPMENT), strings.ToLower(mode)):
		setEnv()
	default:
		GetLogConfig().WithFields(logrus.Fields{
			"message": fmt.Sprintf("MODE %s is not valid", mode),
		}).Panic("[MAIN]")
	}
}

func GetVarsConfig() varsConfig {
	return varsCfg
}

func setEnv() {
	env := os.Getenv("ENV")

	if env == "" {
		GetLogConfig().WithFields(logrus.Fields{
			"message": "ENV is not set",
		}).Panic("[CONFIGURE VARS]")
	}

	path := fmt.Sprintf("/env/%s.env", env)
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		GetLogConfig().WithFields(logrus.Fields{
			"message": "read env fail",
			"detail":  err,
		}).Panic("[CONFIGURE VARS]")
	}

	if err := viper.Unmarshal(&varsCfg); err != nil {
		GetLogConfig().WithFields(logrus.Fields{
			"message": "load env fail",
			"detail":  err,
		}).Panic("[CONFIGURE VARS]")
	}
	GetLogConfig().WithFields(logrus.Fields{
		"message": "load env success",
	}).Infoln("[CONFIGURE VARS]")
}
