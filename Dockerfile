FROM golang:1.24.1

WORKDIR /app
RUN go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /main
EXPOSE 8080
CMD ["/main"]
