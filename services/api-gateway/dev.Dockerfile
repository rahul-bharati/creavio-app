FROM golang:1.25-alpine

LABEL author="Rahul Bharati"
LABEL description="Api Gateway for Creavio"

WORKDIR /src

COPY go.mod  ./

RUN --mount=type=cache,target=/go/pkg/mod go mod download

COPY services/api-gateway ./services/api-gateway

RUN --moount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /src/services/api-gateway ./services/api-gateway/cmd/main.go


RUN printf '#!bin/sh\nexec /src/services/api-gateway/main\n' > /user/local/bin/run.sh && chmod +x /usr/local/bin/run.sh

EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/run.sh"]