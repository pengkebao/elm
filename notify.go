package elm

import (
	"encoding/json"
	"net/url"
)

type Notify struct {
	AppId     string `json:"app_id"`
	Data      string `json:"data"`
	Salt      int    `json:"salt"`
	Signature string `json:"signature"`
}

/**
body 网页原生内容
v 返回数据
返回数据参考
蜂鸟配送"系统已接单"状态返回JSON数据如下:

{
    "partner_order_code": "BG658907200991",
    "order_status": 1,
    "push_time": 1466095163344,
    "carrier_driver_name": "",
    "carrier_driver_phone": "",
    "description": "",
    "error_code":""
}
订单"已分配骑手"状态返回JSON数据如下:

{
    "partner_order_code": "BG658907200991",
    "order_status": 20,
    "push_time": 1466095163344,
    "carrier_driver_name": "张三",
    "carrier_driver_phone": "18602030493",
    "description": "",
    "station_name":"亦庄配送站",
    "station_tel":"18602393333",
    "error_code":""
}
订单"骑手已到店"状态返回JSON数据如下:

{
    "partner_order_code": "BG658907200991",
    "order_status": 80,
    "push_time": 1466095163344,
    "carrier_driver_name": "张三",
    "carrier_driver_phone": "18602030493",
    "description": "",
    "station_name":"亦庄配送站",
    "station_tel":"18602393333",
    "error_code":""
}
订单"配送中"的状态返回JSON数据如下:

{
    "partner_order_code": "BG659915200312",
    "order_status": 2,
    "push_time": 1466129638461,
    "carrier_driver_name": "周智伟",
    "carrier_driver_phone": "18501336429",
    "description": "",
    "error_code":""
}
订单"已送达"状态返回JSON数据如下:

{
    "partner_order_code": "BG659915200312",
    "order_status": 3,
    "push_time": 1466132529983,
    "carrier_driver_name": "周智伟",
    "carrier_driver_phone": "18501336429",
    "cancel_reason": 1,
    "description": "",
    "error_code":""
}
订单"系统拒单/配送异常"状态返回JSON数据如下:

{
    "partner_order_code": "BG659915200312",
    "order_status": 5,
    "push_time": 1466132529983,
    "carrier_driver_name": "",
    "carrier_driver_phone": "",
    "description": "订单重复",
    "error_code":"ORDER_REPETITION"
}
参数	数据类型	说明
partner_order_code	string	商户自己的订单号
order_status	int	状态码
push_time	long	状态推送时间(毫秒)
carrier_driver_name	string	蜂鸟配送员姓名
carrier_driver_phone	string	蜂鸟配送员电话
description	string	描述信息
address	string	定点次日达服务独有的字段: 微仓地址
latitude	long	定点次日达服务独有的字段: 微仓纬度
longitude	long	定点次日达服务独有的字段: 微仓经度
cancel_reason	int	订单取消原因. 1:用户取消, 2:商家取消
error_code	string	错误编码
*/
func NewNotify() *Notify {
	return &Notify{}
}

func (this *Notify) Check(body []byte, v interface{}) (err error) {
	err = json.Unmarshal(body, this)
	if err != nil {
		return err
	}
	//验证签名
	elm := new(ELM)
	err = elm.verifSign(this)
	if err != nil {
		return err
	}
	data, err := url.QueryUnescape(this.Data)
	if err != nil {
		return err
	}
	//返回数据
	err = json.Unmarshal([]byte(data), v)
	if err != nil {
		return err
	}
	return err
}
