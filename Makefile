dockerinit:
	 docker run --name postgresdisko -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

postgres:
	docker exec -it postgresdisko psql

createdb:
	docker exec -it postgresdisko createdb --username=root --owner=root diskoDB 

dropdb:
	docker exec -it postgresdisko dropdb go-chat

migratecreate:
	migrate create -ext sql -dir db/migrations add_users_table

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/diskoDB?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/diskoDB?sslmode=disable" -verbose down 

.PHONY: dockerinit postgres createdb dropdb migrateup migratedown

