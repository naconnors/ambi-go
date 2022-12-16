db_user := "postgres"
db_password := "postgres"
db_host := "localhost"
db_name := "ambi_go_dev"

.phony: install
install: 
	go install github.com/pressly/goose/v3/cmd/goose@latest

.phony: migrate-up
migrate-up: install
	goose -dir ./migrations/ postgres "user=$(db_user) password=$(db_password) host=$(db_host) dbname=$(db_name) sslmode=disable" up

.phony: migrate-down
migrate-down:
	goose -dir ./migrations/ postgres "user=$(db_user) password=$(db_password) host=$(db_host) dbname=$(db_name) sslmode=disable" down

.phony: run
run: vendor
	go run ./...

.phony: vendor
vendor:
	go mod vendor