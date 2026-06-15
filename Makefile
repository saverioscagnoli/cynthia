build-app:
	go build -o bin/cynthia ./cmd/app

run-app:
	go build -o bin/cynthia ./cmd/app
	./bin/cynthia $(ARGS)

build-pkapi:
	go build -0 bin/pkapi ./cmd/pkapi/api.go

run-pkapi:
	go build -o bin/pkapi ./cmd/pkapi/api.go
	./bin/pkapi $(ARGS)

run-seed:
	go run ./cmd/pkapi/seed.go

app-healthcheck:
	go run ./cmd/healthcheck

clean:
	rm -rf bin/
