package adapter

import (
	"gowebstarter/infra/config"
	"gowebstarter/model/entity"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

const (
	JwtIdentityKey = "id"
	JwtNameKey     = "name"
)

func ginJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	conf := config.GetConf()
	timeoutDuration := time.Duration(conf.Auth.JWT.Timeout) * time.Second
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Key:         []byte(conf.Auth.JWT.Secret),
		Realm:       conf.Auth.JWT.Realm,
		Timeout:     timeoutDuration,
		MaxRefresh:  time.Duration(conf.Auth.JWT.MaxRefresh) * time.Second,
		IdentityKey: JwtIdentityKey,
		// JWT PayLoad 自定义保存的数据
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*entity.SysUser); ok {
				return jwt.MapClaims{
					JwtIdentityKey: v.ID,
					JwtNameKey:     v.Name,
				}
			}
			return jwt.MapClaims{}
		},

		// Sign in logic
		Authenticator: func(c *gin.Context) (interface{}, error) {
			return login(c) // 登录逻辑
		},
		// 登录失败或 TOKEN 无效
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// 登录成功后
		LoginResponse: func(c *gin.Context, code int, token string, expire time.Time) {
			c.JSON(code, gin.H{
				"code":   code,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		// TOKEN 有效
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			openid := claims[JwtIdentityKey]
			userID := openid.(float64)
			log.Println("userID:", userID)
			// 是否需增加 TOKEN 缓存验证
			return nil
		},
		// after IdentityHandler
		Authorizator: func(data interface{}, c *gin.Context) bool {
			log.Println(data) // return value of IdentityHandler
			return true
		},

		RefreshResponse: nil,
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
	})

	return authMiddleware, err
}

func GetIdFromJWT(c *gin.Context) (int64, error) {
	claims := jwt.ExtractClaims(c)
	openid := claims[JwtIdentityKey]
	userID := openid.(float64)
	return int64(userID), nil
}

func GetNameFromJWT(c *gin.Context) (string, error) {
	claims := jwt.ExtractClaims(c)
	name := claims[JwtNameKey]
	return name.(string), nil
}
