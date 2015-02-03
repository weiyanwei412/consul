package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	consulapi "github.com/armon/consul-api"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Session struct {
	ID string
}

type Value struct {
	Key         string
	CreateIndex uint64
	ModifyIndex uint64
	LockIndex   uint64
	Flags       uint64
	Value       []byte
	Session     string
}

func IsSession() {
	client := new(http.Client)
	getkvurl := &url.URL{
		Scheme: "http",
		Host:   "192.168.2.71:8500",
		Path:   "/v1/kv/service/mysql-1/leader",
	}
	getreg, err := http.NewRequest("GET", getkvurl.String(), nil)
	if err != nil {
		beego.Error("获得service/mysql-1/leader的信息NewRequest失败", err)
		return
	}
	beego.Info("获得service/mysql-1/leader的信息NewRequest成功")
	getreg.Header.Set("Content-Type", "application/json")
	getresg, err := client.Do(getreg)
	if getresg.StatusCode != 200 {
		beego.Error("获得service/mysql-1/leader失败", err)
		return
	}
	beego.Info("获得service/mysql-1/leader成功")
	getkvbody, err := ioutil.ReadAll(getresg.Body)
	if err != nil {
		beego.Error("读取service/mysql-1/leader数据失败", err)
		return
	}
	beego.Info("读取service/mysql-1/leader数据成功", string(getkvbody))
	values := []Value{}
	err = json.Unmarshal(getkvbody, &values)
	if err != nil {
		beego.Error("解析service/mysql-1/leader信息失败", err)
		return
	}
	beego.Info("解析service/mysql-1/leader信息成功")
	if len(values) <= 0 {
		beego.Error("service/mysql-1/leader没有数据")
	}
	beego.Info("service/mysql-1/leader有数据")
	if values[0].Session != "" {
		fmt.Println("service/mysql-leader中有session ID")
		beego.Info("service/mysql-leader中有session ID")
		time.Sleep(1000)
		return
	}
	ServiceCheck()
}

func SetConn() {
	client := new(http.Client)
	sessionjson := `{"LockDelay":"15s","Name":"mysql","Node":"consul-agent1","Checks":["serfHealth","service:mysql-1"]}`
	sessionurl := &url.URL{
		Scheme: "http",
		Host:   "192.168.2.71:8500",
		Path:   "/v1/session/create",
	}
	sessionreg, err := http.NewRequest("PUT", sessionurl.String(), bytes.NewBufferString(sessionjson))
	if err != nil {
		beego.Error("session NewRequest方法执行失败:", err)
		return
	}
	beego.Info("NewRequest方法执行成功")
	sessionreg.Header.Set("Content-Type", "application/json")
	//执行PUT请求
	sessionresg, err := client.Do(sessionreg)
	beego.Info(sessionresg.StatusCode)
	if sessionresg.StatusCode != 200 {
		beego.Error("session 创建失败:", err)
		return
	}
	beego.Info("session 创建成功")
	//读取sessionresg数据到sessionbody
	sessionbody, err := ioutil.ReadAll(sessionresg.Body)
	if err != nil {
		beego.Error("读取session数据失败！", err)
		return
	}
	beego.Info("读取session数据成功")
	session := Session{}
	//解析sessionbody数据到session
	err = json.Unmarshal(sessionbody, &session)
	if err != nil {
		beego.Error("解析session ID失败:", err)
		return
	}
	beego.Info("解析session ID成功")
	beego.Info("session ID:", session.ID)
	time.Sleep(20000000000)
	rawquery := "acquire=" + session.ID
	beego.Info("acquire:", rawquery)
	kvurl := &url.URL{
		Scheme:   "http",
		Host:     "192.168.2.71:8500",
		Path:     "/v1/kv/service/mysql-1/leader",
		RawQuery: rawquery,
	}
	kvjson := `{"Node":"consul-agent1","IP":"192.168.2.61","Port":3306,"username":"root","password":"111111"}`
	kvreg, err := http.NewRequest("PUT", kvurl.String(), bytes.NewBufferString(kvjson))
	if err != nil {
		beego.Error("acquire NewRequest方法执行失败:", err)
		return
	}
	beego.Info("acquire NewRequest方法执行成功")
	kvreg.Header.Set("Content-Type", "application/json")
	//执行PUT请求
	kvresp, err := client.Do(kvreg)
	if kvresp.StatusCode != 200 {
		beego.Error("acquire 加入锁失败:", err)
		return
	}
	beego.Info("acquire 加入锁成功")
	defer kvresp.Body.Close()
	//读取kvresg数据到kvbody
	kvbody, err := ioutil.ReadAll(kvresp.Body)
	if err != nil {
		beego.Error("读取 acquire 数据失败:", err)
		return
	}
	beego.Info("读取 acquire 数据成功")
	beego.Info("kv acquire:", string(kvbody))

}

func Empty() {
	client := new(http.Client)
	emptykvurl := &url.URL{
                Scheme: "http",
                Host:   "192.168.2.71:8500",
                Path:   "/v1/kv/service/mysql-1/leader",
        }
        emptyreg, err := http.NewRequest("PUT", emptykvurl.String(), nil)
        if err != nil {
                beego.Error("更新service/mysql-1/leader的NewRequest失败", err)
                return
        }
        beego.Info("更新service/mysql-1/leader的NewRequest成功")
        emptyreg.Header.Set("Content-Type", "application/json")
        emptyresg, err := client.Do(emptyreg)
        if emptyresg.StatusCode != 200 {
                beego.Error("更新/service/mysql-1/leader失败", err)
                return
        }
        beego.Info("更新service/mysql-1/leader成功")
        slave()
}

func ServiceCheck() {
	client := new(http.Client)
	healthmysqlurl := &url.URL{
		Scheme: "http",
		Host:   "192.168.2.71:8500",
		Path:   "/v1/health/checks/mysql-1",
	}
	healthmysqlreg, err := http.NewRequest("GET", healthmysqlurl.String(), nil)
	if err != nil {
		beego.Error("获得mysql-1服务状态的NewRequest失败", err)
		return
	}
	beego.Info("获得mysql-1服务状态的NewRequest成功", err)
	healthmysqlreg.Header.Set("Content-Type", "application/json")
	healthmysqlresg, err := client.Do(healthmysqlreg)
	if healthmysqlresg.StatusCode != 200 {
		beego.Error("获得mysql-1服务状态失败", err)
		return
	}
	beego.Info("获得mysql-1服务状态成功")
	gethealthbody, err := ioutil.ReadAll(healthmysqlresg.Body)
	if err != nil {
		beego.Error("mysql-1服务状态返回信息读取失败", err)
		return
	}
	beego.Info("mysql-1服务状态返回信息读取成功")
	healthvalue := []consulapi.HealthCheck{}
	err = json.Unmarshal(gethealthbody, &healthvalue)
	if err != nil {
		beego.Error("mysql-1服务状态返回信息解析失败", err)
		return
	}
	beego.Info("mysql-1服务状态返回信息解析成功")
	if len(healthvalue) <= 0 {
		beego.Error("没有mysql-1服务")
		return
	}
	beego.Info("有mysql-1服务")
	var islocal bool
	for index := range healthvalue {
		if healthvalue[index].Node == "consul-agent2" {
			islocal = true
			beego.Info("本机mysql-1服务正常")
			break
		}

	}
	if !islocal {
		beego.Info("本机的mysql-1服务不正常")
		return
	}else {
		Empty()
	}
}
