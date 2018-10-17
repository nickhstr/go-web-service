# Stage 1 - Build app
FROM golang:1.11.1@sha256:63ec0e29aeba39c0fe2fc6551c9ca7fa16ddf95394d77ccee75bc7062526a96c

WORKDIR /app
COPY ./app ./app
COPY ./main.go \
    ./go.mod \
    ./go.sum \
    ./
RUN CGO_ENABLED=0 go build -o ./bin/go-web-service ./main.go

# Stage 2 - Run app
FROM scratch

COPY --from=0 /app/bin/go-web-service go-web-service
ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./go-web-service"]
