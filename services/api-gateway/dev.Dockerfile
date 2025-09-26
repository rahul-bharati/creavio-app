FROM golang:1.25-alpine

LABEL author="Rahul Bharati"
LABEL description="Api Gateway for Creavio"

WORKDIR /src
ENV CGO_ENABLED=0

COPY go.mod  ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY services/api-gateway ./services/api-gateway

RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /src/ ./services/api-gateway/cmd/main.go


EXPOSE 8000

ENTRYPOINT ["/src/main"]