migrate-up :
	migrate -path infra/migration -database "postgres://root:root@localhost:5432/go-chat?sslmode=disable" up

migrate-down :
	migrate -path infra/migration -database "postgres://root:root@localhost:5432/go-chat?sslmode=disable" down