package main

import (
	"net/http"

	"gowhole/aframe/ginweb/dao/weixin"
	"gowhole/aframe/ginweb/model"

	"github.com/gin-gonic/gin"
)

var (
	r    = gin.New()
	err  error
	user *model.UserInfo
)

func init() {
	r.GET("/auth", func(c *gin.Context) {
		weixin.OAuth(c)
	})
	r.GET("/callback", func(c *gin.Context) {
		_, user, err = weixin.CallBack(c)
	})
}

func main() {
	// r.Use(middleware.Logger())
	r.LoadHTMLGlob("../views/**/*")
	r.GET("/", func(c *gin.Context) {
		res := make(gin.H)
		if user != nil {
			res["nickname"] = user.Nickname
		}
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
