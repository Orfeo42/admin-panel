# Build.
FROM golang:1.23.4-alpine AS build-stage

RUN go install github.com/a-h/templ/cmd/templ@latest

WORKDIR /app

COPY . .

RUN templ generate
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api/main.go

# Deploy.
FROM alpine:latest  AS release-stage

WORKDIR /

COPY --from=build-stage /app/main .
COPY --from=build-stage /app/cmd/web/assets /assets

EXPOSE 8080

CMD ["./main"]
