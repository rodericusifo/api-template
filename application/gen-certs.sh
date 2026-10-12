#!/bin/bash
# Copyright 2025 Rodericus Ifo Krista
# SPDX-License-Identifier: MIT

# Set the certificates directory
CERTS_DIR="certs"

# Services to generate certificates for, read from stack.mk's APPLICATIONS
# variable (single source of truth for the service list).
# api-gateway gets a client certificate; the rest get server certificates.
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
STACK_MK="$SCRIPT_DIR/../stack.mk"

if [ ! -f "$STACK_MK" ]; then
    echo "Error: stack.mk not found at $STACK_MK" >&2
    exit 1
fi

applications_line=$(grep -E '^APPLICATIONS[[:space:]]*:=' "$STACK_MK" | sed -E 's/^APPLICATIONS[[:space:]]*:=[[:space:]]*//')

if [ -z "$applications_line" ]; then
    echo "Error: APPLICATIONS variable not found in $STACK_MK" >&2
    exit 1
fi

read -ra SERVICES <<< "$applications_line"

# ========== Generate Certificate Authority (CA) ==========
echo "Generating CA Key and Certificate..."
mkdir -p ./$CERTS_DIR
openssl genrsa -out ./$CERTS_DIR/ca.key 2048
MSYS_NO_PATHCONV=1 openssl req -x509 -new -nodes -key ./$CERTS_DIR/ca.key -sha256 -days 365 \
    -subj "/C=US/ST=California/L=San Francisco/O=My Organization/CN=My CA" \
    -out ./$CERTS_DIR/ca.crt

# ========== Copy Certificate Authority (CA) to Services Directory ==========
echo "Copy CA Key and Certificate to Services..."
for service in "${SERVICES[@]}"; do
    dest="$service/$CERTS_DIR"
    cp ./$CERTS_DIR/ca.key "./$dest/ca.key"
    cp ./$CERTS_DIR/ca.crt "./$dest/ca.crt"
done
rm -rf ./$CERTS_DIR

# ========== Generate Service Certificates and Keys ==========
for service in "${SERVICES[@]}"; do
    dest="$service/$CERTS_DIR"
    cnf="./$dest/service-san.cnf"

    echo "Generating Certificate ($service)..."

    # APPLICATION is read by service-san.cnf via OpenSSL's $ENV:: config
    # expansion, so the CN/SAN always match the service's own directory name
    # without duplicating it in a committed, per-service config file.
    export APPLICATION="$service"

    openssl genrsa -out "./$dest/service.key" 2048
    openssl req -new -key "./$dest/service.key" -out "./$dest/service.csr" \
        -config "$cnf"
    openssl x509 -req -in "./$dest/service.csr" -CA "./$dest/ca.crt" -CAkey "./$dest/ca.key" \
        -CAcreateserial -out "./$dest/service.crt" -days 365 -sha256 \
        -extfile "$cnf" -extensions v3_req
done

# ========== Verify Certificates ==========
for service in "${SERVICES[@]}"; do
    dest="$service/$CERTS_DIR"

    echo "Verifying Certificate ($service)"

    openssl verify -CAfile "./$dest/ca.crt" "./$dest/service.crt"
done
