package main

import "github.com/pengkebao/elm"
import "fmt"

func main() {
	//创建订单
	// orderInfo := elm.NewCreateOrder()
	// orderInfo.OrderType = 1
	// orderInfo.PartnerOrderCode = "2922"
	// orderInfo.NotifyUrl = "https://ckb.mobi"
	// TransportInfo := new(elm.TransportInfo)
	// TransportInfo.TransportName = "天下第一店"
	// TransportInfo.TransportAddress = "大门口"
	// TransportInfo.TransportLongitude = "120.285906"
	// TransportInfo.TransportLatitude = "30.311656"
	// TransportInfo.PositionSource = 3
	// TransportInfo.TransportTel = "18354224988"
	// orderInfo.TransportInfo = TransportInfo
	// orderInfo.OrderTotalAmount = "6"
	// orderInfo.OrderActualAmount = "5.5"
	// orderInfo.OrderRemark = "用户备注"
	// orderInfo.IsInvoiced = 0
	// orderInfo.OrderPaymentStatus = 1
	// orderInfo.OrderPaymentMethod = 1
	// orderInfo.IsAgentPayment = 0
	// orderInfo.GoodsCount = 1
	// ReceiverInfo := new(elm.ReceiverInfo)
	// ReceiverInfo.ReceiverName = "杰哥"
	// ReceiverInfo.ReceiverPrimaryPhone = "18658888643"
	// ReceiverInfo.ReceiverAddress = "九堡家苑三区"
	// ReceiverInfo.ReceiverLongitude = "120.285906"
	// ReceiverInfo.ReceiverLatitude = "30.311656"
	// ReceiverInfo.PositionSource = 3
	// orderInfo.ReceiverInfo = ReceiverInfo
	// OrderItem := new(elm.OrderItem)
	// OrderItem.ItemName = "商品名称"
	// OrderItem.ItemQuantity = 1
	// OrderItem.ItemPrice = "2.50"
	// OrderItem.ItemActualPrice = "2.40"
	// OrderItem.IsNeedPackage = 1
	// OrderItem.IsAgentPurchase = 0
	// orderInfo.OrderItems = append(orderInfo.OrderItems, OrderItem)
	// err := orderInfo.Send()
	// fmt.Println(err)
	//查询订单
	// query := elm.NewQueryOrder()
	// query.PartnerOrderCode = "2922"
	// res := new(elm.QueryOrderRes)

	// err := query.Query(res)
	// if err == nil {
	// 	fmt.Println(res.EventLogDetails[1])
	// }
	// fmt.Println(err)

	//取消订单
	// cancleOrder := elm.NewCancelOrder()
	// cancleOrder.PartnerOrderCode = "2922"
	// cancleOrder.OrderCancelDescription = "不要了"
	// cancleOrder.OrderCancelCode = 0
	// cancleOrder.OrderCancelReasonCode = 2
	// cancleOrder.OrderCancelTime = (time.Now().Unix() * 1000)
	// err := cancleOrder.Send()
	// fmt.Println(err)

	//投诉订单
	// ComplaintOrder := elm.NewComplaintOrder()
	// ComplaintOrder.PartnerOrderCode = "2922"
	// ComplaintOrder.OrderComplaintDesc = "送了一天还没到"
	// ComplaintOrder.OrderComplaintTime = (time.Now().Unix() * 1000)
	// ComplaintOrder.OrderComplaintCode = 200

	// err := ComplaintOrder.Send()
	// fmt.Println(err)

	// //创建门店
	// createStore := elm.NewCreateStore()
	// createStore.Address = "九堡家苑三区"
	// createStore.ChainStoreCode = "292291"
	// createStore.ChainStoreName = "天下第一店"
	// createStore.ContactPhone = "18658888643"
	// createStore.Longitude = "120.285906"
	// createStore.Latitude = "30.311656"
	// createStore.PositionSource = 3
	// createStore.ServiceCode = 1
	// err = createStore.Send()
	//查询门店
	// QueryStore := elm.NewQueryStore()
	// QueryStore.ChainStoreCode = append(QueryStore.ChainStoreCode, "A001")
	// QueryStore.ChainStoreName = append(QueryStore.ChainStoreCode, "饿了么BOD7")
	// var QueryStoreRes []elm.QueryStoreRes
	// err := QueryStore.Query(&QueryStoreRes)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(QueryStoreRes)

	//骑手位置
	Carrier := new(elm.Carrier)
	Carrier.PartnerOrderCode = "2922"
	CarrierRes := new(elm.CarrierQueryRes)
	err := Carrier.Query(CarrierRes)
	fmt.Println(err)
	fmt.Println(CarrierRes)
}
