REPOSITORY=github.com/rjjatson
SERVICE=bot-kuraifu-chan

RUN_BUILDER= docker run --rm -v $(CURDIR):/usr/src/$(REPOSITORY)/$(SERVICE) -w /usr/src/$(REPOSITORY)/$(SERVICE) \
			-e GOPATH=/usr -e CGO_ENABLED=0 -e GOOS=linux \
            rjjatson/golang-builder:1.12-alpine

build:
	$(RUN_BUILDER) dep ensure -v
	$(RUN_BUILDER) go build -o bin ./cmd/$(SERVICE)
	docker build --tag="$(SERVICE):latest" .

run:
	docker-compose up