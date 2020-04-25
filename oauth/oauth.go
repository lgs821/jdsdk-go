package oauth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/lgs821/jdsdk-go/context"
	"github.com/lgs821/jdsdk-go/util"
)

const (
	redirectOauthURL      = "https://open-oauth.jd.com/oauth2/to_login?app_key=%s&response_type=code&redirect_uri=%s&scope=%s&state=%s"
	accessTokenURL        = "https://open-oauth.jd.com/oauth2/access_token?app_key=%s&app_secret=%s&grant_type=authorization_code&code=%s"
	refreshAccessTokenURL = "https://open-oauth.jd.com/oauth2/refresh_token?app_key=%s&app_secret=%s&grant_type=refresh_token&refresh_token=%s"
)

//Oauth 保存用户授权信息
type Oauth struct {
	*context.Context
}

//NewOauth 实例化授权信息
func NewOauth(context *context.Context) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	return auth
}

//GetRedirectURL 获取跳转的url地址
func (oauth *Oauth) GetRedirectURL(redirectURI, scope, state string) (string, error) {
	//url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(redirectOauthURL, oauth.AppKey, urlStr, scope, state), nil
}

//Redirect 跳转到网页授权
func (oauth *Oauth) Redirect(writer http.ResponseWriter, req *http.Request, redirectURI, scope, state string) error {
	location, err := oauth.GetRedirectURL(redirectURI, scope, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, 302)
	return nil
}

// ResAccessToken 获取用户授权access_token的返回结果
type ResAccessToken struct {
	Code         int64  `json:"code"`
	Msg          string `json:"msg"`
	RequestID    string `json:"requestId"`
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}

// GetUserAccessToken 通过网页授权的code 换取access_token
func (oauth *Oauth) GetUserAccessToken() (accessToken string, err error) {
	accessTokenCacheKey := fmt.Sprintf("jd_access_token_%s", oauth.AppKey)
	val := oauth.Cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}
	return
}

// GetUserAccessTokenFromServer 通过网页授权的code 换取access_token
func (oauth *Oauth) GetUserAccessTokenFromServer(code string) (result ResAccessToken, err error) {

	urlStr := fmt.Sprintf(accessTokenURL, oauth.AppKey, oauth.AppSecret, code)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Code != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.Code, result.Msg)
		return
	}
	accessTokenCacheKey := fmt.Sprintf("jd_access_token_%s", oauth.AppKey)
	expires := result.ExpiresIn - 1500
	err = oauth.Cache.Set(accessTokenCacheKey, result.AccessToken, time.Duration(expires)*time.Second)
	return
}

//RefreshAccessToken 刷新access_token
func (oauth *Oauth) RefreshAccessToken(refreshToken string) (result ResAccessToken, err error) {
	urlStr := fmt.Sprintf(refreshAccessTokenURL, oauth.AppKey, oauth.AppSecret, refreshToken)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Code != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.Code, result.Msg)
		return
	}
	return
}
