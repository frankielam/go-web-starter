# go-web-starter
Go web starter


### Branches
* http 
```
git checkout http
go run main.go
curl http://127.0.0.1:8080/
```

* gin
```
git checkout gin
go get github.com/gin-gonic/gin 
go run main.go
curl http://127.0.0.1:8080/
```

* gin-adapter
```
git checkout gin-adapter
go get github.com/spf13/viper
go get gopkg.in/natefinch/lumberjack.v2
```

* gin-gorm
```
git checkout gin-gorm
go get gorm.io/gorm
go get gorm.io/driver/postgres
// If using MySQL Database
// go get gorm.io/driver/mysql
```

* gin-redis
```
git checkout gin-redis
go get github.com/redis/go-redis/v9
```