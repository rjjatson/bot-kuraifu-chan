SERVICE="bot-kuraifu-chan"

build:
	docker run --rm -v $(CURDIR):/usr/src/$(SERVICE) -w /usr/src/$(SERVICE) golang:1.12-alpine \
	go build -o bin ./cmd/bot-kuraifu-chan
	docker build --tag="$(SERVICE):latest" .

run:
	docker run $(SERVICE):latest

#push:
#	docker tag $(SERVICE):latest $(ECR_ADDRESS):latest
#	docker push $(ECR_ADDRESS):latest