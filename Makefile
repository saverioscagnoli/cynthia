build-bot:
	go build -o bin/cynthia ./bot

run-bot:
	go build -o bin/cynthia ./bot
	./bin/cynthia $(ARGS)

build-pkapi:
	go build -0 bin/pkapi ./pokemon

run-pkapi:
	go build -o bin/pkapi ./pokemon
	./bin/pkapi $(ARGS)

clean:
	rm -rf bin/
