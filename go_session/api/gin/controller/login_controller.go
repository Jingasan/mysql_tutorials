package controller

import (
	model_db "gin/model/db"
	model_redis "gin/model/redis"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getSignup(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func postSignup(c *gin.Context) {
	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	user, err := model_db.Signup(id, pw)
	if err != nil {
		c.Redirect(http.StatusFound, "/signup")
		return
	}
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	model_redis.NewSession(c, cookieKey, user.UserId)
	c.Redirect(http.StatusFound, "/")
}

func getLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func postLogin(c *gin.Context) {
	id := c.PostForm("user_id")
	pw := c.PostForm("password")
	user, err := model_db.Login(id, pw)
	if err != nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	model_redis.NewSession(c, cookieKey, user.UserId)
	c.Redirect(http.StatusFound, "/")
}

func getLogout(c *gin.Context) {
	cookieKey := os.Getenv("LOGIN_USER_ID_KEY")
	model_redis.DeleteSession(c, cookieKey)
	c.Redirect(http.StatusFound, "/")
}
