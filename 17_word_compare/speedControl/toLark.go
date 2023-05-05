package speedControl

import (
	"encoding/json"
	"fmt"
	utils "go_Learning/100_utils"
	"go_Learning/17_word_compare/util"
	"sync"
	"time"
)

var SafeMap sync.Map

type Resp struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data TextList `json:"data"`
}

type TextList struct {
	TextList []string `json:"text_list"`
}

func PictureBasicRecognize(larkReq util.LarkMsgStruct, t string) error {
	t0 := time.Now().Unix()

	url := "https://open.feishu.cn/open-apis/optical_char_recognition/v1/image/basic_recognize"
	header := make(map[string]string, 0)
	header["Authorization"] = "Bearer t-g1043f9qQGEK46KWAGAMSEUZ5HEKUVFPXEBWASG2"
	header["Content-Type"] = "application/json; charset=utf-8"

	reqBody := struct {
		Image string `json:"image"`
	}{
		Image: larkReq.Image,
	}

	body, err := json.Marshal(&reqBody)
	if err != nil {
		return err
	}

	request, err := utils.HttpRequest(url, "POST", string(body), header)
	if err != nil {
		return err
	}
	var resp Resp
	err = json.Unmarshal(request, &resp)
	if err != nil {
		return err
	}

	//if larkReq.Sku == strconv.Itoa(10) && larkReq.ReSendCount < 2 {
	//	resp.Data.TextList = []string{}
	//}
	//if len(resp.Data.TextList) == 0 && larkReq.ReSendCount < 2 {
	//	larkReq.Position = "重发"
	//	larkReq.ReSendCount++
	//	util.SendToLarkChan(larkReq)
	//	// todo return
	//	//return nil
	//}

	fmt.Println(fmt.Sprintf("result：%v,    time: %v,    sku: %v,   position: %v,  resendcount:%v   t: %v ",
		len(resp.Data.TextList), time.Now().Unix()-t0, larkReq.Sku, larkReq.Position, larkReq.ReSendCount, t))

	// TODO 落地数据库
	//SafeMap.Store(larkReq.Sku+larkReq.Position, true)


	return nil
}


