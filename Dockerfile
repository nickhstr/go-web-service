# The same base alpine image is used for both stages. Ideally, a scratch image would
# be preferable when running the app, however, this makes use of any packages which
# rely on C libraries much more difficult to manage.
# While the final image size won't be as small as one based on scratch, compatibility
# with all packages remains intact.

# Stage 1 - Build app
FROM golang:1.11.2-alpine3.8@sha256:e7462ca504afc789d289f2bb5fd471815cc11833439d2fe4e61915b190045359

# Install git
RUN apk update
RUN apk add git gcc g++

WORKDIR /app
COPY ./app ./app
COPY ./main.go \
    ./go.mod \
    ./go.sum \
    ./
RUN go build -a -o ./bin/go-web-service ./main.go

# Stage 2 - Run app
FROM golang:1.11.2-alpine3.8@sha256:e7462ca504afc789d289f2bb5fd471815cc11833439d2fe4e61915b190045359

COPY --from=0 /app/bin/go-web-service go-web-service
ENV GO_ENV=production \
    PORT=3000

EXPOSE 3000
CMD ["./go-web-service"]
