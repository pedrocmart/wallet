.PHONY: local
local:
	@docker-compose up --build

.PHONY: stop-local
stop-local:
	@docker-compose down

.PHONY: test
test:
	@go test -race -cover ./internal/wallet/...