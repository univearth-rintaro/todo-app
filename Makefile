# 環境変数の読み込み
ifneq (,$(wildcard ./env/api.env))
    include ./env/api.env
    export $(shell sed 's/=.*//' ./env/api.env)
endif

# 依存関係のインストール
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	go mod tidy

# サーバのビルド
.PHONY: build
build:
	@echo "Building server..."
	go build -o main ./app/main.go

# サーバの実行
.PHONY: run
run:
	@echo "Running server..."
	./main

# テストの実行
.PHONY: test
test:
	@echo "Running tests..."
	docker-compose run --rm tester

# Dockerコンテナの起動
.PHONY: up
up:
	@echo "Starting Docker containers..."
	docker-compose up -d

# Dockerコンテナの停止
.PHONY: down
down:
	@echo "Stopping Docker containers..."
	docker-compose down

# Swaggerの生成
.PHONY: swagger
swagger:
	@echo "Generating Swagger documentation..."
	swag init -g app/main.go -o docs

# すべてのタスクを実行
.PHONY: all
all: deps build test swagger

# クリーンアップ
.PHONY: clean
clean:
	@echo "Cleaning up..."
	rm -f main
	docker-compose down -v
