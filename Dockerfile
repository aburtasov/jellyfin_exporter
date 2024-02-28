FROM golang:1.18 AS build
WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o jellyfin_exporter ./main.go


FROM alpine:latest
COPY --from=build /build/jellyfin_exporter /app/jellyfin_exporter
WORKDIR /app
EXPOSE 9249
ENTRYPOINT ["./jellyfin_exporter"]