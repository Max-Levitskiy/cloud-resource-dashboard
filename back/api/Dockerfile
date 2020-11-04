FROM golang:alpine AS builder
RUN mkdir /build
WORKDIR /build
COPY . .
RUN go build -o api api.go

FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/api .
COPY --from=builder /build/conf/*.yaml ./conf/
ENV CONFIG_PATH=/root/conf
CMD ["./api"]
