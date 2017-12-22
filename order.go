package elm

import "encoding/json"

type TransportInfo struct {
	TransportName      string `json:"transport_name"`
	TransportAddress   string `json:"transport_address"`
	TransportLongitude string `json:"transport_longitude"`
	TransportLatitude  string `json:"transport_latitude"`
	PositionSource     int64  `json:"position_source"`
	TransportTel       string `json:"transport_tel"`
}

type ReceiverInfo struct {
	ReceiverName         string `json:"receiver_name"`
	ReceiverPrimaryPhone string `json:"receiver_primary_phone"`
	ReceiverAddress      string `json:"receiver_address"`
	ReceiverLongitude    string `json:"receiver_longitude"`
	ReceiverLatitude     string `json:"receiver_latitude"`
	PositionSource       int64  `json:"position_source"`
}

type OrderItem struct {
	ItemName        string `json:"item_name"`
	ItemQuantity    int    `json:"item_quantity"`
	ItemPrice       string `json:"item_price"`
	ItemActualPrice string `json:"item_actual_price"`
	IsNeedPackage   int    `json:"is_need_package"`
	IsAgentPurchase int    `json:"is_agent_purchase"`
}

type CreateOrder struct {
	PartnerOrderCode   string         `json:"partner_order_code"`
	NotifyUrl          string         `json:"notify_url"`
	OrderType          int64          `json:"order_type"`
	TransportInfo      *TransportInfo `json:"transport_info"`
	OrderTotalAmount   string         `json:"order_total_amount"`
	OrderActualAmount  string         `json:"order_actual_amount"`
	OrderRemark        string         `json:"order_remark"`
	IsInvoiced         int64          `json:"is_invoiced"`
	OrderPaymentStatus int64          `json:"order_payment_status"`
	OrderPaymentMethod int64          `json:"order_payment_method"`
	IsAgentPayment     int64          `json:"is_agent_payment"`
	GoodsCount         int64          `json:"goods_count"`
	ReceiverInfo       *ReceiverInfo  `json:"receiver_info"`
	OrderItems         []*OrderItem   `json:"items_json"`
}

func NewCreateOrder() *CreateOrder {
	return &CreateOrder{}
}

func (this *CreateOrder) Send() (err error) {
	elm := new(ELM)
	return elm.createOrder(this)
}

type CancelOrder struct {
	PartnerOrderCode       string `json:"partner_order_code"`
	OrderCancelReasonCode  int64  `json:"order_cancel_reason_code"`
	OrderCancelCode        int64  `json:"order_cancel_code"`
	OrderCancelDescription string `json:"order_cancel_description"`
	OrderCancelTime        int64  `json:"order_cancel_time"`
}

func NewCancelOrder() *CancelOrder {
	return &CancelOrder{}
}

func (this *CancelOrder) Send() (err error) {
	elm := new(ELM)
	return elm.cancelOrder(this)
}

type QueryOrder struct {
	PartnerOrderCode string `json:"partner_order_code"`
}

func NewQueryOrder() *QueryOrder {
	return &QueryOrder{}
}

/**
{
        "transport_station_id": 1234,
        "transport_station_tel": "13112345678",
        "carrier_driver_id": 1,
        "carrier_driver_name": "张三",
        "carrier_driver_phone": "13112345678",
        "estimate_arrive_time": 1469088084266,
        "order_total_delivery_cost": 0,
        "order_total_delivery_discount": 0,
        "order_status": 1,
        "abnormal_code":"ORDER_OUT_OF_DISTANCE_ERROR",
        "abnormal_desc": "订单超区",
        "event_log_details": [
            {
                "order_status": 1,
                "occur_time": 1469088084269,
                "carrier_driver_name": "张三",
                "carrier_driver_phone": "13112345678"
            },
            {
                "order_status": 20,
                "occur_time": 1469088084269,
                "carrier_driver_name": "李四",
                "carrier_driver_phone": "13112345679"
            },
            {
                "order_status": 5,
                "occur_time": 1469088084269,
                "carrier_driver_name": "",
                "carrier_driver_phone": ""
            }
        ]
	}
	**/
type EventLogDetail struct {
	OrderStatus        int64  `json:"order_status"`
	OccurTime          int64  `json:"occur_time"`
	CarrierDriverName  string `json:"carrier_driver_name"`
	CarrierDriverPhone string `json:"carrier_driver_phone"`
}

type QueryOrderRes struct {
	TransportStationId         string            `json:"transport_station_id"`
	TransportStationTel        string            `json:"transport_station_tel"`
	CarrierDriverId            int               `json:"carrier_driver_id"`
	CarrierDriverName          string            `json:"carrier_driver_name"`
	CarrierDriverPhone         string            `json:"carrier_driver_phone"`
	EstimateArriveTime         int64             `json:"estimate_arrive_time"`
	OrderTotalDeliveryCost     int               `json:"order_total_delivery_cost"`
	OrderTotalDeliveryDiscount int               `json:"order_total_delivery_discount"`
	OrderStatus                int               `json:"order_status"`
	AbnormalCode               string            `json:"abnormal_code"`
	AbnormalDesc               string            `json:"abnormal_desc"`
	EventLogDetails            []*EventLogDetail `json:"event_log_details"`
}

func (this *QueryOrder) Query(v interface{}) (err error) {
	elm := new(ELM)
	err = elm.queryOrder(this)
	if err != nil {
		return err
	}
	data, err := json.Marshal(elm.Data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}
	return nil
}

/**
partner_order_code	string	128	是	商户订单号
order_complaint_code	int	3	是	订单投诉编码（230:其他, 150:未保持餐品完整, 160:服务态度恶劣, 190:额外索取费用, 170:诱导收货人或商户退单, 140:提前点击送达, 210:虚假标记异常, 220:少餐错餐, 200:虚假配送, 130:未进行配送）
order_complaint_desc	string	128	否	订单投诉描述（order_complaint_code为230时必填）
order_complaint_time	long	-	是	订单投诉时间（毫秒）
/anubis-webapi/v2/order/complaint
*/

type ComplaintOrder struct {
	PartnerOrderCode   string `json:"partner_order_code"`
	OrderComplaintCode int64  `json:"order_complaint_code"`
	OrderComplaintDesc string `json:"order_complaint_desc"`
	OrderComplaintTime int64  `json:"order_complaint_time"`
}

func NewComplaintOrder() *ComplaintOrder {
	return &ComplaintOrder{}
}

func (this *ComplaintOrder) Send() (err error) {
	elm := new(ELM)
	err = elm.complaintOrder(this)
	return err
}
