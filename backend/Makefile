.PHONY: mongo

MONGO_CONTAINER_NAME := budget-mongo

check: lint

lint:
	$(eval out = $(shell golint ./... | grep -v 'should have comment or be unexported' | sed s/$$/xxx/))
	@echo "$(out)" | grep -v '^$$' | sed $$'s/xxx /\\\n/g'
	@exit $(shell echo "$(out)" | sed $$'s/xxx /\\\n/g' | grep -v '^$$' | wc -l)

mongo-docker:
	@docker start $(MONGO_CONTAINER_NAME) >/dev/null \
	|| docker run --name $(MONGO_CONTAINER_NAME) -d -P mongo:3.6 >/dev/null
	$(eval port = $(shell docker inspect budget-mongo -f '{{with (index (index .NetworkSettings.Ports "27017/tcp") 0)}}{{.HostPort}}{{end}}'))
	@echo "export MONGODB_URI=mongodb://localhost:$(port)/budget"

api-dev:
	$(shell make mongo-docker); go run ./cmd/api/main.go

generate:
	go generate ./...