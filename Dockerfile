# Multi-stage
# FROM golang:1.19.2 as builder
# WORKDIR /app/hello
# COPY ./hello ./
# RUN CGO_ENABLED=0 GOOS=linux go build -o /hello

# FROM gcr.io/distroless/base-debian11 as final
# WORKDIR /
# COPY --from=builder /hello /hello
# USER nonroot:nonroot
# ENTRYPOINT ["/hello"]

# Hot Reload
FROM golang:latest
WORKDIR /app/hello
COPY ./hello ./
RUN go get github.com/githubnemo/CompileDaemon
RUN go install -mod=mod github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main