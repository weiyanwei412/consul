package main

import (
	"bytes"
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Session struct {
	ID string
}

type Value struct {
	Node string
	IP string
	Port int
	username string
	password string
	holdon string
}

func PUT() {	
	time.Sleep(5000000000)
	client := new(http.Client)
	kvgeturl := &url.URL{
		Scheme: "http",
		Host: "192.168.2.71:8500",
		Path: "/v1/kv/service/mysql-1/leader",
		RawQuery: "raw",
	}
	//
	kvgetreg, err := http.NewRequest("GET",kvgeturl.String(),nil)
	if err != nil {
		beego.Error("get kv value request failed!", err)
		return
	}
	beego.Info("get kv value request Success1!")
	kvgetreg.Header.Set("Content-Type","application/json")
	//执行GET请求
	kvgetresg, err := client.Do(kvgetreg)
	if kvgetresg.StatusCode !=200 {
		beego.Error("get kv api 发送失败!",err)
		return
	}
	beego.Info("get kv api 发送成功!")
	//读kvgetresg里面的数据到kvgetbody
	kvgetbody, err := ioutil.ReadAll(kvgetresg.Body)
/*	if err != nil {
		beego.Error("读取get kv value 数据失败!",err)
		return
	}*/
	beego.Info("body:",string(kvgetbody))
	//判断kvgetbody里面的数据是否为空，如果不为空表示本机是当前leader
	if string(kvgetbody) != "" {
//		beego.Error("kvgetbody为空!")
//		return
		beego.Info("body:",string(kvgetbody))
		beego.Info("读取 get kv value 数据成功!",err)
		value := Value{}
		err = json.Unmarshal(kvgetbody,&value)
		if err !=nil {
			beego.Error("get kv value 解析失败!",err)
			return
		}
		beego.Info("get kv value 解析成功!")
		if value.Node == "consul-agent1" {
			beego.Info("本机是leader")
			return
		}
	}
	time.Sleep(500000000)
	sessionjson := `{"LockDelay":"15s","Name":"mysql","Node":"consul-agent1","Checks":["serfHealth","service:mysql-1"]}`
	sessionurl := &url.URL{
		Scheme: "http",
		Host:   "192.168.2.71:8500",
		Path:   "/v1/session/create",
	}
	sessionreg, err := http.NewRequest("PUT", sessionurl.String(), bytes.NewBufferString(sessionjson))
	if err != nil {
		beego.Error("session request failed:", err)
		return
	}
	beego.Info("session request Success")
	sessionreg.Header.Set("Content-Type", "application/json")
	//执行PUT请求
	sessionresg, err := client.Do(sessionreg)
	beego.Info(sessionresg.StatusCode)
	if sessionresg.StatusCode != 200{
		beego.Error("session api 发送失败！", err)
		return
	}
	beego.Info("session api 发送成功")
	//读取sessionresg数据到sessionbody
	sessionbody, err := ioutil.ReadAll(sessionresg.Body)
	if err != nil {
		beego.Error("读取session数据失败！", err)
		return
	}
	session := Session{}
	//解析sessionbody数据到session
	err = json.Unmarshal(sessionbody, &session)
	if err != nil {
		beego.Error("session 解析失败！", err)
		return
	}
	beego.Info("session 解析成功！")
	beego.Info("session ID:", session.ID)
	time.Sleep(10000000000)
	rawquery := "acquire=" + session.ID
	kvurl := &url.URL{
		Scheme:   "http",
		Host:     "192.168.2.71:8500",
		Path:     "/v1/kv/service/mysql-1/leader",
		RawQuery: rawquery,
	}
	kvjson := `{"Node":"consul-agent1","IP":"192.168.2.61","Port":3306,"username":"root","password":"111111","holdon":"N"}`
	kvreg, err := http.NewRequest("PUT", kvurl.String(), bytes.NewBufferString(kvjson))
	if err != nil {
		beego.Error("kv request failed:", err)
		return
	}
	beego.Info("kv request Success!")
	kvreg.Header.Set("Content-Type", "application/json")
	//执行PUT请求
	kvresp, err := client.Do(kvreg)
	if kvresp.StatusCode != 200 {
		beego.Error("kv acquire api 发送失败！", err)
		return
	}
	beego.Info("kv acquire  发送成功！")
	defer kvresp.Body.Close()
	//读取kvresg数据到kvbody
	kvbody, err := ioutil.ReadAll(kvresp.Body)
	if err != nil {
		beego.Error("读取kv 数据 失败！:", err)
		return
	}
	beego.Info("读取kv  数据成功！")
	beego.Info("kv acquire:",string(kvbody))
}
