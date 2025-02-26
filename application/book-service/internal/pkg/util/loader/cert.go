// Copyright 2025 Rodericus Ifo Krista
// SPDX-License-Identifier: MIT

package loader

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func LoadCert(certFile, keyFile string) (tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return tls.Certificate{}, err
	}
	return cert, nil
}

func LoadCA(caFile string) (*x509.CertPool, error) {
	caCert, err := os.ReadFile(caFile)
	if err != nil {
		return nil, err
	}
	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)
	return caPool, nil
}
