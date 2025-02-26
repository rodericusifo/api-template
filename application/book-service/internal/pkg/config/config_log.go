// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package config

import (
	"github.com/sirupsen/logrus"
)

var (
	log *logrus.Logger
)

func ConfigureLog() {
	log = logrus.New()
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
		TimestampFormat:        "2006-01-02 15:04:05 MST",
	})
	log.SetReportCaller(true)
	log.WithFields(logrus.Fields{
		"message": "setting log success",
	}).Infoln("[CONFIGURE LOG]")
}

func GetLogConfig() *logrus.Logger {
	return log
}
