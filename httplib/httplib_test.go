package httplib

import (
	"encoding/json"
	"github.com/siddontang/go/log"
	"testing"
)

type TestGetIo struct {
	Status int64  `json:"status"`
	Data   string `json:"data"`
}

func TestGet(t *testing.T) {
	body, terr := Get("http://m.sh.189.cn/service/node/crypto?data=oKXUCj1MOddnp-sCpGi1J1dg3TyM,abc&key=wechat-mobile-201604&type=0")
	if terr != nil {

	}
	//{"status":200,"data":"9abd60a5334646b5498a490da70972b6c22902ddcad30fc929a1411f5b2360c6,532efc805a8aa7a31b84640bf05ec726"}

	var tji TestGetIo
	json.Unmarshal(body, &tji)
	jso, _ := json.Marshal(&tji)
	log.Info("TestGet Response" + string(jso))
}

//func TestPost(t *testing.T) {
//	var PayParams map[string]interface{}
//	PayParams = make(map[string]interface{});
//	PayParams["money"] = "5"
//	PayParams["number"] = "17721021494"
//	PayParams["bpnum"] = ""
//	PayParams["payid"] = ""
//	PayParams["openid"] = "oKXUCj1MOddnp-sCpGi1J1dg3TyM"
//	PayParams["from"] = "disney"
//	PayParams["channel"] = "0"
//	PayParams["note"] = "迪士尼活动"
//	body, terr := Post("http://httpbin.org/post", PayParams);
//	if terr != nil {
//
//	}
//	fmt.Println("Post 请求结果" + string(body))
//}
//
//func TestPostForm(t *testing.T) {
//	var PayParams map[string]interface{}
//	PayParams = make(map[string]interface{});
//	PayParams["mobile"] = "17721021494"
//	PayParams["channel"] = "2"
//	PayParams["desKey"] = "dzqd-wt-flow"
//
//	body, terr := PostForm("http://172.16.50.138:8091/csb/1.0/Encrypt", PayParams);
//	if terr != nil {
//
//	}
//	fmt.Println("PostForm 请求结果" + string(body))
//}

//func TestPut(t *testing.T) {
//	var PayParams map[string]interface{}
//	PayParams = make(map[string]interface{})
//	PayParams["contact_id_card"] = "1111"
//	PayParams["contact_name"] = "111"
//	PayParams["contact_phone"] = "111"
//	PayParams["into_park_time"] = "11"
//
//	body, terr := Put("http://httpbin.org/put", PayParams)
//	//body, terr := Put("http://172.16.50.141:8080/v1/disney_pay/draw_prize/123", PayParams);
//	if terr != nil {
//
//	}
//	fmt.Println("put 请求结果" + string(body))
//}
