# CONFIG


- 優先順位
1. コマンド引数
1. 設定ファイル
1. 環境変数
1. デフォルト

- 使用例
- mysql user
1. コマンド引数
```sh
program.exe --mysql-user root
```
1. 設定ファイル
```json
{
    "mysql": {
        "user": "root"
    }
}
```
```yaml
mysql:
    user: root
```
1. 環境変数
```sh
export MYSQL_USER=root
```
1. デフォルト
```go
getConfig(["mysql","user"],"root")
```







