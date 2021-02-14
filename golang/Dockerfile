FROM golang:1.15 AS builder

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
RUN go install ./services/cryptopricepublisher/...

# Build smaller image
FROM scratch

#RUN apk --no-cache add ca-certificates

COPY --from=builder /go/bin /

## This container exposes port 8080 to the outside world
EXPOSE 8080

## Run the executable
ENTRYPOINT ["/cryptopricepublisher"]