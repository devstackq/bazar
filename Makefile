build:
	go mod download && CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/bazar/main.go

run: build app
	docker-compose up --build server

run: build postgres
	docker-compose up --build postgresdb