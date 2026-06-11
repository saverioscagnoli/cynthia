build-bot:
	go build -o bin/cynthia ./bot

run-bot:
	go build -o bin/cynthia ./bot
	./bin/cynthia $(ARGS)

clean:
	rm -rf bin/
