FROM golang:1.23.2-alpine3.20
WORKDIR /app
RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
RUN go install github.com/air-verse/air@latest
CMD ["air"]