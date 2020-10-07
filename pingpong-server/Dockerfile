FROM golang:latest
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
# RUN apk add build-base
COPY . .
# RUN go build -o ./cmd/main

# RUN go get -d -v ./go.mod
CMD ["go", "run", "main.go"]
# CMD ["./cmd/main"]