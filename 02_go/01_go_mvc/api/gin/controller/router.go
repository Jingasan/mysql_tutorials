package controller

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*html")
	router.GET("/", ShowAllBlog)
	router.GET("/show/:id", ShowOneBlog)
	router.GET("/create", ShowCreateBlog)
	router.POST("/create", CreateBlog)
	router.GET("/edit/:id", ShowEditBlog)
	router.POST("/edit", EditBlog)
	router.GET("/delete/:id", ShowCheckDeleteBlog)
	router.POST("/delete", DeleteBlog)
	router.POST("/table", CreateTable)
	router.DELETE("/table", DeleteTable)
	return router
}
