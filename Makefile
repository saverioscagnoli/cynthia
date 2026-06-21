build-app:
	go build -o bin/cynthia ./cmd/app

run-app:
	go build -o bin/cynthia ./cmd/app
	./bin/cynthia $(ARGS)

run-seed:
	go run ./store/seed.go

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
