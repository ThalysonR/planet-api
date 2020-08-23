FROM golang:1.15.0-alpine AS builder
WORKDIR /app
COPY . .
RUN go get -u github.com/swaggo/swag/cmd/swag
RUN $GOPATH/bin/swag init
RUN go build -o planets-api

FROM alpine:3.12
WORKDIR /app
COPY --from=builder /app/planets-api .
EXPOSE 8080
CMD /app/planets-api