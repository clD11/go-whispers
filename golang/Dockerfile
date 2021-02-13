# Dockerfile References: https://docs.docker.com/engine/reference/builder/

FROM golang:1.15

# Enable to use modules
# ENV GO111MODULE=on

# Add Maintainer Info
LABEL maintainer="clD11"

# Create directory
RUN mkdir -p $GOPATH/src/github.com/clD11/go-whispers

# Set the Current Working Directory inside the container
WORKDIR $GOPATH/src/github.com/clD11/go-whispers

## Download all the dependencies cache for faster builds
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . ./golang/services

#RUN ls -l $GOPATH/src/github.com/clD11/go-whispers
#
### Install the package
#RUN go install -v ./..
#
## This container exposes port 8080 to the outside world
#EXPOSE 8080
#
## Run the executable
#ENTRYPOINT ["cryptopricepublisher"]