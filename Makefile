default: help

.PHONY: help
help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

.PHONY: dependencies
dependencies: ## install dependencies
	go mod tidy
	go mod verify
	go mod vendor

.PHONY: tools
tools: ## get tools
	git config core.hooksPath .githooks
	go get -u github.com/golang/mock/gomock
	go get -u github.com/golang/mock/mockgen

.PHONY: generate-mocks
generate-mocks: ## generate mocks
	mockgen -package=mock_redis -source internal/pkg/redis/redis.go -destination=internal/pkg/redis/mocks/redis_mock.go
	mockgen -package=mock_validator -source internal/pkg/validator/validator.go -destination=internal/pkg/validator/mocks/validator_mock.go

.PHONY: generate-api-schema
generate-api-schema: ## generate api schema using swagger
	./scripts/bundle-api.sh
	./scripts/generate-code.sh

# ifeq (run-local,$(firstword $(MAKECMDGOALS)))
#   # use the rest as arguments for "run-local"
#   RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
#   # ...and turn them into do-nothing targets
#   $(eval $(RUN_ARGS):dummy;@:)
# endif

# dummy: ## don't touch this, removing this mucks up the run-local script
# 	@:

.PHONY: run-local
run-local: ## run the application locally
	go run cmd/main.go

# .PHONY: migrate
# migrate: ## wraps golang-migrate. Use with arguments such as 'up', 'down 2', 'version' etc. run 'migrate help for more info'
# 	migrate -database '$(DB_URI)' -path ./migrations $(RUN_ARGS)

.PHONY: migrate-create
migrate-create: ## creates migrations with one argument for a suffix
	migrate -database '$(db_uri)' create -dir migrations/$(service) -ext .sql $(migration_name)

