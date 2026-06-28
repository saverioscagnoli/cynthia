run-webui:
	cd www && yarn install && yarn dev

build-webui:
	cd www && yarn install && yarn build

build-app:
	go mod vendor
	go build -mod=vendor -o bin/camilla ./cmd/app

run-app:
	go mod vendor
	go build -mod=vendor -o bin/camilla ./cmd/app
	./bin/camilla $(ARGS)

make run-app-prod:
	$(MAKE) build-webui
	$(MAKE) run-app

run-seed:
	python python/seed/main.py

app-healthcheck:
	go run -mod=vendor ./cmd/healthcheck

db-up:
	docker compose -f db-compose.yaml up -d

db-down:
	docker compose -f db-compose.yaml down

db-reset:
	docker compose -f db-compose.yaml down -v
	docker compose -f db-compose.yaml up -d

clean:
	rm -rf bin/
