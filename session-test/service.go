package main

import (
	"bytes"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
)

func RegisterService() {
	client := new(http.Client)
	agent1url := &url.URL{
		Scheme: "http",
		Host:   "192.168.2.61:8500",
		Path:   "/v1/agent/service/register",
	}
	agent1json := `{"Name":"mysql-1","Tags":["master"],"Port":3306,"Check":{"Script":"/root/mysqlcheck.sh","Interval":"10s"}}`
	agent1reg, err := http.NewRequest("PUT", agent1url.String(), bytes.NewBufferString(agent1json))
	if err != nil {
		beego.Error("agent1 register service 创建请求失败", err)
		return
	}
	beego.Info("agent1 register service创建请求成功")
	agent1reg.Header.Set("Content-Type", "application/json")
	//执行PUT请求
	agent1resg, err := client.Do(agent1reg)
	if agent1resg.StatusCode != 200 {
		beego.Error("发送agent1 register service失败", err)
		return
	}
	beego.Info("发送agent1 register service成功")
	//读取agent1resg数据到agent1body
	agent1body, err := ioutil.ReadAll(agent1resg.Body)
	if err != nil {
		beego.Error("读取agent1 register service 返回值失败", err)
		return
	}
	beego.Info("读取agent1 register service 返回值成功",string(agent1body))
/*	agent3url := &url.URL{
		Scheme: "http",
		Host:   "192.168.2.63:8500",
		Path:   "/v1/agent/service/register",
	}
	agent3json := `{"Name":"mysql-1","Tags":["master"],"Port":3306,"Check":{"Script":"/root/mysqlcheck.sh","Interval":"10s"}}`
	agent3reg, err := http.NewRequest("PUT", agent3url.String(), bytes.NewBufferString(agent3json))
	if err != nil {
		beego.Error("agent3 register service 创建请求失败", err)
		return
	}
	beego.Info("agent3 register service 创建请求成功")
	agent3reg.Header.Set("Content-Type", "application/json")
	//执行PUT请求
	agent3resg, err := client.Do(agent3reg)
	if agent3resg.StatusCode != 200 {
		beego.Error("发送agent3 register service失败", err)
		return
	}
	beego.Info("发送agent3 register service成功")
	读取agent1resg数据到agent3body
	agent3body, err := ioutil.ReadAll(agent3resg.Body)
	if err != nil {
		beego.Error("读取agent3 register service 返回值失败", err)
		return
	}
	beego.Info("读取agent3 register service 返回值成功",agent3body)*/
}
