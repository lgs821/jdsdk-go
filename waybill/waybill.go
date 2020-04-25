package waybill

//LdopWaybillReceiveRequest 添加请求字段
type LdopWaybillReceiveRequest struct {
	SalePlat             string  `json:"salePlat"`             //是	无	销售平台
	CustomerCode         string  `json:"customerCode"`         //是	无	商家编码
	OrderID              string  `json:"orderId"`              //是	无	订单号
	ThrOrderID           string  `json:"thrOrderId"`           //否	无	销售平台订单号(例如京东订单号或天猫订单号等等。总长度不要超过100。如果有多个单号，用英文,间隔。例如：7898675,7898676)
	SenderName           string  `json:"senderName"`           //是	无	寄件人姓名，说明：不能为生僻字
	SenderAddress        string  `json:"senderAddress"`        //是	无	寄件人地址，说明：不能为生僻字
	SenderTel            string  `json:"senderTel"`            //否	无	寄件人电话
	SenderMobile         string  `json:"senderMobile"`         //否	无	寄件人手机(寄件人电话、手机至少有一个不为空)
	SenderPostcode       string  `json:"senderPostcode"`       //否	100000	寄件人邮编，长度：6位
	ReceiveName          string  `json:"receiveName"`          //是	无	收件人名称，说明：不能为生僻字
	ReceiveAddress       string  `json:"receiveAddress"`       //是	无	收件人地址，说明：不能为生僻字
	Province             string  `json:"province"`             //否	无	收件人省
	City                 string  `json:"city"`                 //否	无	收件人市
	County               string  `json:"county"`               //否	无	收件人县
	Town                 string  `json:"town"`                 //否	无	收件人镇
	ProvinceID           int     `json:"provinceId"`           //否	无	收件人省编码
	CityID               int     `json:"cityId"`               //否	无	收件人市编码
	CountyID             int     `json:"countyId"`             //否	无	收件人县编码
	TownID               int     `json:"townId"`               //否	无	收件人镇编码
	SiteType             int     `json:"siteType"`             //否	无	站点类型
	SiteID               int     `json:"siteId"`               //否	无	站点编码
	SiteName             string  `json:"siteName"`             //否	无	站点名称
	ReceiveTel           string  `json:"receiveTel"`           //否	无	收件人电话
	ReceiveMobile        string  `json:"receiveMobile"`        //否	无	收件人手机号(收件人电话、手机至少有一个不为空)
	Postcode             string  `json:"postcode"`             //否	100000	收件人邮编，长度：6
	PackageCount         int     `json:"packageCount"`         //是	无	包裹数(大于0，小于1000)
	Weight               float64 `json:"weight"`               //是	2.5	重量(单位：kg，保留小数点后两位)
	VloumLong            float64 `json:"vloumLong"`            //否	无	包裹长(单位：cm,保留小数点后两位)
	VloumWidth           float64 `json:"vloumWidth"`           //否	无	包裹宽(单位：cm，保留小数点后两位)
	VloumHeight          float64 `json:"vloumHeight"`          //否	无	包裹高(单位：cm，保留小数点后两位)
	Vloumn               float64 `json:"vloumn"`               //是	10000	体积(单位：cm3，保留小数点后两位)
	Description          string  `json:"description"`          //否	无	商品描述
	CollectionValue      int     `json:"collectionValue"`      //否	1	是否代收货款(是：1，否：0。不填或者超出范围，默认是0)
	CollectionMoney      float64 `json:"collectionMoney"`      //否	98.00	代收货款金额(保留小数点后两位)
	GuaranteeValue       int     `json:"guaranteeValue"`       //否	1	是否保价(是：1，否：0。不填或者超出范围，默认是0)
	GuaranteeValueAmount float64 `json:"guaranteeValueAmount"` //否	100.00	保价金额(保留小数点后两位)
	SignReturn           int     `json:"signReturn"`           //否	1	签单返还(签单返还类型：0-不返单，1-普通返单，2-校验身份返单，3-电子签返单，4-电子返单+普通返单)
	Aging                int     `json:"aging"`                //否	1	时效(普通：1，工作日：2，非工作日：3，晚间：4。O2O一小时达：5。O2O定时达：6。不填或者超出范围，默认是1)
	TransType            int     `json:"transType"`            //否	1	运输类型(陆运：1，航空：2。不填或者超出范围，默认是1)
	Remark               string  `json:"remark"`               //否	无	运单备注，长度：20,说明：打印面单时备注内容也会显示在快递面单上
	GoodsType            int     `json:"goodsType"`            //否	无	配送业务类型（ 1:普通，3:填仓，4:特配，6:控温，7:冷藏，8:冷冻，9:深冷）默认是1
	OrderType            int     `json:"orderType"`            //否	无	运单类型。(普通外单：0，O2O外单：1)默认为0
	ShopCode             string  `json:"shopCode"`             //否	无	门店编码(只O2O运单需要传，普通运单不需要传)
	OrderSendTime        string  `json:"orderSendTime"`        //否	2014-09-18 08:30:00	预约配送时间（格式：yyyy-MMDd HH:mm:ss）
	WarehouseCode        string  `json:"warehouseCode"`        //否	无	发货仓编码
	AreaProvID           int     `json:"areaProvId"`           //否	无	接货省ID
	AreaCityID           int     `json:"areaCityId"`           //否	无	接货市ID
	ShipmentStartTime    string  `json:"shipmentStartTime"`    //否	1	配送起始时间
	ShipmentEndTime      string  `json:"shipmentEndTime"`      //否	1	配送结束时间
	Idint                string  `json:"idint"`                //否	无	身份证号
	AddedService         string  `json:"addedService"`         //否	无	扩展服务
	ExtendField1         string  `json:"extendField1"`         //否	无	扩展字段1
	ExtendField2         string  `json:"extendField2"`         //否	无	扩展字段2
	ExtendField3         string  `json:"extendField3"`         //否	无	扩展字段3
	ExtendField4         int     `json:"extendField4"`         //否	无	扩展字段4
	ExtendField5         int     `json:"extendField5"`         //否	无	扩展字段5
	SenderCompany        string  `json:"senderCompany"`        //否	北京市大兴区亦庄经济开发区京东大厦	寄件人公司，长度：50，说明：不能为生僻字
	ReceiveCompany       string  `json:"receiveCompany"`       //否	北京市大兴区亦庄经济开发区京东大厦	收件人公司，长度：50，说明：不能为生僻字
	Goods                string  `json:"goods"`                //否	服装、3C等	托寄物名称，长度：200，说明：为避免托运物品在铁路、航空安检中被扣押，请务必下传商品类目或名称，例如：服装、3C等
	GoodsCount           int     `json:"goodsCount"`           //否	无	寄托物数量
	PromiseTimeType      int     `json:"promiseTimeType"`      //否	无	产品类型（1：特惠送 2：特快送 4：城际闪送 5：同城当日达 6：次晨达 7：微小件 8: 生鲜专送 16：生鲜特快 17、生鲜特惠 21：特惠小包）
	Freight              float64 `json:"freight"`              //否	无	运费
	PickUpStartTime      string  `json:"pickUpStartTime"`      //否	1	预约取件开始时间
	PickUpEndTime        string  `json:"pickUpEndTime"`        //否	1	预约取件结束时间
	UnpackingInspection  string  `json:"unpackingInspection"`  //否	无	开箱验货标识
	BoxCode              string  `json:"boxCode"`              //否	无	商家箱号,多个箱号请用逗号分隔，例如三个箱号为：a123,b456,c789
	FileURL              string  `json:"fileUrl"`              //否	无	文件url
}

//JdWaybillReceiveReturn aa
type JdWaybillReceiveReturn struct {
	JingdongLdopWaybillReceiveResponce JingdongLdopWaybillReceiveResponce `json:"jingdong_ldop_waybill_receive_responce"`
}

//JingdongLdopWaybillReceiveResponce aa
type JingdongLdopWaybillReceiveResponce struct {
	Code                   string                 `json:"code"`
	ReceiveorderinfoResult ReceiveorderinfoResult `json:"receiveorderinfo_result"`
}

//ReceiveorderinfoResult 返回数据
type ReceiveorderinfoResult struct {
	ResultCode      int           `json:"resultCode"`      //无	结果编码
	ResultMessage   string        `json:"resultMessage"`   //无	结果描述
	OrderID         string        `json:"orderId"`         //无	订单号
	DeliveryID      string        `json:"deliveryId"`      //无	运单号
	PromiseTimeType int           `json:"promiseTimeType"` //无	产品类型（1：特惠送 2：特快送 4：城际闪送 5：同城当日达 6：次晨达 7：微小件 8: 生鲜专送 16：生鲜特快 17:生鲜特惠）
	PreSortResult   PreSortResult `json:"preSortResult"`   //预分拣结果：面单打印信息
	TransType       int           `json:"transType"`       //无	运输类型（1：陆运 2：航空）
	NeedRetry       bool          `json:"needRetry"`       //无	是否需要重试

}

//PreSortResult 面单数据
type PreSortResult struct {
	SiteID                 int    `json:"siteId"`                 //无	站点id
	SiteName               string `json:"siteName"`               //无	站点名称
	Road                   string `json:"road"`                   //无	路区
	SlideNo                string `json:"slideNo"`                //无	滑道号
	SourceSortCenterID     int    `json:"sourceSortCenterId"`     //无	始发分拣中心id
	SourceSortCenterName   string `json:"sourceSortCenterName"`   //无	始发分拣中心名称
	SourceCrossCode        string `json:"sourceCrossCode"`        //无	始发道口号
	SourceTabletrolleyCode string `json:"sourceTabletrolleyCode"` //无	始发笼车号
	TargetSortCenterID     int    `json:"targetSortCenterId"`     //无	目的分拣中心id
	TargetSortCenterName   string `json:"targetSortCenterName"`   //无	目的分拣中心名称
	TargetTabletrolleyCode string `json:"targetTabletrolleyCode"` //无	目的笼车号
	Aging                  int    `json:"aging"`                  //无	时效
	AgingName              string `json:"agingName"`              //无	时效名称
	SiteType               int    `json:"siteType"`               //无	站点类型
	IsHideName             int    `json:"isHideName"`             //无	是否隐藏姓名
	IsHideContractints     int    `json:"isHideContractints"`     //无	是否隐藏联系方式
}

//LdopDeliveryProviderCancelWayBillRequest 取件单取消
type LdopDeliveryProviderCancelWayBillRequest struct {
	UserPin      string `json:"userPin"`      //否	pin	用户唯一标识,与下单一致可不填写
	WaybillCode  string `json:"waybillCode"`  //是	VA12345678	运单号
	CustomerCode string `json:"customerCode"` //是	010K00000	商家编码
	Source       string `json:"source"`       //是	JOS	来源
	CancelReason string `json:"cancelReason"` //是	客户取消	取消原因
	OperatorName string `json:"operatorName"` //是	张三	操作人
}

//LdopDeliveryProviderCancelWayBillReturn aa
type LdopDeliveryProviderCancelWayBillReturn struct {
	LdopDeliveryProviderCancelWayBillResponse LdopDeliveryProviderCancelWayBillResponse `json:"jingdong_ldop_delivery_provider_cancelWayBill_responce"`
}

//LdopDeliveryProviderCancelWayBillResponse aa
type LdopDeliveryProviderCancelWayBillResponse struct {
	Code                     string                   `json:"code"`
	CancelWayBillresponseDTO CancelWayBillresponseDTO `json:"responseDTO"`
}

//CancelWayBillresponseDTO bb
type CancelWayBillresponseDTO struct {
	StatusCode    int    `json:"statusCode"`
	StatusMessage string `json:"statusMessage"`
}
