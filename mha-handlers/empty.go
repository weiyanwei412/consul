package main

import (
	//	"bytes"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"net/url"
)

func Empty() {
	client := new(http.Client)
	emptykvurl := &url.URL{
		Scheme: "http",
		Host:   "192.168.2.71:8500",
		Path:   "/v1/kv/service/mysql-1/leader",
	}
	//emptykv := ""
	emptyreg, err := http.NewRequest("PUT", emptykvurl.String(), nil)
	if err != nil {
		beego.Error("Update service/mysql-1/leader request failed!", err)
		return
	}
	beego.Info("Update service/mysql-1/leader request success!")
	emptyreg.Header.Set("Content-Type", "application/json")
	emptyresg, err := client.Do(emptyreg)
	if emptyresg.StatusCode != 200 {
		beego.Error("Update service/mysql-1/leader failed!", err)
		return
	}
	beego.Info("Update service/mysql-1/leader success!")

}
