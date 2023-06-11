

psql.up:
	docker-compose up -d 

psql.exclude:
	docker stop instago_db_1
	docker rm instago_db_1

run:
	go run cmd/app/main.go
