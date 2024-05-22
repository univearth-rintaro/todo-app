# Todoアプリケーション

GoとEchoフレームワークを使用したシンプルなTodoアプリケーションです。
バックエンドはGoで、フロントエンドはNext.jsを使用。
データベースはPostgreSQLを使用しています。

## 目次

- [セットアップ](#セットアップ)
- [ビルドと実行](#ビルドと実行)
- [APIドキュメント](#apiドキュメント)
- [ディレクトリ構造](#ディレクトリ構造)
- [貢献](#貢献)
- [ライセンス](#ライセンス)

## セットアップ

### 依存関係のインストール

まず、プロジェクトの依存関係をインストールします。以下のコマンドを実行してください。

```sh
go mod tidy
```

### 環境変数の設定

環境変数を設定するために、.envファイルをプロジェクトのルートディレクトリに作成し、以下のように記述します。

```env
# env/api.env
DB_USER=postgres
DB_PASSWORD=postgres123
DB_NAME=todo
DB_HOST=todo-db
DB_PORT=5432

# env/db.env
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres123
POSTGRES_DB=todo
```

## ビルドと実行

Dockerがインストールされていることを確認し、Dockerコンテナを起動します。

```sh
docker-compose up -d
```

サーバが正常に起動すると、以下のメッセージが表示されます。

```sh
⇨ http server started on [::]:5050
```

### APIドキュメント

APIのドキュメントはSwaggerを使用して生成されています。サーバが起動している状態で、以下のURLにアクセスするとSwagger UIが表示されます。

```bash
http://localhost:5050/swagger/index.html
```

### ディレクトリ構造

プロジェクトのディレクトリ構造は以下の通りです。

```go
.
├── Dockerfile
├── api
│   ├── api_gen.go
│   └── swagger.yaml
├── app
│   ├── main.go
│   ├── middleware
│   │   └── logger.go
│   ├── models
│   │   └── todo.go
│   ├── handlers
│   │   └── todo_handler.go
├── db
│   ├── db.go
│   ├── init.go
│   └── init.sql
├── docker-compose.yml
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── env
│   ├── api.env
│   └── db.env
├── go.mod
├── go.sum
└── tmp
    ├── build-errors.log
    └── main
```

各ディレクトリおよびファイルの役割は以下の通りです。

- **api/** - OpenAPI仕様および生成されたAPIコード
- **app/** - アプリケーションのエントリーポイントおよびミドルウェア、モデル、ハンドラ
- **db/** - データベース接続および初期化スクリプト
- **docker-compose.yml** - Dockerコンテナの設定
- **docs/** - Swaggerドキュメント
- **env/** - 環境変数ファイル
- **tmp/** - 一時ファイル

### 貢献

バグ報告や機能追加の提案は、Issueを通じて行ってください。プルリクエストも歓迎します。

### ライセンス

このプロジェクトは Apache 2.0 ライセンス のもとで公開されています。