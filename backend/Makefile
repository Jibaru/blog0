run:
	@go run cmd/app/main.go

build:
	@swag init -g cmd/app/main.go
	@go build -o bin/app cmd/app/main.go

gen:
	@rm -R internal/infra/persistence/postgres
	@rm -R internal/domain/dao
	@gormless -i internal/domain -o internal/infra/persistence --driver postgres
	@gormless -i internal/domain -o internal/domain/dao --interface