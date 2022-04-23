FROM golang:latest
LABEL name devstack

WORKDIR /root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

ENV  POSTGRES_USER=rwecsnywflfryx
ENV  POSTGRES_PASSWORD=92b83e192256d2b9c4d4173dfb66bedf264c78bd4a602a8972d987ac501f2cc8
ENV  POSTGRES_URI=ec2-99-80-170-190.eu-west-1.compute.amazonaws.com
ENV  POSTGRES_PORT=5432
ENV  POSTGRES_DB=d448svjcam10be

RUN go build  ./cmd/bazar/main.go

CMD ["./main"]