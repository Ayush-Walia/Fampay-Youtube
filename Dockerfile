FROM golang:alpine3.17 as builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o app .

# deployment image
FROM alpine:3.17

WORKDIR /bin/

# copy app from builder
COPY --from=builder /build/app .

CMD [ "./app" ]