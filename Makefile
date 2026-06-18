build-app:
	go build -o bin/cynthia ./cmd/app

run-app:
	go build -o bin/cynthia ./cmd/app
	./bin/cynthia $(ARGS)

build-pkapi:
	go build -o bin/pkapi ./cmd/pkapi/api.go

run-pkapi:
	go build -o bin/pkapi ./cmd/pkapi/api.go
	./bin/pkapi $(ARGS)

run-seed:
	go run ./cmd/pkapi/seed.go

app-healthcheck:
	go run ./cmd/healthcheck

db-up:
	docker compose -f db-compose.yaml up -d

db-down:
	docker compose -f db-compose.yaml down

db-reset:
	docker compose -f db-compose.yaml down -v
	docker compose -f db-compose.yaml up -d

clean:
	rm -rf bin/
