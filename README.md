# Coverage for mysql

___

## Usage

* Create .env file and set params
```
MYSQL_DB_USER=
MYSQL_DB_PASSWORD=
MYSQL_DB_HOST=
MYSQL_DB_PORT=
MYSQL_DB_NAME=
MYSQL_DB_DRIVER=
```

* Use Manager
```go
manager := mysql.NewSqlManager()

manager.Exec("...Query")
```
