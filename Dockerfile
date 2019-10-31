FROM golang:1.13
COPY . /gas
WORKDIR /gas
RUN CC=$(which musl-gcc) go build --ldflags '-w -linkmode external -extldflags "-static"' -o gas main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /gas
COPY --from=0 /gas /gas
CMD ["./gas"]
