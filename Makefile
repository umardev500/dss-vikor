up:
	migrate -database "postgres://root:root@127.0.0.1:5432/db_spk?sslmode=disable" -path database/migrations/schemas up
down:
	migrate -database "postgres://root:root@127.0.0.1:5432/db_spk?sslmode=disable" -path database/migrations/schemas down