.PHONY: mongo

MONGO_CONTAINER_NAME := budget-mongo
LINT_IGNORE=should have comment or be unexported\|comment on exported

check: lint test

lint:
	@echo golint ./...
	$(eval out = $(shell golint ./... | grep -v '$(LINT_IGNORE)' | sed s/$$/xxx/))
	@echo "$(out)" | sed $$'s/xxx[ ]*/\\\n/g' | grep -v '^$$' || true
	@exit $(shell echo "$(out)" | sed $$'s/xxx[ ]*/\\\n/g' | grep -v '^$$' | wc -l)

format: generate
	go fmt ./...
	goimports -l -w .

test:
	go test ./...

mongo-docker:
	@docker start $(MONGO_CONTAINER_NAME) >/dev/null \
	|| docker run --name $(MONGO_CONTAINER_NAME) -d -P mongo:3.6 >/dev/null
	$(eval port = $(shell docker inspect budget-mongo -f '{{with (index (index .NetworkSettings.Ports "27017/tcp") 0)}}{{.HostPort}}{{end}}'))
	@echo "export MONGODB_URI=mongodb://localhost:$(port)/budget"

api-dev:
	$(shell make mongo-docker); go run ./cmd/api/main.go

generate:
	go generate ./...