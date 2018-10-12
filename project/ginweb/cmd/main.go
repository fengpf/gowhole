package main

import (
	"encoding/json"
	"net/http"

	"gowhole/project/ginweb/model"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/pkg/errors"
)

var (
	r = gin.New()
	// r   = gin.Default()// use Logger(), Recovery()
	err error
)

//JSON def json struct
type JSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

//Render writes data with custom ContentType for implement gin Render interface
func (js JSON) Render(w http.ResponseWriter) error {
	var jb []byte
	if jb, err = json.Marshal(js); err != nil {
		return errors.WithStack(err)
	}
	if _, err = w.Write(jb); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// WriteContentType write json ContentType for implement gin Render interface
func (js JSON) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
}

func myRender(c *gin.Context, code int, r render.Render) {
	r.WriteContentType(c.Writer)
	if code > 0 {
		c.Status(code)
	}

	if err := r.Render(c.Writer); err != nil {
		c.Error(err)
	}
}

// JSONOutput def json output
func JSONOutput(c *gin.Context, data interface{}, err error) {
	code := http.StatusOK
	msg := "ok"
	if err != nil {
		code = 500
		msg = err.Error()
	}

	myRender(c, http.StatusOK, JSON{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func main() {

	pprof.Register(r)

	r.GET("/ping", func(c *gin.Context) {
		JSONOutput(c,
			map[string]interface{}{
				"a": 111,
				"b": "vvcvv",
			},
			nil,
		)
	})

	r.GET("/hello", func(c *gin.Context) {
		JSONOutput(c,
			"world",
			nil,
		)
	})

	r.GET("/error", func(c *gin.Context) {
		JSONOutput(c,
			"xxxxx",
			errors.New("bad req"),
		)
	})

	r.Run(":9000")
}

func test() {

	r.LoadHTMLGlob("../views/**/*")
	r.GET("/", func(c *gin.Context) {
		res := make(gin.H)
		res["nickname"] = "nickname"
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
}
