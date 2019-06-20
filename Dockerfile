##################################
# STAGE 1 build executable binary #
##################################
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git bash ca-certificates openssh

WORKDIR $GOPATH/src/github.com/timurkash/task_example
COPY . .

RUN go get -d -v
# Build the binary
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o /go/bin/task_example
###############################
# STAGE 2 build a small image #
###############################
FROM scratch

# Copy our static executable.
COPY --from=builder /go/bin/task_example /go/bin/task_example

# Port on which the service will be exposed.
EXPOSE 3000

# Run binary
CMD ["/go/bin/task_example"]