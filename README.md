# go-redis-handson
- NIFTY Techbookの「Redis・go-redis速習」のハンズオン用リポジトリです。
- Techbookの内容を実際に動かしてみることで、Redisとgo-redisの基本的な使い方を学ぶことができます。

# 実行方法
- go-redis-handsonを動かすには、以下の環境が必要です。
  - Docker
  - Docker Compose
- DockerとDocker Composeのインストールが完了したら、以下のコマンドを実行して、go-redis-handsonをクローンしてください。
```bash
$ git clone https://github.com/Penpen7/go-redis-handson.git
$ cd go-redis-handson
```
- go-redis-handsonディレクトリに移動したら、以下のコマンドを実行して、Dockerイメージをビルドしてください。
```bash
$ docker-compose build
```
- Dockerイメージのビルドが完了したら、以下のコマンドを実行して、Dockerコンテナを起動してください。
```bash
$ docker-compose up
```

# ハンズオンの進め方
- サンプルコードがAirによってホットリロードされるようになっています。
- main.goを修正すれば即時コードが実行されるので、main.goを修正して色々試してみてください。
