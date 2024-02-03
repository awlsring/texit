# BUILD
FROM golang:1.21.5-alpine as build
WORKDIR /src
COPY . .
RUN apk add --no-cache gcc musl-dev
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags '-extldflags "-static"' -v -o api cmd/api/main.go
RUN chmod +x api

# IMAGE
FROM scratch
COPY --from=build /src/api /usr/local/bin/api
EXPOSE 7032
CMD ["api"]
