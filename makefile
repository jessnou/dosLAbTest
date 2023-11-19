postgres:
	docker run --name dosLab  -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=root --owner=root dosLab

migrateup:
	migrate -path ./migrations -database "postgresql://root:secret@localhost:5432/dosLab?sslmode=disable" -verbose up