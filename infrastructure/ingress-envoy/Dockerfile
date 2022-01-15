FROM envoyproxy/envoy-dev:latest

RUN apt-get update \
    && apt-get install --no-install-recommends -y curl \
    && apt-get autoremove -y \
    && apt-get clean \
    && rm -rf /tmp/* /var/tmp/* /var/lib/apt/lists/*
