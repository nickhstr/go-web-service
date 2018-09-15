FROM golang:1.11-alpine

COPY ./go-web-service ./go-web-service

ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./go-web-service"]
