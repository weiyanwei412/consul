package main

import (
	"fmt"
	consulapi "github.com/armon/consul-api"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
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
	beego.Info("tag=" + tag + "of" + servicename + "  service exist")
	var isunhealthy bool
	for index := range healthpair {
		for checkindex := range healthpair[index].Checks {
			fmt.Println(healthpair[index].Checks[checkindex].Status)
			if healthpair[index].Checks[checkindex].Status != "passing" {
				isunhealthy = true
			}
		}
	}
	if !isunhealthy {
		beego.Info(tag + "on" + servicename + "  service  health!")
		return
	} else {
		beego.Info(tag + "on" + servicename + "  service Unhealth!")
		ip := beego.AppConfig.String("ip")
		port := beego.AppConfig.String("port")
		username := beego.AppConfig.String("username")
		password := beego.AppConfig.String("password")
		Set(ip, port, username, password)

	}

}
