.PHONY: gqlgen
gqlgen: # gqlgen コード生成
	docker compose exec app go run github.com/99designs/gqlgen

.PHONY: tidy
tidy:
	@docker compose exec app go mod tidy

.PHONY: run
run:
	@docker compose up -d --build

.PHONY: stop
stop:
	@docker compose down


.PHONY: console
console: ## hasura console
	@hasura --project hasura console --admin-secret secret --skip-update-check --endpoint "http://localhost:${HASURA_PORT}"