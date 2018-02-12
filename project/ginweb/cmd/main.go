package main

import (
	"net/http"

	"gowhole/aframe/ginweb/model"

	"github.com/gin-gonic/gin"
)

var (
	r   = gin.New()
	err error
)

func main() {
	// r.Use(middleware.Logger())
	r.LoadHTMLGlob("../views/**/*")
	r.GET("/", func(c *gin.Context) {
		res := make(gin.H)
		res["nickname"] = "not get nickname"
		c.HTML(http.StatusOK, "index.html", res)
	})
	r.GET("/user/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user/index.html", gin.H{
			"title": "user",
		})
	})
	r.POST("/login", func(c *gin.Context) {
		var form model.LoginForm
		if c.Bind(&form) == nil {
			if form.User == "user" && form.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})
	r.Run(":8080")
}
