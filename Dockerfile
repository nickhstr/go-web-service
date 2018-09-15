# Stage 1 - Build app
FROM golang:1.11-alpine

# Add git and gcc/libc packages
RUN apk add --no-cache git build-base

WORKDIR /app
COPY . /app
RUN go build -o /app/bin/go-web-service /app/main.go

RUN apk del git build-base

# Stage 2 - Run app
FROM alpine

COPY --from=0 /app/bin/go-web-service go-web-service
ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./go-web-service"]
