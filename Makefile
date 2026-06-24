run-webui:
	cd www && yarn install && yarn dev

build-webui:
	cd www && yarn install && yarn build

build-app:
	go build -mod=vendor -o bin/cynthia ./cmd/app

run-app:
	go build -mod=vendor -o bin/cynthia ./cmd/app
	./bin/cynthia $(ARGS)

make run-app-prod:
	$(MAKE) build-webui
	go build -mod=vendor -o bin/cynthia ./cmd/app
	./bin/cynthia $(ARGS)

run-seed:
	go run -mod=vendor ./store/seed.go

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
