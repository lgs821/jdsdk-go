package client

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"time"

	"github.com/lgs821/jdsdk-go/context"
	"github.com/lgs821/jdsdk-go/util"
	"github.com/lgs821/jdsdk-go/waybill"
)

const (
	baseAPIURL = "https://api.jd.com/routerjson"
)

//Client 素材管理
type Client struct {
	*context.Context
}

//NewClient init
func NewClient(context *context.Context) *Client {
	cli := new(Client)
	cli.Context = context
	return cli
}

// //SysParams SysParams
// type SysParams struct {
// 	Method      string `json:"method"`            //是	API接口名称
// 	AccessToken string `json:"access_token"`      //必须	采用OAuth授权方式是必填参数
// 	AppKey      string `json:"app_key"`           //是	应用的app_key
// 	Sign        string `json:"sign"`              //是	详见下文“5.签名算法”描述
// 	Timestamp   string `json:"timestamp"`         //是	时间戳，格式为yyyy-MMDd HH:mm:ss，例如：2019-05-01 00:00:00。京东API服务端允许客户端请求时间误差为10分钟
// 	Format      string `json:"format"`            //是	暂时只支持json
// 	ParamJSON   string `json:"360buy_param_json"` //是
// 	V           string `json:"v"`                 //是	API协议版本，参考接口文档版本
// }

//WaybillReceive 下单接口
func (cli *Client) WaybillReceive(req *waybill.LdopWaybillReceiveRequest) (*waybill.JdWaybillReceiveReturn, error) {
	accessToken, err := cli.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req.SalePlat = cli.SalePlat
	apiparams, _ := json.Marshal(req)
	paramsjson := string(apiparams)
	sysparams := map[string]string{
		"method":            "jingdong.ldop.waybill.receive",
		"access_token":      accessToken,
		"app_key":           cli.AppKey,
		"timestamp":         time.Now().Format("2006-01-02 15:04:05"),
		"360buy_param_json": paramsjson,
		"v":                 "2.0",
	}
	sign := cli.generateSign(sysparams)
	sysparams["sign"] = sign
	fmt.Println(sysparams)
	uri := fmt.Sprintf("%s?method=%s&access_token=%s&app_key=%s&sign=%s&timestamp=%s&360buy_param_json=%s&v=2.0", baseAPIURL, url.QueryEscape(sysparams["method"]), sysparams["access_token"], sysparams["app_key"], url.QueryEscape(sysparams["sign"]), url.QueryEscape(sysparams["timestamp"]), url.QueryEscape(sysparams["360buy_param_json"]))
	fmt.Println(uri)
	responseBytes, err := util.PostJSON(uri, req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(responseBytes))
	var rea *waybill.JdWaybillReceiveReturn
	err = json.Unmarshal(responseBytes, &rea)
	if err != nil {
		return nil, err
	}

	return rea, nil
}

//CancelWayBill 取件单取消
func (cli *Client) CancelWayBill(req *waybill.LdopDeliveryProviderCancelWayBillRequest) (*waybill.LdopDeliveryProviderCancelWayBillReturn, error) {
	accessToken, err := cli.GetAccessToken()
	if err != nil {
		return nil, err
	}
	req.Source = "JOS"
	apiparams, _ := json.Marshal(req)
	paramsjson := string(apiparams)
	sysparams := map[string]string{
		"method":            "jingdong.ldop.delivery.provider.cancelWayBill",
		"access_token":      accessToken,
		"app_key":           cli.AppKey,
		"timestamp":         time.Now().Format("2006-01-02 15:04:05"),
		"360buy_param_json": paramsjson,
		"v":                 "2.0",
	}
	sign := cli.generateSign(sysparams)
	sysparams["sign"] = sign
	fmt.Println(sysparams)
	uri := fmt.Sprintf("%s?method=%s&access_token=%s&app_key=%s&sign=%s&timestamp=%s&360buy_param_json=%s&v=2.0", baseAPIURL, url.QueryEscape(sysparams["method"]), sysparams["access_token"], sysparams["app_key"], url.QueryEscape(sysparams["sign"]), url.QueryEscape(sysparams["timestamp"]), url.QueryEscape(sysparams["360buy_param_json"]))
	fmt.Println(uri)
	responseBytes, err := util.PostJSON(uri, req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(string(responseBytes))
	var rea *waybill.LdopDeliveryProviderCancelWayBillReturn
	err = json.Unmarshal(responseBytes, &rea)
	if err != nil {
		return nil, err
	}

	return rea, nil
}
func (cli *Client) generateSign(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	dataParams := cli.AppSecret
	for _, k := range keys {
		dataParams = dataParams + k + params[k]
	}
	dataParams = dataParams + cli.AppSecret
	r := md5.Sum([]byte(dataParams))
	return strings.ToUpper(hex.EncodeToString(r[:]))
}
