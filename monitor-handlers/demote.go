package main

import (
	//	"fmt"
	consulapi "github.com/armon/consul-api"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)

func CheckService() {
	time.Sleep(10000 * time.Millisecond)
	config := &consulapi.Config{
		Datacenter: beego.AppConfig.String("datacenter"),
		Token:      beego.AppConfig.String("token"),
		Address:    beego.AppConfig.String("service_ip") + ":" + beego.AppConfig.String("service_port"),
	}
	client, err := consulapi.NewClient(config)
	if err != nil {
		beego.Error("Create consul-api client failure!", err)
		return
	}
	beego.Info(" Create consul-api client success!")
	servicename := beego.AppConfig.String("servicename")
	tag := beego.AppConfig.String("tag")
	health := client.Health()
	healthpair, _, err := health.Service(servicename, tag, false, nil)
	if err != nil {
		beego.Error("Health check execution nodes and services /v1/health/service/"+servicename+"?tag="+tag+"  failure!", err)
		return
	}
	beego.Info("Health check execution nodes and services /v1/health/service/" + servicename + "?tag=" + tag + "  success!")
	if len(healthpair) <= 0 {
		beego.Error("tag=" + tag + "of" + servicename + "  service does not exist")
		return
	}
	beego.Info("tag=" + tag + "  of  " + servicename + "  service exist")
	//	var isunhealthy bool
	var addr string
	var status string
	for index := range healthpair {
		if healthpair[index].Node.Address != "" {
			addr = healthpair[index].Node.Address
		}
		for checkindex := range healthpair[index].Checks {
			//	fmt.Println(healthpair[index].Checks[checkindex].Status)
			//if healthpair[index].Checks[checkindex].Status != "passing" {
			//		isunhealthy = true
			//	}
			if healthpair[index].Checks[checkindex].Status == "passing" {
				status = healthpair[index].Checks[checkindex].Status
			} else if healthpair[index].Checks[checkindex].Status == "warning" {
				status = healthpair[index].Checks[checkindex].Status
				break
			} else if healthpair[index].Checks[checkindex].Status == "critical" {
				status = healthpair[index].Checks[checkindex].Status
				break
			} else {
				status = "invalid"
				break
			}
		}
	}
	ip := beego.AppConfig.String("ip")
	port := beego.AppConfig.String("port")
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	if status == "passing" {
		beego.Info(tag + "  status is passing")
		return
	} else if status == "warning" {
		Switch := beego.AppConfig.String("switch")
		if strings.EqualFold(Switch, "off") {
			beego.Info("Not set asynchronous")
			return
		} else if strings.EqualFold(Switch, "on") {
			checkio_thread(ip, port, username, password, addr)
			return
		} else {
			beego.Info("config file switch format error,off or on")
			return
		}
	} else if status == "critical" {
		checkio_thread(ip, port, username, password, addr)
	} else {
		beego.Info("Not passing,not waring,not critical ,is invalid state")
		return
	}
	/*if !isunhealthy {
		Switch := beego.AppConfig.String("switch")
		if strings.EqualFold(Switch, "off") {
			beego.Info(tag + "on" + servicename + "  service  health!")
			return
		} else if strings.EqualFold(Switch, "on") {
			Set(ip, port, username, password, 1)
			return
		} else {
			beego.Info("config file switch format error,off or on")
			return
		}
		return
	} else {
		checkio_thread(ip, port, username, password, addr)

	}*/

}
