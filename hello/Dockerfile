FROM golang:1.19.2 as builder
WORKDIR /app

RUN go mod init hello

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /hello

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=builder /hello /hello

USER nonroot:nonroot
ENTRYPOINT ["/hello"]