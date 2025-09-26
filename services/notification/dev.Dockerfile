FROM golang:1.25-alpine

LABEL author="Rahul Bharati"
LABEL description="Notification Microservice app for Creavio"

WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod  ./

RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY services/notification ./services/notification

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /src/ ./services/notification/cmd/main.go

EXPOSE 8002

ENTRYPOINT ["/src/main"]