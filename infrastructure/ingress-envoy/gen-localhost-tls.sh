#!/usr/bin/env bash

DIR=secrets

mkdir -p $DIR

# https://letsencrypt.org/docs/certificates-for-localhost/
openssl req -x509 -out "$DIR/localhost.crt" -keyout "$DIR/localhost.key" \
  -days 365 \
  -newkey rsa:4096 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
