package weixin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"gowhole/aframe/ginweb/model"
	"gowhole/lib/session"
	"gowhole/lib/xrand"

	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

const (
	token           = "mywechat"
	oauth2Host      = "https://open.weixin.qq.com/connect/oauth2/authorize"
	appID           = "wx6825d838f490d674"                   //测试号-wx6825d838f490d674  订阅号-wxc317b6d9e6f1ad6f
	appSecret       = "9769e6449b771846774a6687802829ca"     //测试号-9769e6449b771846774a6687802829ca 订阅号-186ce000a0113c89c8b22722a018edf0
	rediectURI      = "http://dm8vgq.natappfree.cc/callback" //注意：授权回调页面域名是dm8vgq.natappfree.cc
	scope           = "snsapi_userinfo"
	wechatRedirect  = "#wechat_redirect"
	accessTokenHost = "https://sh.api.weixin.qq.com/cgi-bin/token"
	lang            = "zh_CN"
)

var (
	sessionStorage = session.New(20*60, 60*60)
)

//oAuth2URL get oauth url.
func oAuth2URL(state string) string {
	return fmt.Sprintf(oauth2Host+"?appid=%s&redirect_uri=%s&scope=%s&state=%s&response_type=code#wechat_redirect",
		url.QueryEscape(appID),
		url.QueryEscape(rediectURI),
		url.QueryEscape(scope),
		url.QueryEscape(state),
	)
}

//accessTokenURL get oauth url.
func accessTokenURL(code string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		url.QueryEscape(appID),
		url.QueryEscape(appSecret),
		url.QueryEscape(code),
	)
}

func userURL(openid, accessToken string) string {
	return fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?openid=%s&access_token=%s&lang=%s",
		url.QueryEscape(openid),
		url.QueryEscape(accessToken),
		url.QueryEscape(lang),
	)
}

// OAuth 建立必要的 session, 然后跳转到授权页面
func OAuth(c *gin.Context) {
	state := string(xrand.NewHex())
	AuthCodeURL := oAuth2URL(state)
	http.Redirect(c.Writer, c.Request, AuthCodeURL, http.StatusFound)
}

// CallBack 授权后回调页面
func CallBack(c *gin.Context) (code string, user *model.UserInfo, err error) {
	queryValues, err := url.ParseQuery(c.Request.URL.RawQuery)
	if err != nil {
		fmt.Printf("url.ParseQuery %v\n", err)
		return
	}
	code = queryValues.Get("code")
	if code == "" {
		fmt.Printf("CallBack get code empty %v\n", queryValues)
		OAuth(c)
	}
	ac, err := accessTokenOpenID(code)
	if err != nil {
		fmt.Printf("CallBack accessTokenOpenID error %v\n", err)
		return
	}
	if ac == nil {
		fmt.Printf("CallBack ac  %v\n", ac)
		return
	}
	if ac.OpenID == "" || ac.AccessToken == "" {
		fmt.Printf("CallBack OpenID AccessToken empty %v\n", queryValues)
		return
	}
	var valid bool
	if valid, err = isValidToken(ac.OpenID, ac.AccessToken); err != nil {
		fmt.Printf("CallBack isValidToken err %v\n", err)
		return
	}
	if !valid {
		refreshToken(ac.RefreshToken)
	}
	user, err = userInfo(ac)
	return
}

func accessTokenOpenID(code string) (ac *model.AccessToken, err error) {
	var (
		atURL string
	)
	atURL = accessTokenURL(code)
	resp, err := http.Get(atURL)
	if err != nil {
		fmt.Printf("http.Get %v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http.Status %v\n", resp.Status)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if bytes.Contains(body, []byte("errcode")) {
		var errcode *model.ErrorResponse
		err = json.Unmarshal(body, &errcode)
		if err != nil {
			fmt.Printf("get accessToken OpenID %+v\n", errcode)
			return nil, err
		}
		return nil, fmt.Errorf("%s", errcode.ErrMsg)
	}
	if err = json.Unmarshal(body, &ac); err != nil {
		fmt.Printf("json.Unmarshal %v\n", resp.Status)
	}
	spew.Dump("accessTokenOpenID", ac)
	return
}

func userInfo(ac *model.AccessToken) (user *model.UserInfo, err error) {

	l := userURL(ac.OpenID, ac.AccessToken)
	resp, err := http.Get(l)
	if err != nil {
		fmt.Printf("http.Get %v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http.Status %v\n", resp.Status)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if bytes.Contains(body, []byte("errcode")) {
		var errcode *model.ErrorResponse
		err = json.Unmarshal(body, &errcode)
		if err != nil {
			fmt.Printf("get userInfo %+v\n", errcode)
			return nil, err
		}
		return nil, fmt.Errorf("%s", errcode.ErrMsg)
	}
	var result struct {
		u *model.UserInfo
	}
	if err = json.Unmarshal(body, &result.u); err != nil {
		fmt.Printf("json.Unmarshal %v\n", resp.Status)
		return
	}
	user = &model.UserInfo{}
	user = result.u
	return
}

func isValidToken(openid, accessToken string) (isvalid bool, err error) {
	refreshTokenURL := fmt.Sprintf("https://api.weixin.qq.com/sns/auth?openid=%s&access_token=%s",
		url.QueryEscape(openid),
		url.QueryEscape(accessToken),
	)
	resp, err := http.Get(refreshTokenURL)
	if err != nil {
		fmt.Printf("http.Get %v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http.Status %v\n", resp.Status)
		return
	}
	var result *model.ErrorResponse
	body, err := ioutil.ReadAll(resp.Body)
	if err = json.Unmarshal(body, &result); err != nil {
		fmt.Printf("json.Unmarshal %v\n", resp.Status)
	}
	if result.ErrCode == 0 {
		isvalid = true
	}
	spew.Dump("isValidToken", result)
	return
}

func refreshToken(refreshToken string) (ac *model.AccessToken, err error) {
	refreshTokenURL := fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&&refresh_token=%s&grant_type=refresh_token",
		url.QueryEscape(appID),
		url.QueryEscape(refreshToken),
	)
	resp, err := http.Get(refreshTokenURL)
	if err != nil {
		fmt.Printf("http.Get %v\n", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("http.Status %v\n", resp.Status)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if bytes.Contains(body, []byte("errcode")) {
		var errcode *model.ErrorResponse
		err = json.Unmarshal(body, &errcode)
		if err != nil {
			fmt.Printf("refreshToken %+v\n", errcode)
			return nil, err
		}
		return nil, fmt.Errorf("%s", errcode.ErrMsg)
	}
	if err = json.Unmarshal(body, &ac); err != nil {
		fmt.Printf("json.Unmarshal %v\n", resp.Status)
	}
	spew.Dump("refreshToken", ac)
	return
}
