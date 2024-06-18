FROM golang:1.21.10 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .

RUN CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -C cmd -ldflags="-w -s" -o tempcep .

FROM scratch
COPY --from=builder /app/* ./
CMD ["./tempcep"]