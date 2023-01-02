FROM golang:1.17-alpine as builder
RUN apk --no-cache add ca-certificates git
WORKDIR /build

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN CGO_ENABLED=0 go build

# Create final image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder ./.bin/app .
COPY --from=builder ./config/ .
EXPOSE 8080
CMD ["./myapp"]