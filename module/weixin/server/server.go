package server

import (
	"io"
	"net/http"
	"time"

	"gowhole/module/weixin/config"
	"gowhole/module/weixin/context"
	"gowhole/module/weixin/oauth"

	"github.com/davecgh/go-spew/spew"
)

var (
	req  *http.Request
	resp http.ResponseWriter
	c    = context.New(req, resp)
	mux  = http.NewServeMux()
)

//Run for http server.
func Run(config *config.Config) {
	c.APPID = config.APPID
	c.APPSecret = config.APPSecret
	httpFunc("GET", "/callback", oauthHandler)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	panic(s.ListenAndServe())
}

func oauthHandler(w http.ResponseWriter, r *http.Request) {
	_, user, err := oauth.CallBack(c)
	spew.Dump(user, err)
	str := "hello "
	if user != nil {
		str += user.Nickname
	}
	io.WriteString(c.ResponseWriter, str)
}
