# mapkicker

## Environment setup

You need to have [Go](https://golang.org/),
[Node.js](https://nodejs.org/),
[Docker](https://www.docker.com/), and
[Docker Compose](https://docs.docker.com/compose/)
(comes pre-installed with Docker on Mac and Windows)
installed on your computer.

Verify the tools by running the following commands:

```sh
go version
npm --version
docker --version
docker-compose --version
```

### Goの設定

* `GOPATH`環境変数を設定します。 `~/go`などにするのが標準的です。
* `PATH`に`$GOPATH/bin`を追加します。これによって、`go get`でインストールしたツールにpathが通ります。

### ツール

#### goreman (開発モードでのサーバ起動を簡略化する)

```
go get github.com/mattn/goreman
```

#### gin(バックエンドのライブリロードに用いる)

```
go get github.com/codegangsta/gin
```

## Development modeで起動する

### 古い手順 

#### frontendの依存ライブラリ(毎回やる必要はない)

```sh
cd webapp
npm install
```

#### 実行

プロジェクトルートにおいて,

```sh
goreman start
```

## Production modeで起動する

Perform:
```sh
docker-compose up
```
This will build the application and start it together with
its database. Access the application on http://localhost:8080.
