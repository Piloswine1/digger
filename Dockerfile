LABEL org.opencontainers.image.source=https://github.com/Piloswine1/digger
LABEL org.opencontainers.image.licenses=MIT

FROM docker.io/library/golang:1.26-alpine as build

RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o /app ./cmd/main.go

FROM gcr.io/distroless/static-debian12 as final
COPY --from=build /app /app
CMD ["/app"]
