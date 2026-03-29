.PHONY: infra-up infra-down topics build lint test

infra-up:
	docker compose -f infra/docker-compose.yml up -d

infra-down:
	docker compose -f infra/docker-compose.yml down

topics:
	docker exec ridecore-redpanda rpk topic create driver-location-events order-requests assignments dead-letter-orders --partitions 6 --replicas 1

build:
	go build ./cmd/...

lint:
	golangci-lint run ./...

test:
	go test ./... -race -timeout 30s
