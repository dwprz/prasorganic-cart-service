DB_URL := 'postgres://postgres:rahasia@localhost:5436/prasorganic_cart_service?sslmode=disable'

# Migration
# example: make migration name=create_cart_table
.PHONY: migration
migration:
	migrate create -ext sql -dir migration -seq ${name}

.PHONY: migrate-up
migrate-up:
	migrate -database ${DB_URL} -path migration up

.PHONY: migrate-down
migrate-down:
	migrate -database ${DB_URL} -path migration down

.PHONY: licenses
licenses:
	rm -rf ./LICENSES
	go-licenses save ./... --save_path=./LICENSES

.PHONY: start
start:
	rm -f ./cmd/main
	go build -o cmd/main cmd/main.go
	./cmd/main