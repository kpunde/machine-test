FROM golang:1.12-alpine

RUN apk add --no-cache git
RUN ["mkdir", "-p", "/resources/logs"]

# Set the Current Working Directory inside the container
WORKDIR /app/project

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./out/project .

# Run the binary program produced by `go install`
CMD ["./out/project"]