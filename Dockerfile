# build stage
FROM golang:alpine AS builder
RUN apk add --no-cache build-base
WORKDIR /src
COPY . .
RUN go build

# server image
FROM golang:alpine
COPY --from=builder /src/hellknight /usr/local/bin/

ARG BOT_TOKEN

CMD ["/usr/local/bin/hellknight"]