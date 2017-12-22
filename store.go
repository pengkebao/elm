package elm

import "encoding/json"

type CreateStore struct {
	ChainStoreCode string `json:"chain_store_code"`
	ChainStoreName string `json:"chain_store_name"`
	ContactPhone   string `json:"contact_phone"`
	Address        string `json:"address"`
	PositionSource int    `json:"position_source"`
	Longitude      string `json:"longitude"`
	Latitude       string `json:"latitude"`
	ServiceCode    int    `json:"service_code"`
}

func NewCreateStore() *CreateStore {
	return &CreateStore{}
}

func (this *CreateStore) Send() (err error) {
	elm := new(ELM)
	return elm.createStore(this)
}

/**
chain_store_code	string	32	否	门店编码集合，JSON格式字符串（支持数字,字母的组合）
chain_store_name	string	32	是	门店名称集合，JSON格式字符串（后期废弃）
*/
type QueryStore struct {
	ChainStoreCode []string `json:"chain_store_code"`
	ChainStoreName []string `json:"chain_store_name"`
}

func NewQueryStore() *QueryStore {
	return &QueryStore{}
}

/**
"data": [
        {
            "chain_store_code": "A001",
            "chain_store_name": "饿了么BOD7",
            "address": "300弄亚都国际名园5号楼2003室",
            "latitude": "30.6865430000",
            "longitude": "104.0280600000",
            "position_source": 3,
            "city": "上海",
            "contact_phone": "13900000000",
            "service_code": 1,
            "status": 1
        },
        {
            "chain_store_code": "A002",
            "chain_store_name": "饿了么BOD8",
            "address": "300弄亚都国际名园5号楼2003室",
            "latitude": "31.2306375000",
            "longitude": "121.3718891000",
            "position_source": 3,
            "city": "上海",
            "contact_phone": "13900000000",
            "service_code": 1,
            "status": 1
        }
    ]）
**/

type QueryStoreRes struct {
	ChainStoreCode string `json:"chain_store_code"`
	ChainStoreName string `json:"chain_store_name"`
	Address        string `json:"address"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	PositionSource int    `json:"position_source"`
	City           string `json:"city"`
	ContactPhone   int    `json:"contact_phone"`
	ServiceCode    int    `json:"service_code"`
	Status         int    `json:"status"`
}

func (this *QueryStore) Query(v interface{}) (err error) {
	elm := new(ELM)
	err = elm.queryStore(this)
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
