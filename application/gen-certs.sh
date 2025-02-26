#!/bin/bash
# Copyright 2025 Rodericus Ifo Krista
# SPDX-License-Identifier: MIT

# Set the certificates directory
CERTS_DIR="certs"

# Paths to certifacates directory of services
API_GATEWAY_CERTS_DIR="api-gateway/$CERTS_DIR"
AUTH_SERVICE_CERTS_DIR="auth-service/$CERTS_DIR"
AUTHOR_SERVICE_CERTS_DIR="author-service/$CERTS_DIR"
CATEGORY_SERVICE_CERTS_DIR="category-service/$CERTS_DIR"
BOOK_SERVICE_CERTS_DIR="book-service/$CERTS_DIR"

# Paths to the SAN configuration files
API_GATEWAY_CNF="./$API_GATEWAY_CERTS_DIR/api-gateway-san.cnf"
AUTH_SERVICE_CNF="./$AUTH_SERVICE_CERTS_DIR/auth-service-san.cnf"
AUTHOR_SERVICE_CNF="./$AUTHOR_SERVICE_CERTS_DIR/author-service-san.cnf"
CATEGORY_SERVICE_CNF="./$CATEGORY_SERVICE_CERTS_DIR/category-service-san.cnf"
BOOK_SERVICE_CNF="./$BOOK_SERVICE_CERTS_DIR/book-service-san.cnf"

# ========== Generate Certificate Authority (CA) ==========
echo "Generating CA Key and Certificate..."
mkdir -p ./$CERTS_DIR
openssl genrsa -out ./$CERTS_DIR/ca.key 2048
MSYS_NO_PATHCONV=1 openssl req -x509 -new -nodes -key ./$CERTS_DIR/ca.key -sha256 -days 365 \
    -subj "/C=US/ST=California/L=San Francisco/O=My Organization/CN=My CA" \
    -out ./$CERTS_DIR/ca.crt

# ========== Copy Certificate Authority (CA) to Services Directory ==========
echo "Copy CA Key and Certificate to Services..."
dests="$API_GATEWAY_CERTS_DIR $AUTH_SERVICE_CERTS_DIR $AUTHOR_SERVICE_CERTS_DIR $CATEGORY_SERVICE_CERTS_DIR $BOOK_SERVICE_CERTS_DIR"
for dest in $dests; do
  cp ./$CERTS_DIR/ca.key "./$dest/ca.key";
  cp ./$CERTS_DIR/ca.crt "./$dest/ca.crt";
done
rm -rf ./$CERTS_DIR

# ========== Generate API Gateway Certificate and Key ==========
echo "Generating Client Certificate (api-gateway)..."
openssl genrsa -out ./$API_GATEWAY_CERTS_DIR/api-gateway.key 2048
openssl req -new -key ./$API_GATEWAY_CERTS_DIR/api-gateway.key -out ./$API_GATEWAY_CERTS_DIR/api-gateway.csr \
    -config $API_GATEWAY_CNF
openssl x509 -req -in ./$API_GATEWAY_CERTS_DIR/api-gateway.csr -CA ./$API_GATEWAY_CERTS_DIR/ca.crt -CAkey ./$API_GATEWAY_CERTS_DIR/ca.key \
    -CAcreateserial -out ./$API_GATEWAY_CERTS_DIR/api-gateway.crt -days 365 -sha256 \
    -extfile $API_GATEWAY_CNF -extensions v3_req

# ========== Generate Auth Service Certificate and Key ==========
echo "Generating Server Certificate (auth-service)..."
openssl genrsa -out ./$AUTH_SERVICE_CERTS_DIR/auth-service.key 2048
openssl req -new -key ./$AUTH_SERVICE_CERTS_DIR/auth-service.key -out ./$AUTH_SERVICE_CERTS_DIR/auth-service.csr \
    -config $AUTH_SERVICE_CNF
openssl x509 -req -in ./$AUTH_SERVICE_CERTS_DIR/auth-service.csr -CA ./$AUTH_SERVICE_CERTS_DIR/ca.crt -CAkey ./$AUTH_SERVICE_CERTS_DIR/ca.key \
    -CAcreateserial -out ./$AUTH_SERVICE_CERTS_DIR/auth-service.crt -days 365 -sha256 \
    -extfile $AUTH_SERVICE_CNF -extensions v3_req

# ========== Generate Author Service Certificate and Key ==========
echo "Generating Server Certificate (author-service)..."
openssl genrsa -out ./$AUTHOR_SERVICE_CERTS_DIR/author-service.key 2048
openssl req -new -key ./$AUTHOR_SERVICE_CERTS_DIR/author-service.key -out ./$AUTHOR_SERVICE_CERTS_DIR/author-service.csr \
    -config $AUTHOR_SERVICE_CNF
openssl x509 -req -in ./$AUTHOR_SERVICE_CERTS_DIR/author-service.csr -CA ./$AUTHOR_SERVICE_CERTS_DIR/ca.crt -CAkey ./$AUTHOR_SERVICE_CERTS_DIR/ca.key \
    -CAcreateserial -out ./$AUTHOR_SERVICE_CERTS_DIR/author-service.crt -days 365 -sha256 \
    -extfile $AUTHOR_SERVICE_CNF -extensions v3_req

# ========== Generate Category Service Certificate and Key ==========
echo "Generating Server Certificate (category-service)..."
openssl genrsa -out ./$CATEGORY_SERVICE_CERTS_DIR/category-service.key 2048
openssl req -new -key ./$CATEGORY_SERVICE_CERTS_DIR/category-service.key -out ./$CATEGORY_SERVICE_CERTS_DIR/category-service.csr \
    -config $CATEGORY_SERVICE_CNF
openssl x509 -req -in ./$CATEGORY_SERVICE_CERTS_DIR/category-service.csr -CA ./$CATEGORY_SERVICE_CERTS_DIR/ca.crt -CAkey ./$CATEGORY_SERVICE_CERTS_DIR/ca.key \
    -CAcreateserial -out ./$CATEGORY_SERVICE_CERTS_DIR/category-service.crt -days 365 -sha256 \
    -extfile $CATEGORY_SERVICE_CNF -extensions v3_req

# ========== Generate Book Service Certificate and Key ==========
echo "Generating Server Certificate (book-service)..."
openssl genrsa -out ./$BOOK_SERVICE_CERTS_DIR/book-service.key 2048
openssl req -new -key ./$BOOK_SERVICE_CERTS_DIR/book-service.key -out ./$BOOK_SERVICE_CERTS_DIR/book-service.csr \
    -config $BOOK_SERVICE_CNF
openssl x509 -req -in ./$BOOK_SERVICE_CERTS_DIR/book-service.csr -CA ./$BOOK_SERVICE_CERTS_DIR/ca.crt -CAkey ./$BOOK_SERVICE_CERTS_DIR/ca.key \
    -CAcreateserial -out ./$BOOK_SERVICE_CERTS_DIR/book-service.crt -days 365 -sha256 \
    -extfile $BOOK_SERVICE_CNF -extensions v3_req

# ========== Verify Certificates ==========
echo "Verifying Client Certificate (api-gateway)"
openssl verify -CAfile ./$API_GATEWAY_CERTS_DIR/ca.crt ./$API_GATEWAY_CERTS_DIR/api-gateway.crt

echo "Verifying Server Certificate (auth-service)"
openssl verify -CAfile ./$AUTH_SERVICE_CERTS_DIR/ca.crt ./$AUTH_SERVICE_CERTS_DIR/auth-service.crt

echo "Verifying Server Certificate (author-service)"
openssl verify -CAfile ./$AUTHOR_SERVICE_CERTS_DIR/ca.crt ./$AUTHOR_SERVICE_CERTS_DIR/author-service.crt

echo "Verifying Server Certificate (category-service)"
openssl verify -CAfile ./$CATEGORY_SERVICE_CERTS_DIR/ca.crt ./$CATEGORY_SERVICE_CERTS_DIR/category-service.crt

echo "Verifying Server Certificate (book-service)"
openssl verify -CAfile ./$BOOK_SERVICE_CERTS_DIR/ca.crt ./$BOOK_SERVICE_CERTS_DIR/book-service.crt