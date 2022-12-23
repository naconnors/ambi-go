db_user := "postgres"
db_password := "postgres"
db_host := "localhost"
db_name := "ambi_go_dev"

.phony: install
install: 
	go install github.com/pressly/goose/v3/cmd/goose@latest

.phony: setup
setup:
	goose -dir ./migrations/ postgres "user=$(db_user) password=$(db_password) host=$(db_host) sslmode=disable" up-to 20221212141916

.phony: migrate-up
migrate-up: install
	goose -dir ./migrations/ postgres "user=$(db_user) password=$(db_password) host=$(db_host) dbname=$(db_name) sslmode=disable" up

.phony: migrate-down
migrate-down:
	goose -dir ./migrations/ postgres "user=$(db_user) password=$(db_password) host=$(db_host) dbname=$(db_name) sslmode=disable" down-to 20221212141916

.phony: migrate-status
migrate-status:
	goose -dir ./migrations/ postgres "user=$(db_user) password=$(db_password) host=$(db_host) dbname=$(db_name) sslmode=disable" status

.phony: run
run: vendor
	go run ./...

.phony: vendor
vendor:
	go mod vendor