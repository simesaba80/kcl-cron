FROM golang:1.23.2-alpine3.20
WORKDIR /app
RUN go install github.com/air-verse/air@latest
CMD ["air"]