package oauth

import (
	"fmt"

	"net/http"
	"net/url"

	"gowhole/module/weixin/context"
	"gowhole/module/weixin/model"
	"gowhole/module/weixin/util"
)

const (
	rediectURI     = "http://q9gze9.natappfree.cc/callback" //注意：授权回调页面域名是dm8vgq.natappfree.cc
	wechatRedirect = "#wechat_redirect"
	scope          = "snsapi_userinfo"
	state          = "STATE"
	lang           = "zh_CN"

	accessTokenHost       = "https://sh.api.weixin.qq.com/cgi-bin/token"
	oauthCodeURL          = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&scope=%s&state=%s&response_type=code#wechat_redirect"
	accessTokenURL        = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	validAccessTokenURL   = "https://api.weixin.qq.com/sns/auth?openid=%s&access_token=%s"
	refreshAccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&&refresh_token=%s&grant_type=refresh_token"
	userInfoURL           = "https://api.weixin.qq.com/sns/userinfo?openid=%s&access_token=%s&lang=%s"
)

// OAuth 建立必要的 session, 然后跳转到授权页面
func OAuth(c *context.Context) {
	AuthCodeURL := fmt.Sprintf(oauthCodeURL, c.APPID, url.QueryEscape(rediectURI), scope, state)
	http.Redirect(c.ResponseWriter, c.Request, AuthCodeURL, http.StatusFound)
}

// CallBack 授权后回调页面
func CallBack(c *context.Context) (code string, user *model.UserInfo, err error) {
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
	at, err := accessTokenOpenID(c.APPID, c.APPSecret, code)
	if err != nil {
		fmt.Printf("CallBack accessTokenOpenID at(%v), err(%v)\n", at, err)
		return
	}
	if at != nil {
		var valid bool
		if valid, err = validToken(at.OpenID, at.AccessToken); err != nil {
			fmt.Printf("CallBack isValidToken err %v\n", err)
			return
		}
		if !valid {
			refreshToken(c.APPID, at.RefreshToken)
		}
		user, err = userInfo(at)
	}
	return
}

func accessTokenOpenID(appID, appSecret, code string) (at *model.AccessToken, err error) {
	atURL := fmt.Sprintf(accessTokenURL, appID, appSecret, code)
	resp, err := util.HTTPGet(atURL)
	defer resp.Body.Close()
	var result struct {
		model.Error
		model.AccessToken
	}
	if err = util.DecodeJSONHttpResponse(resp.Body, &result); err != nil {
		return
	}
	if result.ErrCode != model.ErrCodeOK {
		err = &result.Error
		return
	}
	at = &result.AccessToken
	return
}

func validToken(openid, accessToken string) (valid bool, err error) {
	validTokenURL := fmt.Sprintf(validAccessTokenURL, openid, accessToken)
	resp, err := util.HTTPGet(validTokenURL)
	defer resp.Body.Close()
	var result model.Error
	if err = util.DecodeJSONHttpResponse(resp.Body, &result); err != nil {
		return
	}
	switch result.ErrCode {
	case model.ErrCodeOK:
		valid = true
		return
	case 42001, 40001, 40014, 40003:
		valid = false
		return
	default:
		err = &result
		return
	}
}

func refreshToken(appID, refreshToken string) (at *model.AccessToken, err error) {
	refreshTokenURL := fmt.Sprintf(refreshAccessTokenURL, appID, refreshToken)
	resp, err := util.HTTPGet(refreshTokenURL)
	defer resp.Body.Close()
	var result struct {
		model.Error
		model.AccessToken
	}
	if err = util.DecodeJSONHttpResponse(resp.Body, &result); err != nil {
		return
	}
	if result.ErrCode != model.ErrCodeOK {
		err = &result.Error
		return
	}
	at = &result.AccessToken
	return
}

func userInfo(ac *model.AccessToken) (user *model.UserInfo, err error) {
	userURL := fmt.Sprintf(userInfoURL, ac.OpenID, ac.AccessToken, lang)
	resp, err := util.HTTPGet(userURL)
	defer resp.Body.Close()
	var result struct {
		model.Error
		model.UserInfo
	}
	if err = util.DecodeJSONHttpResponse(resp.Body, &result); err != nil {
		return
	}
	if result.ErrCode != model.ErrCodeOK {
		err = &result.Error
		return
	}
	user = &result.UserInfo
	return
}
