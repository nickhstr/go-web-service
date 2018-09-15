FROM golang:1.11-alpine

COPY ./go-web-service ./go-web-service

EXPOSE 3000
CMD ["./go-web-service"]
