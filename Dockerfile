FROM golang:1.22 AS build
WORKDIR /go/src/github.com/thmshmm/xk6-opentelemetry
COPY . .
RUN go install go.k6.io/xk6/cmd/xk6@latest
RUN xk6 build --with xk6-opentelemetry=. --with github.com/mostafa/xk6-kafka

FROM alpine:latest as base
RUN apk add --no-cache ca-certificates && \
  addgroup -S app && adduser -S -G app app

FROM scratch
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=base /etc/passwd /etc/passwd
COPY --from=build /go/src/github.com/thmshmm/xk6-opentelemetry/k6 /k6
USER app
ENTRYPOINT ["/k6"]
