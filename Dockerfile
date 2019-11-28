FROM golang:1.13-buster as build

WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

FROM gcr.io/distroless/base

COPY --from=build /go/bin/app /

CMD ["/app"]
