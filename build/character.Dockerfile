# Build application
FROM golang:1.19 AS build
WORKDIR /src
ENV CGO_ENABLED=0
COPY ./ ./
RUN go mod download
ARG APP_VERSION=v0.0.1
RUN go build \
	-ldflags="-X 'github.com/zedGGs/grpc-multiplayer-world-fantasy/pkg/config/default.Version=${APP_VERSION}'" \
	-o /out/character ./cmd/character

# Run server
FROM alpine:3.15.0
WORKDIR /app
COPY --from=build /out/character ./
EXPOSE 8081
ENTRYPOINT [ "./character" ]
