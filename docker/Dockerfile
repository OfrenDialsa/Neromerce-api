FROM golang:1.25.5

WORKDIR /app

RUN mkdir -p /app/.cache/go-build && \
    mkdir -p /app/.cache/go-mod && \
    chown -R 1000:1000 /app

RUN go install github.com/air-verse/air@latest

USER 1000:1000

COPY --chown=1000:1000 . .

ENV GOMODCACHE=/app/.cache/go-mod
ENV GOCACHE=/app/.cache/go-build

RUN go mod download

CMD ["air"]
