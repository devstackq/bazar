FROM golang:latest as build
LABEL name devstack

WORKDIR /root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

#heroku db creds
ENV  POSTGRES_USER=rwecsnywflfryx
ENV  POSTGRES_PASSWORD=04d64cea97a6d17d7b4de6b5e9b0755e2ab090bbdc24b1be7e11943d618c777c
ENV  POSTGRES_URI=ec2-44-199-143-43.compute-1.amazonaws.com
ENV  POSTGRES_PORT=5432
ENV  POSTGRES_DB=dffqhfqt0ef62l
ENV  POSTGRES_DATABASE=dffqhfqt0ef62l

ENV PORT=6969
# ENV APP_PORT=6969

RUN go build  ./cmd/bazar/main.go

FROM heroku/heroku:20

WORKDIR /app
COPY --from=build /root/main /app/

CMD ["./main"]