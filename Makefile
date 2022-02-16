run:
	go run main.go run

migrate:
	go run main.go migrate

migrate-rollback:
	go run main.go migrate:rollback

migrate-refresh:
	go run main.go migrate:refresh

migrate-reset:
	go run main.go migrate:reset

migrate-status:
	go run main.go migrate:status