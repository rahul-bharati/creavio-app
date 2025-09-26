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
    go build -o /src/services/user ./services/user/cmd

RUN printf '#!/bin/sh\nexec /app/user\n' > /usr/local/bin/run.sh && chmod +x /usr/local/bin/run.sh

EXPOSE 8000

ENTRYPOINT ["/usr/local/bin/run.sh"]