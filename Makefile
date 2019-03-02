SERVICE="bot-kuraifu-chan"
#SERVICE=health

build:
	docker run --rm -v $(CURDIR):/usr/src/$(SERVICE) -w /usr/src/$(SERVICE) golang:1.12-alpine \
	go build -o bin ./cmd/$(SERVICE)
	docker build --tag="$(SERVICE):latest" .

run:
	docker run -p 8080:8080 $(SERVICE):latest

#push:
#	docker tag $(SERVICE):latest $(ECR_ADDRESS):latest
#	docker push $(ECR_ADDRESS):latest