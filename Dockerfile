FROM golang:latest
LABEL name devstack

WORKDIR /root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

ENV  POSTGRES_USER=postgres
ENV  POSTGRES_PASSWORD=postgres
ENV  POSTGRES_DATABASE=testdb
ENV  POSTGRES_URI=postgresdb
ENV  POSTGRES_PORT=5432
ENV  POSTGRES_DB=testdb

RUN chmod 777 postgres-data


# CMD [ "./main" ]
FROM postgres:latest
RUN   ./build/sql/create_tables.sql:/docker-entrypoint-initdb.d/create_tables.sql

FROM heroku/heroku:18
WORKDIR /app
# COPY --from=0 /root/main /app

RUN go build  ./cmd/bazar/main.go
CMD ["./main"]