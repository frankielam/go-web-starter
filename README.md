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
// scripts/db.sql
curl http://127.0.0.1:8080/
```


* gin-redis
```
git checkout gin-redis
go get github.com/redis/go-redis/v9
```

* gin-jwt
```
git checkout gin-jwt
go get github.com/appleboy/gin-jwt/v2
curl 'http://127.0.0.1:8080/login' --data-raw '{"username":"frankie", "password":"hello"}'
curl -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTg4NTIxMjksImlkIjoxLCJuYW1lIjoiZnJhbmtpZSIsIm9yaWdfaWF0IjoxNzE4NzY1NzMxfQ.I2dpoMEMIXVfKA11tUXOKqSjgEDTnAkk36BfJwCgvnQ' 'http://127.0.0.1:8080/my/profile'
```