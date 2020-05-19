package jdskd

import (
	"sync"

	"github.com/lgs821/jdsdk-go/cache"
	"github.com/lgs821/jdsdk-go/client"
	"github.com/lgs821/jdsdk-go/context"
	"github.com/lgs821/jdsdk-go/oauth"
)

// Jdsdk struct
type Jdsdk struct {
	Context *context.Context
}

// Config for user
type Config struct {
	AppKey    string
	AppSecret string
	ReturnURL string
	SalePlat  string
	Cache     cache.Cache
}

// NewJdsdk init
func NewJdsdk(cfg *Config) *Jdsdk {
	context := new(context.Context)
	copyConfigToContext(cfg, context)
	return &Jdsdk{context}
}

func copyConfigToContext(cfg *Config, context *context.Context) {
	context.AppKey = cfg.AppKey
	context.AppSecret = cfg.AppSecret
	context.ReturnURL = cfg.ReturnURL
	context.SalePlat = cfg.SalePlat
	context.Cache = cfg.Cache
	context.SetAccessTokenLock(new(sync.RWMutex))
}

//GetClient client
func (wc *Jdsdk) GetClient() *client.Client {
	return client.NewClient(wc.Context)
}

//GetAccessToken 获取access_token
func (wc *Jdsdk) GetAccessToken() (string, error) {
	return wc.Context.GetAccessToken()
}

// GetOauth oauth2网页授权
func (wc *Jdsdk) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(wc.Context)
}
