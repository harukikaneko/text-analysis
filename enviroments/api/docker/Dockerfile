FROM rust:1.63.0 AS builder
ADD . ./
RUN --mount=type=cache,target=/usr/local/cargo/registry \
    --mount=type=cache,target=/target

RUN cargo test && \
    cargo build --release

RUN strip /target/release/text-analysis-api && \
    cp /target/release/text-analysis-api /text-analysis-api

FROM debian:bullseye-slim
RUN apt-get update && \
    apt-get install -y libpq-dev ca-certificates openssl && \ 
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /text-analysis-api /

ENTRYPOINT ["/text-analysis-api"]