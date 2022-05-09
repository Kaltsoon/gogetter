FROM golang:1.17-alpine AS build

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o gogetter

FROM alpine:3.15
RUN apk add ca-certificates

WORKDIR /usr/src/app

COPY --from=build /usr/src/app/gogetter .

ENTRYPOINT ["./gogetter"]
