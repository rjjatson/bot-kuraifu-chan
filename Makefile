SERVICE="bot-kuraifu-chan"

build:
	go build -o bin ./cmd/bot-kuraifu-chan
	docker build --tag="$(SERVICE):latest" .

run:
	docker run $(SERVICE):latest

#push:
#	docker tag $(SERVICE):latest $(ECR_ADDRESS):latest
#	docker push $(ECR_ADDRESS):latest