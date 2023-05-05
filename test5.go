package main

import (
	"encoding/json"
	"fmt"
	utils "go_Learning/100_utils"
	types "go_Learning/101_types"
)

type zxReq struct {
	OrderNumber string `json:"order_number"`
}

type zxResp struct {
	types.BaseResponse
	Data data `json:"data"`
}

type data struct {
	OrderNumber  string   `json:"order_number"`
	SupplierCode string   `json:"supplier_code"`
	FactoryId    int      `json:"factory_id"`
	Detail       []detail `json:"detail"`
}

type detail struct {
	Ctn       string `json:"ctn"`
	MapSn     string `json:"map_sn"`
	RepairNum int    `json:"repair_num"`
}

const number = "PO2311XCSJ010009-4-1"

func main() {
	url := "http://47.89.13.192:6001/synchronous/get-sampling-data"
	header := make(map[string]string, 0)
	header["Authorization"] = "@!89%Tq3Ikol()k|~d!Ez09"
	header["Content-Type"] = "application/json"

	var body zxReq
	body.OrderNumber = number
	data, err := json.Marshal(body)
	if err != nil {
		fmt.Println("SearchQcExamineGoodsInfo err: ", err)
		return
	}
	zxData, err := utils.HttpRequest(url, "POST", string(data), header)
	if err != nil {
		fmt.Println("zx httpUtils.HttpPost err: ", err)
		return
	}
	var resp zxResp
	err = json.Unmarshal(zxData, &resp)
	if err != nil {
		fmt.Println("zx json.Unmarshal err: ", err)
		return
	}
	if resp.ResCode != 0 {
		fmt.Println("=========1: ", resp)
	} else {
		fmt.Println("=========2: ", resp)
	}
}
