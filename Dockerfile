# Stage 1 - Build app
FROM golang:1.11@sha256:e8e4c4406217b415c506815d38e3f8ac6e05d0121b19f686c5af7eaadf96f081

WORKDIR /app
COPY ./app ./app
COPY ./main.go \
    ./go.mod \
    ./go.sum \
    ./
RUN CGO_ENABLED=0 go build -o ./bin/go-web-service ./main.go

# Stage 2 - Run app
FROM alpine@sha256:621c2f39f8133acb8e64023a94dbdf0d5ca81896102b9e57c0dc184cadaf5528

COPY --from=0 /app/bin/go-web-service go-web-service
ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./go-web-service"]
