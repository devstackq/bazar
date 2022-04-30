FROM golang:latest
LABEL name devstack

WORKDIR /root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

#heroku db creds
ENV  POSTGRES_USER=oktarcpjvgaorf
ENV  POSTGRES_PASSWORD=93f71f9ee4342d1fb1190493c57493ad405416c153f8bc281d7db4b3306555ad
ENV  POSTGRES_URI=ec2-52-30-67-143.eu-west-1.compute.amazonaws.com
ENV  POSTGRES_PORT=5432
ENV  POSTGRES_DB=d9pr4qe2pmgqs
ENV  POSTGRES_DATABASE=d9pr4qe2pmgqs


RUN chmod 777 postgres-data

RUN go build  ./cmd/bazar/main.go

CMD ["./main"]