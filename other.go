package elm

import "encoding/json"

/**
/anubis-webapi/v2/order/carrier

*/
type Carrier struct {
	PartnerOrderCode string `json:"partner_order_code"`
}

func NewCarrier() *Carrier {
	return &Carrier{}
}

type CarrierQueryRes struct {
	CarrierPhone string  `json:"carrierPhone"`
	CarrierName  string  `json:"carrierName"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

func (this *Carrier) Query(v interface{}) (err error) {
	elm := new(ELM)
	err = elm.carrier(this)
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
