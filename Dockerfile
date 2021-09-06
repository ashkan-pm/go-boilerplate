FROM golang:1.16
WORKDIR /app

COPY ./Makefile ./Makefile
COPY ./go.mod ./go.mod
COPY ./src ./src

RUN make install
RUN make build

FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache gcompat
COPY --from=0 /app/build/go-boilerplate .
ENTRYPOINT ["./go-boilerplate"]
