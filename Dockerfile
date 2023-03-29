FROM golang:latest AS build-env

WORKDIR /src

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=build-env /app .

RUN chmod +x ./app

CMD ["./app"]