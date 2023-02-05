package controller

import (
	model_db "gin/model/db"
	model_redis "gin/model/redis"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getHome(c *gin.Context) {
	user := model_db.User{}
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	userId := model_redis.GetSession(c, cookieKey)
	if userId != nil {
		user = model_db.GetOneUser(userId.(string))
	}

	c.HTML(http.StatusOK, "home.html", gin.H{
		"user": user,
	})
}
