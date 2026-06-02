.PHONY: up down down-clean test-run check-es check-os

# Docker起動
up:
	docker-compose up -d
	@echo "Waiting for services to be ready..."
	@sleep 5

# Docker停止
down:
	docker-compose down

# クリーンアップ（コンテナとボリューム削除）
down-clean:
	docker-compose down -v

# テスト実行
test-run:
	cd scripts && go run .

# Elasticsearch ヘルスチェック
check-es:
	curl -s http://localhost:9200/_cluster/health?pretty

# OpenSearch ヘルスチェック
check-os:
	curl -s http://localhost:9201/_cluster/health?pretty