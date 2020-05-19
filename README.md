# 京东青龙系统 SDK for Go

使用Golang开发的京东青龙系统，目前只有下单接口和取消运单接口

## 基本配置

```go
memcache := cache.NewMemcache("127.0.0.1:11211")

config := &jdsdk.Config{
		AppKey:       "your appkey",
		AppSecret:    "your appsecret",
		ReturnURL:    "your retrunurl",
		SalePlat:     "0030001",
		Cache:        memcache,
	}
	jd = jdsdk.NewJdsdk(config)
```

**Cache 设置**

Cache主要用来保存全局access_token以及js-sdk中的ticket： 默认采用memcache存储，建议采用redis长期存储。

```go
redisopts := &cache.RedisOpts{
		Host:     "your redis host",
		Password: "your redis password",
	}
	RedisCache := cache.NewRedis(redisopts)
	config := &jdsdk.Config{
		AppKey:       "your appkey",
		AppSecret:    "your appsecret",
		ReturnURL:    "your retrunurl",
		SalePlat:     "0030001",
		Cache:        RedisCache,
	}
	jd = jdsdk.NewJdsdk(config)
```



## 基本API使用

- [Oauth2授权](https://github.com/lgs821/jdsdk-go#Oauth2授权)
- 发起授权
  - 通过code换取access_token
  - 刷新access_token
  
- [青龙系统](https://github.com/lgs821/jdsdk-go#青龙系统)
  - 下单接口
  - 取消运单

## Oauth2授权

具体流程请参考文档：[获取accesstoken](https://open.jd.com/home/home#/doc/common?listId=880)

**1.发起授权**

```go
oauth := jd.GetOauth()
err := oauth.Redirect("跳转的绝对地址", "snsapi_userinfo", "123dd123")
if err != nil {
	fmt.Println(err)
}
```

> 如果不希望直接跳转，可通过 oauth.GetRedirectURL 获取跳转的url

**2.通过code换取access_token**

```go
code := c.Query("code")
resToken, err := oauth.GetUserAccessToken(code)
if err != nil {
	fmt.Println(err)
	return
}
```

**3.刷新access_token**

```go
func (oauth *Oauth) RefreshAccessToken(refreshToken string) (result ResAccessToken, err error)
```

## 青龙系统

### 下单接口

[下单接口API](https://open.jd.com/home/home#/doc/apiAuthPackage?apiCateId=471&apiId=2122&apiName=jingdong.ldop.waybill.receive)

```go
req := &waybill.LdopWaybillReceiveRequest{OrderID: "your orderid" ......}

cli := jd.GetClient()
res, err := cli.WaybillReceive(req)
```

### 取消运单接口

[取消运单API](https://open.jd.com/home/home#/doc/apiAuthPackage?apiCateId=471&apiId=3482&apiName=jingdong.ldop.delivery.provider.cancelWayBill)

```go
cancelreq := &waybill.LdopDeliveryProviderCancelWayBillRequest{WaybillCode: "your waybillCode", CancelReason: "用户取消订单", OperatorName: "your name"}

cli := jd.GetClient()
res, err := cli.CancelWayBill(cancelreq)
```



## License

Apache License, Version 2.0