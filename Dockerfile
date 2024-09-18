FROM golang:1.23.1 AS build
WORKDIR /build
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build .

FROM debian:12-slim
COPY --from=build /build/hello-pss /usr/local/bin/hello-pss
CMD ["hello-pss"]
EXPOSE 8000
