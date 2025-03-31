FROM golang:1.23 AS build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ratelimit ./cmd

FROM scratch
WORKDIR /app
COPY --from=build /app/ratelimit .

EXPOSE 8080
ENTRYPOINT [ "./ratelimit" ]