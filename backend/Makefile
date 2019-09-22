.PHONY: mongo

MONGO_CONTAINER_NAME := budget-mongo

mongo-docker:
	@docker start $(MONGO_CONTAINER_NAME) >/dev/null \
	|| docker run --name $(MONGO_CONTAINER_NAME) -d -P mongo:4.2 >/dev/null
	$(eval port = $(shell docker inspect budget-mongo -f '{{with (index (index .NetworkSettings.Ports "27017/tcp") 0)}}{{.HostPort}}{{end}}'))
	@echo "export MONGODB_URI=mongodb://localhost:$(port)/budget"

api-dev:
	$(shell make mongo-docker); go run ./cmd/api/main.go

generate:
	go generate ./...