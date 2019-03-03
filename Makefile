SERVICE="bot-kuraifu-chan"

build:
	docker run --rm -v $(CURDIR):/usr/src/$(SERVICE) -w /usr/src/$(SERVICE) golang:1.12-alpine \
	go build -o bin ./cmd/$(SERVICE)
	docker build --tag="$(SERVICE):latest" .

run:
	docker run --rm -d -p 8080:8080 --name $(SERVICE) $(SERVICE):latest

remove:
	docker rm $(SERVICE) -f

rebuild: build remove run