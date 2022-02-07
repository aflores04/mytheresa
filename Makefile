dependency:
	@go get -v ./...

docker-app-up:
	@docker-compose build
	@docker-compose up

docker-app-down:
	@docker-compose down

docker-test-up:
	@docker-compose -f ./.it/docker-compose.yml up -d

docker-test-down:
	@docker-compose -f ./.it/docker-compose.yml down

app-up: docker-test-down docker-app-down docker-app-up

app-down: docker-test-down docker-app-down

test: docker-test-down docker-test-up dependency
	@go test -v ./...