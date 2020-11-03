FROM golang:1.15.3-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o go-api

WORKDIR /dist
RUN cp /build/go-api .

# Multistage build
FROM scratch
COPY --from=builder /dist/go-api /
# COPY .config .

EXPOSE 3010
CMD ["/go-api"]