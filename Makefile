build:
	docker-compose build market-app

run: 
	docker-compose up market-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:662011Egor@localhost:30432/postgres?sslmode=disable' up

remove-migrate:
	migrate -path ./schema -database 'postgres://postgres:662011Egor@localhost:30432/postgres?sslmode=disable' down
test:
	go test -v ./...
