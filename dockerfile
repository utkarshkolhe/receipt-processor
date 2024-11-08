FROM golang:1.22

WORKDIR /app
COPY go.mod .
COPY main.go .
COPY main_test.go .
COPY api ./api
COPY test ./test
RUN go get
RUN go mod tidy
RUN go test ./...
RUN go build -o bin .

ENTRYPOINT ["/app/bin"]
