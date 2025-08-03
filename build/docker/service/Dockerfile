FROM golang:1.24.2-alpine3.20 AS build
LABEL authors="IAV"
WORKDIR /PaymentService
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /PaymentService/cmd/app
RUN go build -o build

FROM alpine:3.18 AS prod
COPY --from=build /PaymentService /.
EXPOSE 8080
ENTRYPOINT ["/cmd/app/build"]