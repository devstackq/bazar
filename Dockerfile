#pull golang base  image
FROM golang:latest
LABEL name devstack

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

RUN go build -o main /cmd/bazar/main.go

CMD [ "./main" ]




# FROM alpine:latest

# RUN apk --no-cache add ca-certificates
# WORKDIR /root/

# COPY ./.bin/app .
# COPY ./config/ ./config/