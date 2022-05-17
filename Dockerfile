FROM golang:latest as build
LABEL name devstack

WORKDIR /root

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
EXPOSE 6969

#heroku db creds

ENV  POSTGRES_USER=pkbdytztfofcwf
ENV  POSTGRES_PASSWORD=ee6b8ea0ca73b0c591c993a67127bdb4fb29af099ffc354d32ce68507e809118
ENV  POSTGRES_URI=ec2-54-228-218-84.eu-west-1.compute.amazonaws.com
ENV  POSTGRES_PORT=5432
ENV  POSTGRES_DATABASE=d9pr4qe2pmgqs

ENV PORT=6969
# ENV APP_PORT=6969

RUN go build  ./cmd/bazar/main.go

FROM heroku/heroku:20

WORKDIR /app
COPY --from=build /root/main /app/

CMD ["./main"]