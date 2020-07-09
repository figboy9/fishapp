## minikube始め方
```bash
make start
```

## minikube終わり方
```bash
make stop
```

## skaffold始め方
```bash
make dev
```
青文字でDeployments stabilized in ~　と出るまで待つ。
## skaffold終わり方
ctrl + C

正常に終了するまで待つ。
## DBデータなどを削除する
skaffoldが終了しているのを確認して、
```bash
make rmdata
```
## GraphQL接続方法
8080ポートを使用中だとエラー出る。
#### GraphQLPlayground
クエリを試すことができる。

URL: localhost:8080/playground

#### GraphQLQuery
実際にアプリでクエリを投げるところ

URL: localhost:8080/graphql

## DB接続方法
ローカルのポートを使用済みだとエラー出る。
| dbname   | host | port | username | password |
|:-----------|------------:|:------------:|:------------:|:------------:|
| user_DB   |     localhost |   4306       | root | password
| chat_DB    |        localhost |    5306        | root|password
| image_DB   |     localhost |   6306     |root |password
| post_DB    |      localhost |    7306  |  root |password
