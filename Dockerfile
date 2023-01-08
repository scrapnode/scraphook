# syntax=docker/dockerfile:1
FROM golang:1.18-alpine as build
WORKDIR /app

RUN apk add build-base

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./bin/scraphook -buildvcs=false

FROM alpine:3
WORKDIR /app

# healthcheck
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.13 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

COPY --from=build /app/db ./db
COPY --from=build /app/configs.props.example ./secrets/configs.props
COPY --from=build /app/.version ./.version

COPY --from=build /app/docker-entrypoint.sh ./docker-entrypoint.sh
RUN chmod +x /app/docker-entrypoint.sh

COPY --from=build /app/bin/scraphook ./scraphook

EXPOSE 8080
EXPOSE 8081
ENTRYPOINT ["/app/docker-entrypoint.sh"]