FROM golang:1.25-alpine

LABEL author="Rahul Bharati"
LABEL description="User Microservice app for Creavio"

WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod  ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY services/user ./services/user

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /src/ ./services/user/cmd/main.go

EXPOSE 8001

ENTRYPOINT ["/src/main"]