# Build Svelte files

FROM node:10-alpine as client_builder

COPY /client /node/app/client/

WORKDIR /node/app/client

RUN yarn && yarn build

# Build Go Server

FROM golang:latest as base_builder

WORKDIR /go/src/app

ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

# Distroless for smaller sized containers

FROM gcr.io/distroless/base

COPY --from=base_builder /go/src/app/public /public/

COPY --from=base_builder /go/src/app/.env /

COPY --from=base_builder /go/bin/app /

COPY --from=client_builder /node/app/public/build/ /public/build/

CMD ["/app"]
