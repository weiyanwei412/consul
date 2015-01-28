package main

import (
	"bytes"
	//"database/sql"
	"encoding/json"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Session struct {
	ID string
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
		beego.Error("更新/ervice/mysql-1/leader的NewRequest失败", err)
		return
	}
	beego.Info("更新/ervice/mysql-1/leader的NewRequest成功")
	emptyreg.Header.Set("Content-Type", "application/json")
	emptyresg, err := client.Do(emptyreg)
	if emptyresg.StatusCode != 200 {
		beego.Error("更新/service/mysql-1/leader失败", err)
		return
	}
	beego.Info("更新service/mysql-1/leader成功")

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
	//time.Sleep(10000000000)
	rawquery := "acquire=" + session.ID
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
