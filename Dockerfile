FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./edge-sim .


FROM alpine:3

COPY --from=builder /app/edge-sim ./
RUN chmod +x ./edge-sim

ENV BASE_SEED=123456
ENV DEVICE_ID=-1

ENTRYPOINT ["./edge-sim"]