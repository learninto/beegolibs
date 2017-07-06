package httplib

import (
	"encoding/json"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"time"
)

var httpConnectTimeout time.Duration
var httpReadWriteTimeout time.Duration

func init() {
	//hct, _ := beego.AppConfig.Int64("httpconnecttimeout")
	//hrwt, _ := beego.AppConfig.Int64("httpconnecttimeout")
	//
	//httpConnectTimeout = time.Duration(hct) * time.Second
	//httpReadWriteTimeout = time.Duration(hrwt) * time.Second

	httpConnectTimeout = 5 * time.Second
	httpReadWriteTimeout = 5 * time.Second
}

func Get(reqUrl string) ([]byte, error) {
	logs.Info("request-get-url: " + reqUrl)
	req := httplib.Get(reqUrl).SetTimeout(httpConnectTimeout, httpReadWriteTimeout)
	data, err := req.Bytes()
	logs.Info("request-get-data: "+string(data), err)
	return data, err
}

func Post(reqUrl string, params interface{}) ([]byte, error) {
	//logs.Info("request-post-url: " + reqUrl)
	//logInfo("request-post-params: ", params)

	req, err := httplib.Post(reqUrl).SetTimeout(httpConnectTimeout, httpReadWriteTimeout).JSONBody(params)
	data, err := req.Bytes()
	//logs.Info("request-post-data: "+string(data), err)
	return data, err
}

func Put(reqUrl string, params map[string]interface{}) (content []byte, err error) {
	//logs.Info("request-put-url: " + reqUrl)
	//logs.Info("request-put-url: " + reqUrl)
	//logs.Info("request-put-params: ", params)
	//start := time.Now()
	req, err := httplib.Put(reqUrl).SetTimeout(httpConnectTimeout, httpReadWriteTimeout).JSONBody(params)
	data, err := req.Bytes()
	//logs.Info("request-put-data: "+string(data), err)
	//end := time.Now()
	//delta := end.Sub(start)
	//logs.Info("request-put-response-time:", delta)
	return data, err
}

func logInfo(msg string, params interface{}) {
	logParams, _ := json.Marshal(params)
	logs.Info(msg, string(logParams))
}
