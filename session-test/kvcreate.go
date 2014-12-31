package main

import (
        "bytes"
        "github.com/astaxie/beego"
        "io/ioutil"
        "net/http"
        "net/url"
)

func KvCreate() {
        client := new(http.Client)
        kvurl := &url.URL{
                Scheme: "http",
                Host:   "192.168.2.71:8500",
                Path:   "/v1/kv/service/mysql-1/leader",
        }
        kvjson := `{"Node":"consul-agent2","IP":"192.168.2.62","Port":3306,"username":"root","password":"111111","holdon":"N"}`
        kvreg, err := http.NewRequest("PUT",kvurl.String(),bytes.NewBufferString(kvjson))
        if err != nil {
                beego.Error("create kv 请求创建失败", err)
                return
        }
        beego.Info("create kv 请求创建成功")
        kvreg.Header.Set("Content-Type", "application/json")
	//执行PUT请求
        kvresg, err := client.Do(kvreg)
        if kvresg.StatusCode != 200 {
                beego.Error("发送create kv 请求失败", err)
                return
        }
        beego.Info("发送create kv请求成功")
	//读取kvresg数据到kvbody
        kvbody, err := ioutil.ReadAll(kvresg.Body)
        if err != nil {
                beego.Error("读取create kv 返回值失败", err)
                return
        }
        beego.Info("读取create kv请求成功",string(kvbody))
}
