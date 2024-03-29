FROM golang:1.16 AS builder

ARG SERVICE_NAME

# Create directory
RUN mkdir -p $GOPATH/src/github.com/clD11/go-whispers/golang

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/clD11/go-whispers/golang

# Copy everything from the current directory to the workdir inside the container
COPY . .

RUN go mod download

ENV GOOS linux
ENV CGO_ENABLED 0
ENV GOARCH amd64

## Install binary into go/bin
RUN go install ./services/$SERVICE_NAME/...

# Build smaller image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin /
COPY --from=builder /go/bin /

ARG SERVICE_NAME
RUN mv $SERVICE_NAME go-whispers-app

ARG EXPOSE_PORT
EXPOSE $EXPOSE_PORT

## Run the executable
ENTRYPOINT ["/go-whispers-app"]