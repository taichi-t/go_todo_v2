.PHONY: gqlgen
gqlgen: # gqlgen コード生成
	docker compose exec app go run github.com/99designs/gqlgen


.PHONY: tidy
tidy:
	@docker compose exec app go mod tidy