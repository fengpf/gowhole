package config

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

//Config for one wx Official Accounts.
type Config struct {
	APPID     string
	APPSecret string
	Token     string
}

//New for wx config
func New() *Config {
	return &Config{
		APPID:     appID,
		APPSecret: appSecret,
		Token:     token,
	}
}
