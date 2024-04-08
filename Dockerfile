# syntax=docker/dockerfile:1

ARG GO_VERSION=1.22
ARG OS_NAME=alpine
ARG OS_VERSION=3.19

ARG BUILD_IMAGE=golang:${GO_VERSION}-${OS_NAME}${OS_VERSION}
ARG RUN_IMAGE=${OS_NAME}:${OS_VERSION}

FROM ${BUILD_IMAGE} AS build

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN go build -o /app/bin/banner-service /app/cmd/banner-service/main.go

FROM ${RUN_IMAGE}

COPY --from=build /app/bin /bin

EXPOSE ${PORT}

CMD ["/bin/banner-service"]