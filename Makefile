.DEFAULT_GOAL := help

build-up: ## Build docker image and up container
	docker compose up -d --build

up: ## コンテナ起動
	docker compose up -d

down: ## コンテナダウン
	docker compose down

in: ## Appのコンテナに入る
	docker compose exec app sh

logs: ## ログ確認
	docker compose logs -f

ps: ## コンテナステータスの確認
	docker compose ps

dry-migrate: ## マイグレーションテスト実行(実行されない)
	mysqldef -u ${DB_USER} -p ${DB_PASSWORD} -h ${DB_HOST} -P ${DB_PORT} ${DB_NAME} --dry-run < ./_tools/mysql/schema.sql

migrate:  ## マイグレーション実行
	mysqldef -u ${DB_USER} -p ${DB_PASSWORD} -h ${DB_HOST} -P ${DB_PORT} ${DB_NAME} < ./_tools/mysql/schema.sql

sqlboiler: ## SQLBoilerでのモデル自動生成
	sqlboiler mysql -c config/sqlboiler.toml  -o models -p models --no-tests --wipe
	
.PHONY: moq
moq: ## mockの作成(コンテナ内で実行すること)
	# サービス層のモック作成
	@docker compose exec app moq -fmt goimports -out ./handler/moq_test.go ./handler \
		RegisterTemporaryUserService

help: ## コマンド説明一覧の表示
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'