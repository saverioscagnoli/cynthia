build-bot:
	go build -o bin/cynthia ./cmd/bot

run-bot:
	go build -o bin/cynthia ./cmd/bot
	./bin/cynthia $(ARGS)

build-pkapi:
	go build -0 bin/pkapi ./cmd/pkapi/api.go

run-pkapi:
	go build -o bin/pkapi ./cmd/pkapi/api.go
	./bin/pkapi $(ARGS)

clean:
	rm -rf bin/
