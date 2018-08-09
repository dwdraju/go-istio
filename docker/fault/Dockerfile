FROM golang:1.10 AS build-env
ADD . /app
WORKDIR /app
RUN go get -d ./... && \
    CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:3.6
WORKDIR /app/
COPY --from=build-env /app/main .
ENTRYPOINT ["./main"]
