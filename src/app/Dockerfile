FROM golang:1.15.2

RUN mkdir /app
RUN mkdir -p /upload
WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main hello.go
EXPOSE 8080
CMD ["/app/main"]