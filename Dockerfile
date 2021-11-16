FROM golang:1.15-alpine3.12 AS payment-builder

WORKDIR /app
ADD . /app

RUN go build -o payment-challenge ./cmd/payment-challenge

FROM alpine:3.12 as payment-app

COPY --from=cas-builder /app/payment-challenge /app/

CMD [ "/app/payment-challenge" ]
