build-bot:
	go build -o bin/cynthia ./bot

build-other:
	go build -o bin/other ./other

run-bot: build-bot
	./bin/cynthia

run-other: build-other
	./bin/other

clean:
	rm -rf bin/
