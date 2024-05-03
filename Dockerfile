FROM golang:1.22.2-alpine3.19 as base
WORKDIR /app

COPY go.* .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd

FROM alpine:3.19
WORKDIR /app
COPY --from=base /app/main /app

EXPOSE 3333

CMD [ "./main" ]