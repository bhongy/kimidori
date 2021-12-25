# ingress-envoy

## Setup

On local machine, generate self-signed certificate for localhost (requires `openssl`). `ingress-envoy` needs to mount the certificate and the key as secrets in order to serve requests over TLS.

```sh
bash gen-localhost-tls.sh
```

```sh
docker-compose up --build
```

Then visit [https://localhost:8020](https://localhost:8020)
