FROM golang:latest as builder
WORKDIR /shelly_ht_exporter
COPY . .
RUN go mod download
RUN go build

FROM debian:buster-slim
WORKDIR /
COPY --from=builder /shelly_ht_exporter/shelly_ht_exporter /bin/shelly_ht_exporter
EXPOSE 9439
CMD ["shelly_ht_exporter"]
