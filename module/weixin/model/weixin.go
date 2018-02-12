package model

//AccessToken for oauth2.0.
type AccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	ExpiresIn    int64  `json:"expires_in"`
}

//ErrorResponse for wx.
type ErrorResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

//UserInfo for wx.
type UserInfo struct {
	IsSubscriber int    `json:"subscribe"` // 用户是否订阅该公众号标识, 值为0时, 代表此用户没有关注该公众号, 拉取不到其余信息
	OpenID       string `json:"openid"`    // 用户的标识, 对当前公众号唯一
	Nickname     string `json:"nickname"`  // 用户的昵称
	Sex          int    `json:"sex"`       // 用户的性别, 值为1时是男性, 值为2时是女性, 值为0时是未知
	Language     string `json:"language"`  // 用户的语言, zh_CN, zh_TW, en
	City         string `json:"city"`      // 用户所在城市
	Province     string `json:"province"`  // 用户所在省份
	Country      string `json:"country"`   // 用户所在国家
	// 用户头像, 最后一个数值代表正方形头像大小(有0, 46, 64, 96, 132数值可选, 0代表640*640正方形头像), 用户没有头像时该项为空
	HeadImageURL  string `json:"headimgurl"`
	SubscribeTime int64  `json:"subscribe_time"`    // 用户关注时间, 为时间戳. 如果用户曾多次关注, 则取最后关注时间
	UnionID       string `json:"unionid,omitempty"` // 只有在用户将公众号绑定到微信开放平台帐号后, 才会出现该字段.
	Remark        string `json:"remark"`            // 公众号运营者对粉丝的备注, 公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupID       int64  `json:"groupid"`           // 用户所在的分组ID
	TagIDList     []int  `json:"tagid_list"`        // Tag List
}
