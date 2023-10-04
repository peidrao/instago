db.up:
	docker-compose up -d 

db.down:
	docker stop instago_db_1

db.exclude:
	docker stop instago_db_1
	docker rm instago_db_1

test:
	go test -v ./tests/...

run:
	go run cmd/app/main.go

lint: 
	/home/pedro/go/bin/golangci-lint run

fmt:
	/home/pedro/go/bin/golangci-lint run --fix
