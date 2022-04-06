#pull golang base  image
FROM golang:latest
LABEL name devstack

WORKDIR /root/

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

RUN go build -o main ./cmd/bazar/main.go

CMD [ "./main" ]
