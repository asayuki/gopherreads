run:
	@templ generate && go run cmd/gopherreads/main.go

migration:
	@migrate create -ext sql -dir database/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrations/migrate.go up

migrate-down:
	@go run cmd/migrations/migrate.go down