package adapter

import (
	"errors"
	"gowebstarter/infra/db"
	"gowebstarter/model/req"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) (interface{}, error) {
	var loginReq req.LoginReq
	if err := c.BindJSON(&loginReq); err != nil {
		return nil, err
	}
	user, err := db.GetSysUserByName(c, loginReq.Username)
	if err != nil {
		slog.Warn("get user by name failed", "err", err)
		return nil, err
	}
	if user.Password != loginReq.Password {
		slog.Warn("password not match", "user", user)
		// password not match err
		return nil, errors.New("invalid username or password")
	}
	return user, nil
}
