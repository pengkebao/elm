package elm

import (
	"fmt"
)

func main() {
	//创建订单
	orderInfo := NewCreateOrder()
	orderInfo.OrderType = 1
	orderInfo.PartnerOrderCode = "2922"
	orderInfo.NotifyUrl = "https://ckb.mobi"
	TransportInfo := new(TransportInfo)
	TransportInfo.TransportName = "天下第一店"
	TransportInfo.TransportAddress = "大门口"
	TransportInfo.TransportLongitude = "120.285906"
	TransportInfo.TransportLatitude = "30.311656"
	TransportInfo.PositionSource = 3
	TransportInfo.TransportTel = "18354224988"
	orderInfo.TransportInfo = TransportInfo
	orderInfo.OrderTotalAmount = "6"
	orderInfo.OrderActualAmount = "5.5"
	orderInfo.OrderRemark = "用户备注"
	orderInfo.IsInvoiced = 0
	orderInfo.OrderPaymentStatus = 1
	orderInfo.OrderPaymentMethod = 1
	orderInfo.IsAgentPayment = 0
	orderInfo.GoodsCount = 1
	ReceiverInfo := new(ReceiverInfo)
	ReceiverInfo.ReceiverName = "杰哥"
	ReceiverInfo.ReceiverPrimaryPhone = "18658888643"
	ReceiverInfo.ReceiverAddress = "九堡家苑三区"
	ReceiverInfo.ReceiverLongitude = "120.285906"
	ReceiverInfo.ReceiverLatitude = "30.311656"
	ReceiverInfo.PositionSource = 3
	orderInfo.ReceiverInfo = ReceiverInfo
	OrderItem := new(OrderItem)
	OrderItem.ItemName = "商品名称"
	OrderItem.ItemQuantity = 1
	OrderItem.ItemPrice = "2.50"
	OrderItem.ItemActualPrice = "2.40"
	OrderItem.IsNeedPackage = 1
	OrderItem.IsAgentPurchase = 0
	orderInfo.OrderItems = append(orderInfo.OrderItems, OrderItem)
	err := orderInfo.Send()
	//创建门店
	createStore := NewCreateStore()
	createStore.Address = "九堡家苑三区"
	createStore.ChainStoreCode = "292291"
	createStore.ChainStoreName = "天下第一店"
	createStore.ContactPhone = "18658888643"
	createStore.Longitude = "120.285906"
	createStore.Latitude = "30.311656"
	createStore.PositionSource = 3
	createStore.ServiceCode = 1
	err = createStore.Send()
	//查询门店
	QueryStore := NewQueryStore()
	QueryStore.ChainStoreCode = "2922"
	QueryStore.ChainStoreName = "天下第一店"
	QueryStoreRes := new(QueryStoreRes)
	err = QueryStore.Query(QueryStoreRes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(QueryStoreRes)
}
