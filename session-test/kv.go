package main

import (
	"bytes"
	"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func KvAcquire(sessioninput string) {
	client := new(http.Client)
	time.Sleep(5000000000)
	beego.Info(sessioninput)
	acquireurl := &url.URL{
		Scheme:   "http",
		Host:     "192.168.2.71:8500",
		Path:     "/v1/kv/service/mysql-1/leader",
		RawQuery: "acquire=" + sessioninput,
	}
	acquirejson := `{"Node":"consul-agent1","IP":"192.168.2.61","Port":3306,"username":"root","password":"","holdon":"N"}`
	acquirereg, err := http.NewRequest("PUT", acquireurl.String(), bytes.NewBufferString(acquirejson))
	if err != nil {
		beego.Error("acquire请求创建失败", err)
		return
	}
	beego.Info("acquire请求创建成功")
	acquirereg.Header.Set("Content-Type","application/json")
	//执行PUT请求
	acquireresg, err := client.Do(acquirereg)
	beego.Info(acquireresg.StatusCode)
	if acquireresg.StatusCode != 200 {
		beego.Error("发送acquire失败", err)
		return
	}
	beego.Info("发送acquire成功")
	defer acquireresg.Body.Close()
	//读取acquireresg数据到acquirebody
	acquirebody, err := ioutil.ReadAll(acquireresg.Body)
	if err != nil {
		beego.Error("读取acquire操作返回值失败", err)
		return
	}
	beego.Info("读取acquire操作返回值成功",string(acquirebody))
	if string(acquirebody) != "true" {
		beego.Error("添加锁失败", err)
	}
	beego.Info("添加锁成功")
}

func KvRelease(sessioninput string){
	time.Sleep(10000000000)
	client := new(http.Client)
        releaseurl := &url.URL{
                Scheme:   "http",
                Host:     "192.168.2.71:8500",
                Path:     "/v1/kv/service/mysql-1/leader",
                RawQuery: "release=" + sessioninput,
        }
        releasereg, err := http.NewRequest("PUT", releaseurl.String(), nil)
        if err != nil {
                beego.Error("release请求创建失败", err)
                return
        }
        beego.Info("release请求创建成功")
        releasereg.Header.Set("Content-Type", "application/json")
	//执行PUT请求
        releaseresg, err := client.Do(releasereg)
        if releaseresg.StatusCode != 200 {
                beego.Error("发送release请求失败", err)
                return
        }
        beego.Info("发送release请求成功")
        defer releaseresg.Body.Close()
	//读取releaseresg数据到releasebody
        releasebody, err := ioutil.ReadAll(releaseresg.Body)
        if err != nil {
                beego.Error("读取release操作返回值失败")
                return
        }
        beego.Info("读取release 操作返回值成功")
        if string(releasebody) != "true" {
                beego.Error("释放锁失败", err)
                return
        }
        beego.Info("释放锁成功")

}
