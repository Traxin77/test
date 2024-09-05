FROM golang:1.23.0

WORKDIR /app

# Copy the Go module files and install dependencies
COPY go.mod .
COPY server.go .

RUN go get
RUN go build -o bin .
ENTRYPOINT [ "/app/bin" ]