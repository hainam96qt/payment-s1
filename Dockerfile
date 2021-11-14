FROM golang:1.15-alpine3.12 AS payment-builder

WORKDIR /app
ADD . /app

RUN go build -o controller ./cmd/controller

FROM alpine:3.12 as payment-app

COPY --from=cas-builder /app/controller /app/

CMD [ "/app/payment-challenge" ]
