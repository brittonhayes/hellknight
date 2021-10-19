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
ENV BOT_TOKEN $BOT_TOKEN

EXPOSE 8080
CMD ["/usr/local/bin/hellknight"]