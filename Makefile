build:
	go build -o bin/grankings .

exec: build
	./bin/grankings

run:
	clear && go run .

clean:
	rm bin/grankings
