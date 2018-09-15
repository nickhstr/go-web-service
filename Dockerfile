# Stage 1 - Build app
FROM golang:1.11@sha256:e8e4c4406217b415c506815d38e3f8ac6e05d0121b19f686c5af7eaadf96f081

WORKDIR /app
COPY . /app
RUN CGO_ENABLED=0 go build -o /app/bin/go-web-service /app/main.go

# Stage 2 - Run app
FROM alpine

COPY --from=0 /app/bin/go-web-service go-web-service
ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./go-web-service"]
