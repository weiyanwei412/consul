package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"database/sql"
	consulapi "github.com/armon/consul-api"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type Kv struct {
	Node     string
	Ip       string
	Port     int
	Username string
	Password string
}

func GetLeader() error {
	config := &consulapi.Config{
		Datacenter: beego.AppConfig.String("datacenter"),
		Token:      beego.AppConfig.String("token"),
		Address:    beego.AppConfig.String("server_ip") + ":" + beego.AppConfig.String("server_port"),
	}
	servicename := beego.AppConfig.String("servicename")
	client, err := consulapi.NewClient(config)
	if err != nil {
		beego.Error("Create a consul-api client fails", err)
		return err
	}
	kv := client.KV()
	//Get is used to lookup a single key
	for {
		kvPair, _, err := kv.Get("service/"+servicename+"/leader", nil)
		if err != nil {
			beego.Error("Get a key failure", err)
			return err
		}
		if kvPair == nil {
			beego.Error("service/" + servicename + "/leader not found, please create the key.")
			return err
		}
		//Are there external connection string provided
		if kvPair.Session != "" {
			beego.Info("Connection string has been set")
			if kvPair.Value != nil {
				beego.Info("service/mysql-1/leader set value")
				break
			} else {
				beego.Info("service/mysql-1/leader no set value")
				continue
			}
		}
	}
	err = Conn()
	if err != nil {
		beego.Error("conn mysql fails:", err)
		return nil
	}
	return nil

}

func Conn() error {
	server_ip := beego.AppConfig.String("server_ip")
	server_port := beego.AppConfig.String("server_port")
	table_name := beego.AppConfig.String("table_name")
	servicename := beego.AppConfig.String("servicename")
	database := beego.AppConfig.String("database")
	cli := new(http.Client)
	url := &url.URL{
		Scheme:   "http",
		Host:     server_ip + ":" + server_port,
		Path:     "/v1/kv/service/" + servicename + "/leader",
		RawQuery: "raw",
	}
	reg, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		beego.Error("Create a get request fails:", err)
		return err
	}
	reg.Header.Set("Content-Type", "application/json")
	//执行get请求
	resp, err := cli.Do(reg)
	if err != nil {
		beego.Error("Send http request fails:", err)
		return err
	}
	//分隔符
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		beego.Error("Failed to read data:", err)
		return err
	}
	kv := Kv{}
	//解析body到kv
	err = json.Unmarshal([]byte(body), &kv)
	if err != nil {
		beego.Error("Failed parsing json:", err)
		return err
	}
	//连接mysql数据库
	db, err := sql.Open("mysql", kv.Username+":"+kv.Password+"@tcp("+kv.Ip+":"+strconv.Itoa(kv.Port)+")/"+database+"?charset=utf8")
	//fmt.Println("username", kv.Username)
	if err != nil {
		beego.Error("Connect to the database fails:", err)
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		beego.Error("db.ping()", err)
		return err
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	//sql := "select * from" + table_name
	raw, err := db.Query("select * from " + table_name)
	if err != nil {
		beego.Error("select data failed:", err)
		return err
	}
	defer raw.Close()
	for raw.Next() {
		var id int
		var names string
		var created_at string
		err = raw.Scan(&id, &names, &created_at)
		if err != nil {
			beego.Error("error", err)
		}
		fmt.Println("out", id, names, created_at)
	}
	//当前时间
	nowtime := time.Now().Format("2006-01-02 15:04:05")
	sqlstr := "insert into " + table_name + "(`name`,`created_at`) values(?,?)"
	if _, err := tx.Exec(sqlstr, "lindan", nowtime); err != nil {
		beego.Error("Insert data failed:", err)
		tx.Rollback()
		return err
	}
	row, err := db.Query("select name from " + table_name + " order by id desc limit 1")
	if err != nil {
		beego.Error("select data failed:", err)
		tx.Rollback()
		return err
	}
	defer row.Close()
	var name string
	for row.Next() {
		err = row.Scan(&name)
		if err != nil {
			beego.Error("error", err)
		}
		fmt.Println("output", name)
	}
	sqlstr = "update " + table_name + " set name='lindan1' where name=? order by id desc limit 1"
	if _, err := tx.Exec(sqlstr, name); err != nil {
		beego.Error("Insert data failed:", err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		beego.Error("commit failed:", err)
		tx.Rollback()
		return err
	}
	return nil
}
