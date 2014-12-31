package main

import (
        "bytes"
        "encoding/json"
        "github.com/astaxie/beego"
        "io/ioutil"
        "net/http"
        "net/url"
)

type Session struct {
        ID string
}

func SessionCearte()(sessionoutput string){
	client := new(http.Client)	
	sessionurl := &url.URL{
                Scheme: "http",
                Host:   "192.168.2.71:8500",
                Path:   "/v1/session/create",
        }
        sessionjson := `{"LockDelay":"15s","Name":"mysql","Node":"consul-agent1","Checks":["serfHealth","service:mysql-1"]}`
//      sessionjson := `{"LockDelay":"15s","Name":"mysql","Node":"consul-agent1","Checks":["serfHealth"]}`
        sessionreg, err := http.NewRequest("PUT",sessionurl.String(),bytes.NewBufferString(sessionjson))
        if err != nil {
                beego.Error("session create请求创建失败", err)
                return
        }
        beego.Info("session create请求创建成功")
        sessionreg.Header.Set("Content-Type","application/json")
	//执行PUT请求
        sessionresg, err := client.Do(sessionreg)
        beego.Info(sessionresg.StatusCode)
        if sessionresg.StatusCode != 200 {
                beego.Error("发送session create请求失败", err)
                return
        }
        beego.Info("发送session create请求成功")
	//读取sessionresg数据到sessionbody
        sessionbody, err := ioutil.ReadAll(sessionresg.Body)
        if err != nil {
                beego.Error("读取session create 返回值失败", err)
                return
        }
        beego.Info("读取session create 返回值成功",string(sessionbody))
        session := Session{}
	//解析sessionbody数据到session
        err = json.Unmarshal(sessionbody, &session)
        if err != nil {
                beego.Error("解析session create 返回值失败", err)
                return
        }
        beego.Info("解析session create返回值成功", session.ID)
	return session.ID
}
